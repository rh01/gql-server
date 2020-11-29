package resolver

import (
	"github.com/sirupsen/logrus"
	"report/internal/graphql/store"
)

type MutationResolver struct {
	store.Report
	Logger *logrus.Entry
}
