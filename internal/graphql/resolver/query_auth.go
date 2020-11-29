package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return q.Query.Users()
}

func (q QueryResolver) User(ctx context.Context, id bson.ObjectId) (*models.User, error) {
	return q.Query.User(id)
}

