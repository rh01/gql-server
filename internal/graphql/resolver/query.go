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

func (q QueryResolver) Ticket(ctx context.Context, id bson.ObjectId) (*models.Ticket, error) {
	panic("implement me")
}

func (q QueryResolver) TicketPretty(ctx context.Context, id bson.ObjectId) (*models.TicketPretty, error) {
	panic("implement me")
}

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

func (q QueryResolver) TicketByYearWeek(ctx context.Context, year int, week int) (*models.Ticket, error) {
	panic("implement me")
}

func (q QueryResolver) ListTickets(ctx context.Context, pageIndex int, pageSize int, filter string) (*models.TicketList, error) {
	panic("implement me")
}

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
