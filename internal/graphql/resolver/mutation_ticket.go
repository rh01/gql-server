package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (m MutationResolver) CreateTicket(ctx context.Context, input *models.CreateTicketInput) (*models.Ticket, error) {
	return m.Mutation.CreateTicket(input)
}

func (m MutationResolver) DeleteTicket(ctx context.Context, id bson.ObjectId) (*models.DeleteTicket, error) {
	return m.Mutation.DeleteTicket(id)
}

func (m MutationResolver) UpdateTicket(ctx context.Context, id bson.ObjectId, input models.UpdateTicketInput) (*models.UpdateTicket, error) {
	return m.Mutation.UpdateTicket(id, input)
}
