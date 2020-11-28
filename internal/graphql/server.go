package graphql

import (
	"net/http/pprof"

	"context"
	"report/internal/graphql/store"
	"report/internal/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Server ..
type Server struct {
	Config ServerConfig

	Loggr  logger.Loggr
	Logger *logrus.Entry
	DB     store.Database
	Router *mux.Router
	// Providers providers.Providers
	Ctx    context.Context
	Report store.Report
}

// New ..
// Creates a new server
func New(ctx context.Context, c ServerConfig) (*Server, error) {

	s := &Server{}
	s.Config = c
	s.Ctx = ctx
	s.Loggr = c.Logger
	s.Logger = s.Loggr.GetLogger("shiftserver")
	// s.NSQ.Consumer.Address = c.NSQ.ConsumerAddress
	// s.NSQ.Producer.Address = c.NSQ.ProducerAddress

	s.DB = store.NewDatabase(c.Store.Name, c.Session)
	// s.Shift = s.DB.InitShiftStore()
	// s.Resolver = &resolver.Shift{}

	r := mux.NewRouter()
	s.Router = r

	// // pprof
	r.HandleFunc("/debug/pprof", pprof.Index)
	r.HandleFunc("/debug/symbol", pprof.Symbol)
	r.HandleFunc("/debug/profile", pprof.Profile)

	// r.HandleFunc("/debug/pprof/", pprof.Index)
	// r.HandleFunc("/debug/pprof/heap", pprof.Index)
	// r.HandleFunc("/debug/pprof/mutex", pprof.Index)
	// r.HandleFunc("/debug/pprof/goroutine", pprof.Index)
	// r.HandleFunc("/debug/pprof/threadcreate", pprof.Index)
	// r.HandleFunc("/debug/pprof/block", pprof.Index)
	// r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	r.Handle("/debug/heap", pprof.Handler("heap"))
	r.Handle("/debug/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/block", pprof.Handler("block"))

	// s.Providers = providers.New(s.Loggr, s.Shift)

	// initialize graphql based services
	// err := s.registerGraphQLServices()
	// if err != nil {
	// 	return nil, err
	// }

	// initialize oauth2 providers
	// s.registerEndpointServices()

	// s.registerWebSocketServices()

	// 预留
	err := s.bootstrap()
	if err != nil {
		return nil, err
	}

	// err := NewAuthServer(ctx, r, c)
	// if err != nil {
	// 	return nil, err
	// }

	return s, nil
}

func (s *Server) registerGraphQLServices() error {
	r := s.Router

	// initialize graphql
	h := handler.New(&handler.Config{
		// Schema:   schm,
		Pretty:   true,
		GraphiQL: true,
	})
	r.Handle("/graphql", h)
	r.Handle("/graphql/", h)

	// Graphql endpoint works with websocket only for subscription
	// psh := pubsub.NewGraphqlWSHandler(s.Pubsub, s.Loggr)
	// r.Handle("/subscription", psh)

	return nil
}

func (s *Server) registerWebSocketServices() {
	//r := s.Router
	//r.HandleFunc("/api/ws/")
}

// Registers the GRPC services ...
// func (s *Server) registerGRPCServices(grpcServer *grpc.Server) {
// 	api.RegisterShiftServer(grpcServer, shift.NewServer(s.Loggr, s.Ctx, s.Shift, s.Vault, s.Pubsub, s.Resolver))
// }

// RegisterHTTPServices Registers the exposed http services
func RegisterHTTPServices(ctx context.Context, router *mux.Router, grpcAddress string, dialopts []grpc.DialOption) error {
	return nil
}
