//go:generate go run ../../../cmd/gen/main.go
package resolver

import (
	"github.com/sirupsen/logrus"
	"report/internal/graphql/generated"
	"report/internal/graphql/store"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	store.Report
	Logger *logrus.Entry
}

func NewResolver(s store.Report, l *logrus.Entry) *Resolver {
	return &Resolver{Report: s, Logger: l}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &MutationResolver{
		Report: r.Report,
		Logger: r.Logger,
	}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &QueryResolver{
		Report: r.Report,
		Logger: r.Logger,
	}
}
