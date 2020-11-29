package store

import "report/internal/graphql/models"

type mutation struct {
	Store
}

// Cap ...
type Mutation interface {
	Interface

	DeleteCap(id string) (*models.DeleteCap, error)
	CreateCap(input *models.CreateCapInput) (*models.Cap, error)
	UpdateCap(id string, input models.UpdateCapInput) (*models.UpdateCap, error)
}

// newMutationStore ..
func newMutationStore(d Database) Mutation {
	s := &mutation{}
	s.Database = d
	s.CollectionName = ""
	return s
}
