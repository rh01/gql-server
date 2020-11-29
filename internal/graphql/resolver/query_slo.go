package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) Slo(ctx context.Context, id bson.ObjectId) (*models.Slo, error) {
	panic("implement me")
}

func (q QueryResolver) SloPretty(ctx context.Context, id bson.ObjectId) (*models.SloPretty, error) {
	panic("implement me")
}

func (q QueryResolver) SloByYearWeek(ctx context.Context, year int, week int) (*models.Slo, error) {
	panic("implement me")
}

func (q QueryResolver) ListSlos(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.SloList, error) {
	panic("implement me")
}

