package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

func (q QueryResolver) Failure(ctx context.Context, id bson.ObjectId) (*models.Failure, error) {
	return q.Query.Failure(id)
}

func (q QueryResolver) FailurePretty(ctx context.Context, id bson.ObjectId) (*models.FailurePretty, error) {
	return q.Query.FailurePretty(id)
}

func (q QueryResolver) FailureByYearWeek(ctx context.Context, year int, week int) (*models.Failure, error) {
	return q.Query.FailureByYearWeek(year, week)
}

func (q QueryResolver) ListFailures(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.FailureList, error) {
	return q.Query.ListFailures(pageIndex, pageSize, filter)
}
