package resolver

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
)

// ListCaps queries ...
func (q QueryResolver) ListCaps(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.CapList, error) {
	q.Logger.Debugf("batch fetching... caps")
	capList, err := q.Query.ListCaps(pageIndex, pageSize, filter)
	if err != nil {
		q.Logger.Errorf("cfetch caps failed, err: %v", err)
	}
	return capList, err
}

// Cap queries ...
func (q QueryResolver) Cap(ctx context.Context, id bson.ObjectId) (*models.Cap, error) {
	q.Logger.Debugf("fetching... cap by id: %s", id.Hex())
	cap, err := q.Query.Cap(id)
	if err != nil {
		q.Logger.Errorf("fetch cap failed, err: %v", err)
	}
	return cap, err
}

// CapByYearWeek queries ...
func (q QueryResolver) CapByYearWeek(ctx context.Context, year int, week int) (*models.Cap, error) {
	q.Logger.Debugf("fetching... cap by year: %d, week: %d", year, week)
	cap, err := q.Query.CapByYearWeek(year, week)
	if err != nil {
		q.Logger.Errorf("fetch cap failed  by year: %d, week: %d, err: %v", year, week, err)
	}
	return cap, err
}

