package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) OnlineCount(ctx context.Context, id bson.ObjectId) (*models.OnlineCount, error) {
	return q.Query.OnlineCount(id)
}

func (q QueryResolver) OnlineCountByYearWeek(ctx context.Context, year int, week int) (*models.OnlineCount, error) {
	return q.Query.OnlineCountByYearWeek(year, week)
}

func (q QueryResolver) ListOnlineCounts(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.OnlineCountList, error) {
	return q.Query.ListOnlineCounts(pageIndex, pageSize, filter)
}

func (q QueryResolver) AllProductOnlineCount(ctx context.Context, year int, week int) (*models.OnlineCountAllProduct, error) {
	return q.Query.AllProductOnlineCount(year, week)
}
