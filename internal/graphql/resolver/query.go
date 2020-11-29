package resolver

import (
	"context"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
	"report/internal/graphql/store"
)

type QueryResolver struct {
	store.Report

	Logger *logrus.Entry
}

func (q QueryResolver) Cap(ctx context.Context, id bson.ObjectId) (*models.Cap, error) {
	q.Logger.Debugf("fetching... cap by id: %s", id.Hex())
	cap, err := q.Query.Cap(id)
	if err != nil {
		q.Logger.Errorf("fetch cap failed, err: %v", err)
	}
	return cap, err
}

func (q QueryResolver) CapByYearWeek(ctx context.Context, year string, week string) (*models.Cap, error) {
	panic("implement me")
}
