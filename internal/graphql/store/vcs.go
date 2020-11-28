package store

import (
	"report/api/types"

	"gopkg.in/mgo.v2/bson"
)

type vcs struct {
	Store // store
}

//Store related database operations
type Vcs interface {
	Interface

	// VCS Settings
	UpdateVCS(vcs types.VCS) error
}

func newVcsStore(d Database) Vcs {
	s := &vcs{}
	s.Database = d
	s.CollectionName = "vcs"
	return s
}

func (s *vcs) GetVCSByID(id string) (types.VCS, error) {

	var result types.VCS
	err := s.FindByID(id, &result)
	return result, err
}

func (s *vcs) UpdateVCS(vcs types.VCS) error {
	_, err := s.Upsert(bson.M{"_id": vcs.ID}, vcs)
	return err
}
