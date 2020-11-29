package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (m MutationResolver) DeleteFailure(ctx context.Context, id bson.ObjectId) (*models.DeleteFailure, error) {
	return m.Mutation.DeleteFailure(id)
}

func (m MutationResolver) CreateFailure(ctx context.Context, input *models.CreateFailureInput) (*models.Failure, error) {
	return m.Mutation.CreateFailure(input)
}

func (m MutationResolver) UpdateFailure(ctx context.Context, id bson.ObjectId, input models.UpdateFailureInput) (*models.UpdateFailure, error) {
	return m.Mutation.UpdateFailure(id, input)
}
