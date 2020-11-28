package graphql

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"report/internal/graphql/store"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/justinas/alice"

	"report/internal/pkg/logger"
	"time"

	"gopkg.in/yaml.v2"
)

// Run ...
// run server
func Run() error {
	cfgFile := "/Users/shh/gs/report/config/config.yaml"
	cfgData, err := ioutil.ReadFile(cfgFile)

	var c Config
	switch {
	case os.IsNotExist(err):
	case err == nil:
		if err := yaml.Unmarshal(cfgData, &c); err != nil {
			return fmt.Errorf("Failed to parse config file %s: %v", cfgFile, err)
		}
	default:
		log.Println(err)
	}

	// override configuration with environment variables.
	if storeServer := os.Getenv("STORE_SERVER"); storeServer != "" {
		c.Store.Server = storeServer
	}

	// server Config
	sc := ServerConfig{}

	// logger
	loggr, err := logger.New(c.Logger.Level, c.Logger.Format)
	if err != nil {
		return fmt.Errorf("invalid config: %v", err)
	}

	if c.Logger.Level != "" {
		log.Println(fmt.Printf("config using log level: %s", c.Logger.Level))
	}
	sc.Logger = loggr

	logger := loggr.GetLogger("main")

	// parse MongoDB configuration
	// parse db config
	sc.Store.Timeout, err = time.ParseDuration(c.Store.Timeout)
	if err != nil {
		return fmt.Errorf("Failed to parse database timeout duration %s :%v", c.Store.Timeout, err)
	}

	sc.Store.RetryIn, err = time.ParseDuration(c.Store.Retry)
	if err != nil {
		return fmt.Errorf("Failed to parse database retryin duration %s :%v", c.Store.Retry, err)
	}
	sc.Store.AutoReconnect = true

	ctx := context.Background()

	// set rest of databse properties to server config
	sc.Store.Name = c.Store.Name
	sc.Store.Username = c.Store.Username
	sc.Store.Password = c.Store.Password
	sc.Store.Server = c.Store.Server
	sc.Store.Monotonic = c.Store.Monotonic

	// open the db connection & retries
	session, err := store.Connect(sc.Logger, sc.Store)
	if err != nil {
		logger.Fatalln(fmt.Errorf("Failed to initalize store (database) connection : %v", err))
	}
	sc.Session = session
	defer session.Close()

	// cors policy
	corsOpts := handlers.AllowedOrigins([]string{"*"})
	corsHandler := handlers.CORS(corsOpts)
	recoveryHandler := handlers.RecoveryHandler()

	publicChain := alice.New(recoveryHandler, corsHandler)

	s, err := New(ctx, sc)
	if err != nil {
		logger.Fatalln(fmt.Errorf("Failed to initialize the server [%v]", err))
	}

	// startup two channel for http and grpc server, waiting for failed
	errch := make(chan error, 1)

	var serv *http.Server
	//start http
	go func() {
		errch <- func() error {
			dialopts := []grpc.DialOption{grpc.WithInsecure()}
			router := mux.NewRouter()
			err := RegisterHTTPServices(ctx, router, c.Web.GRPC, dialopts)
			if err != nil {
				return fmt.Errorf("Error when registering services.. : %v", err)
			}
			s.Router.Handle("/", router)
			serv := &http.Server{Addr: c.Web.HTTP, Handler: publicChain.Then(s.Router)}
			logger.Info("Exposing HTTP services on ", c.Web.HTTP)
			err = serv.ListenAndServe()
			return fmt.Errorf("Listing on %s failed with : %v", c.Web.HTTP, err)
		}()
	}()

	//graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	// http server
	go func() {

		errch <- func() error {
			<-sigs
			// logger.Infoln("Stopping GRPC Server..")
			// grpcServer.GracefulStop()
			logger.Infoln("Stopping HTTP(S) Server..")
			serv.Shutdown(ctx)
			return fmt.Errorf("Server gracefully stopped ")
		}()
	}()

	return <-errch

}
