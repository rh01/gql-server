package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) Slo(ctx context.Context, id bson.ObjectId) (*models.Slo, error) {
	return q.Query.Slo(id)
}

func (q QueryResolver) SloPretty(ctx context.Context, id bson.ObjectId) (*models.SloPretty, error) {
	return q.Query.SloPretty(id)
}

func (q QueryResolver) SloByYearWeek(ctx context.Context, year int, week int) (*models.Slo, error) {
	return q.Query.SloByYearWeek(year, week)
}

func (q QueryResolver) ListSlos(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.SloList, error) {
	return q.Query.ListSlos(pageIndex, pageSize, filter)
}

