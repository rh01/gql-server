package store

import (
	"fmt"
	"report/internal/pkg/logger"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/sirupsen/logrus"
)

// Config ..
type Config struct {
	Server        string
	Name          string
	Username      string
	Password      string
	Timeout       time.Duration
	Monotonic     bool
	AutoReconnect bool

	// old info
	IdleConnection int
	MaxConnection  int
	Log            bool
	RetryIn        time.Duration
}

// Connect ..
// Open the database connection and returns the session
func Connect(loggr logger.Loggr, cfg Config) (*mgo.Session, error) {

	logger := loggr.GetLogger("maintain-report")

	// DB Initialization
	var session *mgo.Session
	var err error
	tryit := true
	for tryit {

		logger.Infoln("Connecting to database...")
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{cfg.Server},
			Username: cfg.Username,
			Password: cfg.Password,
			Database: cfg.Name,
			Timeout:  cfg.Timeout,
		})
		if err != nil {

			if !cfg.AutoReconnect {
				return nil, err
			}

			logger.Errorln(fmt.Sprintf("Connecting database failed, retrying in %v [", cfg.RetryIn), err, "]")
			time.Sleep(cfg.RetryIn)

		} else {

			// Ping function checks the database connectivity
			dberr := session.Ping()
			if dberr != nil {
				logger.Errorln(fmt.Sprintf("Ping DB failed, retrying in %v []", cfg.RetryIn), err, "]")
			} else {
				logger.Infoln("Database connected successfully")
				tryit = false
			}
		}
	}

	// set the configurations
	session.SetMode(mgo.Monotonic, cfg.Monotonic)

	// starting a background process to automatically reconnect to database in case of disconnects.
	if cfg.AutoReconnect {
		autoReconnect(logger, cfg, session)
	}
	return session, nil
}

// autoReconnect ..
// Launches a separate go routine to perform below operations.
// 1. Ping the db session to ensure the conenction is live
// 2. Retries to conenect with database if connectivity broke.
func autoReconnect(logger *logrus.Entry, cfg Config, session *mgo.Session) {
	go func() {
		disconnected := false
		for {
			time.Sleep(cfg.RetryIn)
			// Ping function checks the database connectivity
			err := session.Ping()
			if err != nil {
				disconnected = true
				logger.Errorln(fmt.Sprintf("DB ping failed, something went wrong. Reconnecting in %d seconds", cfg.RetryIn))

				//Trying to refresh the db connection
				session.Refresh()
			} else if disconnected {
				logger.Infoln("Reconnected to database successfully.")
				disconnected = false
			}
		}
	}()
}
