package graphql

import (
	"report/internal/graphql/store"
	"report/internal/pkg/logger"

	"gopkg.in/mgo.v2"
)

// Config ..
type Config struct {
	Store    Store    `json:"store"`
	Web      Web      `json:"web"`
	Logger   Logger   `json:"logger"`
	Identity Identity `json:"dex"`
}

// ServerConfig ..
type ServerConfig struct {
	Store    store.Config
	Logger   logger.Loggr
	Session  *mgo.Session
	Identity Identity
}

// Identity ..
type Identity struct {
	Issuer      string
	HostAndPort string
	caPath      string
	ID          string
	Secret      string
	RedirectURI string
}

// Store ..
type Store struct {
	Server    string
	Name      string
	Username  string
	Password  string
	Timeout   string
	Monotonic bool
	Retry     string // duration

	// old info
	IdleConnection int
	MaxConnection  int
	Log            bool
}

// Web ..
// Holds the web server configuration
type Web struct {
	HTTP string `json:"http"`
	GRPC string `json:"grpc"`
}

// Logger ..x``
type Logger struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}
