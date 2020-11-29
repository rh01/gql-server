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

	// Cap queries
	Cap(id bson.ObjectId) (*models.Cap, error)
	CapByYearWeek(year int, week int) (*models.Cap, error)
	ListCaps(pageIndex int, pageSize int, filter string) (*models.CapList, error)

	// Ticket
	Ticket(id bson.ObjectId) (*models.Ticket, error)
	TicketPretty(id bson.ObjectId) (*models.TicketPretty, error)
	TicketByYearWeek(year int, week int) (*models.Ticket, error)
	ListTickets(pageIndex int, pageSize int, filter string) (*models.TicketList, error)

	Failure(id bson.ObjectId) (*models.Failure, error)
	FailurePretty(id bson.ObjectId) (*models.FailurePretty, error)
	FailureByYearWeek(year int, week int) (*models.Failure, error)
	ListFailures(pageIndex int, pageSize int, filter string) (*models.FailureList, error)

	Slo(id bson.ObjectId) (*models.Slo, error)
	SloPretty(id bson.ObjectId) (*models.SloPretty, error)
	SloByYearWeek(year int, week int) (*models.Slo, error)
	ListSlos(pageIndex int, pageSize int, filter string) (*models.SloList, error)

	OnlineCount(id bson.ObjectId) (*models.OnlineCount, error)
	OnlineCountByYearWeek(year int, week int) (*models.OnlineCount, error)
	ListOnlineCounts(pageIndex int, pageSize int, filter string) (*models.OnlineCountList, error)
	AllProductOnlineCount(year int, week int) (*models.OnlineCountAllProduct, error)
}

func (q query) Ticket(id bson.ObjectId) (*models.Ticket, error) {
	panic("implement me")
}

func (q query) TicketPretty(id bson.ObjectId) (*models.TicketPretty, error) {
	panic("implement me")
}

func (q query) Failure(id bson.ObjectId) (*models.Failure, error) {
	panic("implement me")
}

func (q query) FailurePretty(id bson.ObjectId) (*models.FailurePretty, error) {
	panic("implement me")
}

func (q query) FailureByYearWeek(year int, week int) (*models.Failure, error) {
	panic("implement me")
}

func (q query) ListFailures(pageIndex int, pageSize int, filter string) (*models.FailureList, error) {
	panic("implement me")
}

func (q query) Slo(id bson.ObjectId) (*models.Slo, error) {
	panic("implement me")
}

func (q query) SloPretty(id bson.ObjectId) (*models.SloPretty, error) {
	panic("implement me")
}

func (q query) SloByYearWeek(year int, week int) (*models.Slo, error) {
	panic("implement me")
}

func (q query) ListSlos(pageIndex int, pageSize int, filter string) (*models.SloList, error) {
	panic("implement me")
}

func (q query) TicketByYearWeek(year int, week int) (*models.Ticket, error) {
	panic("implement me")
}

func (q query) ListTickets(pageIndex int, pageSize int, filter string) (*models.TicketList, error) {
	panic("implement me")
}

func (q query) OnlineCount(id bson.ObjectId) (*models.OnlineCount, error) {
	panic("implement me")
}

func (q query) OnlineCountByYearWeek(year int, week int) (*models.OnlineCount, error) {
	panic("implement me")
}

func (q query) ListOnlineCounts(pageIndex int, pageSize int, filter string) (*models.OnlineCountList, error) {
	panic("implement me")
}

func (q query) AllProductOnlineCount(year int, week int) (*models.OnlineCountAllProduct, error) {
	panic("implement me")
}

// newMutationStore ..
func newQueryStore(d Database) Query {
	s := &query{}
	s.Database = d
	s.CollectionName = ""
	return s
}
