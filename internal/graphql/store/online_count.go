package store

import (
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (m *mutation) DeleteOnlineCount(id bson.ObjectId) (*models.DeleteOnlineCount, error) {
	panic("implement me")
}

func (m *mutation) CreateOnlineCount(input *models.CreateOnlineCountInput) (*models.OnlineCount, error) {
	panic("implement me")
}

func (m *mutation) UpdateOnlineCount(id bson.ObjectId, input models.UpdateOnlineCountInput) (*models.UpdateOnlineCount, error) {
	panic("implement me")
}
