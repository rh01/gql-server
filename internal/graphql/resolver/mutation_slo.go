package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (m MutationResolver) DeleteSlo(ctx context.Context, id bson.ObjectId) (*models.DeleteSlo, error) {
	return m.Mutation.DeleteSlo(id)

}

func (m MutationResolver) CreateSlo(ctx context.Context, input *models.CreateSloInput) (*models.Slo, error) {
	return m.Mutation.CreateSlo(input)
}

func (m MutationResolver) UpdateSlo(ctx context.Context, id bson.ObjectId, input models.UpdateSloInput) (*models.UpdateSlo, error) {
	return m.Mutation.UpdateSlo(id, input)
}