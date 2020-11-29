package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("implement me")
}

func (q QueryResolver) User(ctx context.Context, id bson.ObjectId) (*models.User, error) {
	panic("implement me")
}

