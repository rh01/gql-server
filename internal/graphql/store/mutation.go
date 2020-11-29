package store

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
	"report/internal/pkg/logger"
)

type mutation struct {
	Store
	Loggr  logger.Loggr
	Logger *logrus.Entry
}

// Mutation ...
type Mutation interface {
	Interface

	// Cap mutations
	DeleteCap(id bson.ObjectId) (*models.DeleteCap, error)
	CreateCap(input *models.CreateCapInput) (*models.Cap, error)
	UpdateCap(id bson.ObjectId, input models.UpdateCapInput) (*models.UpdateCap, error)

	// Ticket mutations
	DeleteTicket(id bson.ObjectId) (*models.DeleteTicket, error)
	CreateTicket(input *models.CreateTicketInput) (*models.Ticket, error)
	UpdateTicket(id bson.ObjectId, input models.UpdateTicketInput) (*models.UpdateTicket, error)

	//OnlineCount mutations
	DeleteOnlineCount(id bson.ObjectId) (*models.DeleteOnlineCount, error)
	CreateOnlineCount(input *models.CreateOnlineCountInput) (*models.Cap, error)
	UpdateOnlineCount(id bson.ObjectId, input models.UpdateOnlineCountInput) (*models.UpdateOnlineCount, error)
}


func (m *mutation) DeleteOnlineCount(id bson.ObjectId) (*models.DeleteOnlineCount, error) {
	panic("implement me")
}

func (m *mutation) CreateOnlineCount(input *models.CreateOnlineCountInput) (*models.Cap, error) {
	panic("implement me")
}

func (m *mutation) UpdateOnlineCount(id bson.ObjectId, input models.UpdateOnlineCountInput) (*models.UpdateOnlineCount, error) {
	panic("implement me")
}

// newMutationStore ..
func newMutationStore(d Database) Mutation {
	s := &mutation{}
	s.Database = d
	var err error
	s.Loggr, err = logger.New("info", "json")
	if err != nil {
		return nil
	}
	s.Logger = s.Loggr.GetLogger("MONGODB")
	return s
}
