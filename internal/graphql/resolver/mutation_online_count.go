package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (m MutationResolver) DeleteOnlineCount(ctx context.Context, id bson.ObjectId) (*models.DeleteOnlineCount, error) {
	return m.Mutation.DeleteOnlineCount(id)
}

func (m MutationResolver) CreateOnlineCount(ctx context.Context, input *models.CreateOnlineCountInput) (*models.OnlineCount, error) {
	return m.Mutation.CreateOnlineCount(input)
}

func (m MutationResolver) UpdateOnlineCount(ctx context.Context, id bson.ObjectId, input models.UpdateOnlineCountInput) (*models.UpdateOnlineCount, error) {
	return m.Mutation.UpdateOnlineCount(id, input)
}
