package graphql

import (
	"net/http/pprof"

	"context"
	"report/internal/graphql/identity/oauth2/providers"
	"report/internal/graphql/store"
	"report/internal/graphql/vcs"
	"report/internal/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Server ..
type Server struct {
	Config ServerConfig

	Loggr     logger.Loggr
	Logger    *logrus.Entry
	DB        store.Database
	Router    *mux.Router
	Providers providers.Providers
	Ctx       context.Context
	Report    store.Report
}

// New ..
// Creates a new server
func New(ctx context.Context, c ServerConfig) (*Server, error) {

	s := &Server{}
	s.Config = c
	s.Ctx = ctx
	s.Loggr = c.Logger
	s.Logger = s.Loggr.GetLogger("maintain-report")

	s.DB = store.NewDatabase(c.Store.Name, c.Session)
	s.Report = s.DB.InitReportStore()
	
	// resolver...
	// s.Resolver = &resolver.Shift{}

	r := mux.NewRouter()
	s.Router = r

	// // pprof
	r.HandleFunc("/debug/pprof", pprof.Index)
	r.HandleFunc("/debug/symbol", pprof.Symbol)
	r.HandleFunc("/debug/profile", pprof.Profile)

	r.Handle("/debug/heap", pprof.Handler("heap"))
	r.Handle("/debug/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/block", pprof.Handler("block"))

	s.Providers = providers.New(s.Loggr, s.Report)

	// initialize graphql based services
	err := s.registerGraphQLServices()
	if err != nil {
		return nil, err
	}

	// initialize oauth2 providers
	s.registerEndpointServices()

	// initialize websocket2 service
	// s.registerWebSocketServices()

	// 预留
	err = s.bootstrap()
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
	r.Handle("/graphql", nil)
	r.Handle("/graphql/", nil)

	// Graphql endpoint works with websocket only for subscription
	// psh := pubsub.NewGraphqlWSHandler(s.Pubsub, s.Loggr)
	// r.Handle("/subscription", psh)

	return nil
}

func (s *Server) registerEndpointServices() {

	// VCS service to link repositories.
	vcsServ := vcs.NewService(s.Loggr, s.DB, s.Providers, s.Report)

	// Oauth2 providers
	s.Router.HandleFunc("/api/link/{provider}", vcsServ.Authorize)
	s.Router.HandleFunc("/api/link/{provider}/callback", vcsServ.Authorized)
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
