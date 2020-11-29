package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (m MutationResolver) DeleteFailure(ctx context.Context, id bson.ObjectId) (*models.DeleteFailure, error) {
	panic("implement me")
}

func (m MutationResolver) CreateFailure(ctx context.Context, input *models.CreateFailureInput) (*models.Failure, error) {
	panic("implement me")
}

func (m MutationResolver) UpdateFailure(ctx context.Context, id bson.ObjectId, input models.UpdateFailureInput) (*models.UpdateFailure, error) {
	panic("implement me")
}
