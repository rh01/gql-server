package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

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
