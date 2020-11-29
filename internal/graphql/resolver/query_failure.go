package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) Failure(ctx context.Context, id bson.ObjectId) (*models.Failure, error) {
	panic("implement me")
}

func (q QueryResolver) FailurePretty(ctx context.Context, id bson.ObjectId) (*models.FailurePretty, error) {
	panic("implement me")
}

func (q QueryResolver) FailureByYearWeek(ctx context.Context, year int, week int) (*models.Failure, error) {
	panic("implement me")
}

func (q QueryResolver) ListFailures(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.FailureList, error) {
	panic("implement me")
}

