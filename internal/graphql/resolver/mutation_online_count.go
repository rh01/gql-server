package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (m MutationResolver) DeleteOnlineCount(ctx context.Context, id bson.ObjectId) (*models.DeleteOnlineCount, error) {
	panic("implement me")
}

func (m MutationResolver) CreateOnlineCount(ctx context.Context, input *models.CreateOnlineCountInput) (*models.OnlineCount, error) {
	panic("implement me")
}

func (m MutationResolver) UpdateOnlineCount(ctx context.Context, id bson.ObjectId, input models.UpdateOnlineCountInput) (*models.UpdateOnlineCount, error) {
	panic("implement me")
}
