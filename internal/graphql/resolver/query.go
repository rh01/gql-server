package resolver

import (
	"github.com/sirupsen/logrus"
	"report/internal/graphql/store"
)

type QueryResolver struct {
	store.Report

	Logger *logrus.Entry
}
