package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) OnlineCount(ctx context.Context, id bson.ObjectId) (*models.OnlineCount, error) {
	panic("implement me")
}

func (q QueryResolver) OnlineCountByYearWeek(ctx context.Context, year int, week int) (*models.OnlineCount, error) {
	panic("implement me")
}

func (q QueryResolver) ListOnlineCounts(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.OnlineCountList, error) {
	panic("implement me")
}

func (q QueryResolver) AllProductOnlineCount(ctx context.Context, year int, week int) (*models.OnlineCountAllProduct, error) {
	panic("implement me")
}
