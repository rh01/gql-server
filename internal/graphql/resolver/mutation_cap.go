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

func (m MutationResolver) CreateTicket(ctx context.Context, input *models.CreateTicketInput) (*models.Ticket, error) {
	panic("implement me")
}

func (m MutationResolver) DeleteSlo(ctx context.Context, id bson.ObjectId) (*models.DeleteSlo, error) {
	panic("implement me")
}

func (m MutationResolver) CreateSlo(ctx context.Context, input *models.CreateSloInput) (*models.Slo, error) {
	panic("implement me")
}

func (m MutationResolver) UpdateSlo(ctx context.Context, id bson.ObjectId, input models.UpdateSloInput) (*models.UpdateSlo, error) {
	panic("implement me")
}

func (m MutationResolver) DeleteFailure(ctx context.Context, id bson.ObjectId) (*models.DeleteFailure, error) {
	panic("implement me")
}

func (m MutationResolver) CreateFailure(ctx context.Context, input *models.CreateFailureInput) (*models.Failure, error) {
	panic("implement me")
}

func (m MutationResolver) UpdateFailure(ctx context.Context, id bson.ObjectId, input models.UpdateFailureInput) (*models.UpdateFailure, error) {
	panic("implement me")
}

func (m MutationResolver) DeleteTicket(ctx context.Context, id bson.ObjectId) (*models.DeleteTicket, error) {
	panic("implement me")
}

func (m MutationResolver) UpdateTicket(ctx context.Context, id bson.ObjectId, input models.UpdateTicketInput) (*models.UpdateTicket, error) {
	panic("implement me")
}

func (m MutationResolver) DeleteOnlineCount(ctx context.Context, id bson.ObjectId) (*models.DeleteOnlineCount, error) {
	panic("implement me")
}

func (m MutationResolver) CreateOnlineCount(ctx context.Context, input *models.CreateOnlineCountInput) (*models.Cap, error) {
	panic("implement me")
}

func (m MutationResolver) UpdateOnlineCount(ctx context.Context, id bson.ObjectId, input models.UpdateOnlineCountInput) (*models.UpdateOnlineCount, error) {
	panic("implement me")
}

// Cap mutations
func (m MutationResolver) DeleteCap(ctx context.Context, id bson.ObjectId) (*models.DeleteCap, error) {
	return m.Mutation.DeleteCap(id)
}

func (m MutationResolver) CreateCap(ctx context.Context, input *models.CreateCapInput) (*models.Cap, error) {
	return m.Mutation.CreateCap(input)
}

func (m MutationResolver) UpdateCap(ctx context.Context, id bson.ObjectId, input models.UpdateCapInput) (*models.UpdateCap, error) {
	return m.Mutation.UpdateCap(id, input)
}
