package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) Ticket(ctx context.Context, id bson.ObjectId) (*models.Ticket, error) {
	panic("implement me")
}

func (q QueryResolver) TicketPretty(ctx context.Context, id bson.ObjectId) (*models.TicketPretty, error) {
	panic("implement me")
}


func (q QueryResolver) TicketByYearWeek(ctx context.Context, year int, week int) (*models.Ticket, error) {
	panic("implement me")
}

func (q QueryResolver) ListTickets(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.TicketList, error) {
	panic("implement me")
}
