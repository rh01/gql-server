package resolver

import (
	"context"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
	"report/internal/graphql/store"
)

type MutationResolver struct {
	store.Report
	Logger *logrus.Entry
}

func (m MutationResolver) DeleteCap(ctx context.Context, id bson.ObjectId) (*models.DeleteCap, error) {
	panic("implement me")
}

func (m MutationResolver) CreateCap(ctx context.Context, input *models.CreateCapInput) (*models.Cap, error) {
	return m.Mutation.CreateCap(input)
}

func (m MutationResolver) UpdateCap(ctx context.Context, id bson.ObjectId, input models.UpdateCapInput) (*models.UpdateCap, error) {
	panic("implement me")
}

func (m MutationResolver) CreateTickets(ctx context.Context, input models.NewTicket) (bool, error) {
	panic("implement me")
}

func (m MutationResolver) DeleteTickets(ctx context.Context, id *bson.ObjectId, week *string, year *string) (int, error) {
	panic("implement me")
}

func (m MutationResolver) UpdateTicket(ctx context.Context, id *bson.ObjectId, input models.NewTicket) (bool, error) {
	panic("implement me")
}