package store

import (
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

type query struct {
	Store
}

// Query ...
type Query interface {
	Interface

	Cap(id bson.ObjectId) (*models.Cap, error)
	CapByYearWeek(year string, week string) (*models.Cap, error)
	ListCaps(pageIndex int, pageSize int, filter string) (*models.CapList, error)
	Tickets(id *string, year *string, week *string, pageIndex int, pageSize int) (*models.TicketRes, error)
}

// newMutationStore ..
func newQueryStore(d Database) Query {
	s := &query{}
	s.Database = d
	s.CollectionName = ""
	return s
}
