package store

import (
	"encoding/json"
	"report/internal/pkg/utils"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	recordNotFound = "record not found"
)

// Interface ..
// Abstract database interactions.
type Interface interface {
	Execute(handleFunc func(c *mgo.Collection))
	Save(model interface{}) error
	Upsert(selector interface{}, model interface{}) (*mgo.ChangeInfo, error)
	Update(selector interface{}, model interface{}) error
	UpdateID(id interface{}, model interface{}) error
	FindAll(query interface{}, model interface{}) error
	FindOne(query interface{}, model interface{}) error
	FindByID(id string, model interface{}) error
	FindByObjectID(id bson.ObjectId, model interface{}) error
	Exist(selector interface{}) (bool, error)
	Remove(id interface{}) error
	RemoveBySelector(selector interface{}) error
	RemoveMultiple(ids []interface{}) error
	GetSession() *mgo.Session
}

// Store ..
// A base datasource that performs actualy sql interactions.
type Store struct {
	Database       Database
	CollectionName string
}

// GetSession : Get a session from mongodb
func (s *Store) GetSession() *mgo.Session {
	return s.Database.Session
}

// Execute given func with a active session against the database
func (s *Store) Execute(handle func(c *mgo.Collection)) {

	ses := s.Database.Session.Copy()
	defer ses.Close()
	handle(ses.DB(s.Database.Name).C(s.CollectionName))
	return
}

// Exist Checks whether the given document exist in a collection
func (s *Store) Exist(selector interface{}) (bool, error) {

	var count int
	var err error
	s.Execute(func(c *mgo.Collection) {
		count, err = c.Find(selector).Count()
	})

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

//Save Insert operation for a model on a collection
func (s *Store) Save(model interface{}) error {

	var err error
	s.Execute(func(c *mgo.Collection) {
		err = c.Insert(model)
	})
	return err
}

// Upsert for a model on a collection, based on a selector
func (s *Store) Upsert(selector interface{}, model interface{}) (*mgo.ChangeInfo, error) {

	var info *mgo.ChangeInfo
	var err error
	s.Execute(func(c *mgo.Collection) {
		info, err = c.Upsert(selector, model)
	})
	return info, err
}

// Update the collection based on the selector (query fields)
func (s *Store) Update(selector interface{}, model interface{}) error {

	var err error
	s.Execute(func(c *mgo.Collection) {
		err = c.Update(selector, model)
	})
	return err
}

//UpdateID Update the collection based on the id (primary key)
func (s *Store) UpdateID(id interface{}, model interface{}) error {

	var err error
	s.Execute(func(c *mgo.Collection) {
		err = c.UpdateId(id, model)
	})
	return err
}

// FindAll the document matches the query on a collection.
func (s *Store) FindAll(query interface{}, model interface{}) error {

	var err error
	s.Execute(func(c *mgo.Collection) {
		err = c.Find(query).All(model)
	})
	return err
}

// FindOne document matches the query on a collection
func (s *Store) FindOne(query interface{}, model interface{}) error {

	var err error
	s.Execute(func(c *mgo.Collection) {
		err = c.Find(query).One(model)
	})
	return err
}

// FindByID matches the document by _id
func (s *Store) FindByID(id string, model interface{}) error {
	return s.FindByObjectID(bson.ObjectIdHex(id), model)
}

//FindByObjectID FindByID matches the document by _id
func (s *Store) FindByObjectID(id bson.ObjectId, model interface{}) error {
	return s.FindOne(bson.M{"_id": id}, model)
}

// Remove a document based on id
func (s *Store) Remove(id interface{}) error {

	var err error
	s.Execute(func(c *mgo.Collection) {
		err = c.RemoveId(id)
	})
	return err
}

//RemoveBySelector Remove a document based on a selector
func (s *Store) RemoveBySelector(selector interface{}) error {

	var err error
	s.Execute(func(c *mgo.Collection) {
		err = c.Remove(selector)
	})
	return err
}

// RemoveMultiple document based on gived ids
func (s *Store) RemoveMultiple(ids []interface{}) error {

	var err error
	s.Execute(func(c *mgo.Collection) {
		err = c.Remove(bson.M{"_id": bson.M{"$in": ids}})
	})
	return err
}

func encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func decode(data []byte, out interface{}) error {
	return json.Unmarshal(data, out)
}

// NewID ..
// Creates a new UUID and returns string
func NewID() string {
	return utils.NewUUID()
}
