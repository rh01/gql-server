package store

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
	"time"
)

func (m *mutation) DeleteCap(id string) (*models.DeleteCap, error) {
	if err := m.GetCapStore("cap").Remove(id); err != nil {
		return &models.DeleteCap{Success: false}, fmt.Errorf("cannot find cap with id: %s, error: %v", id, err)
	}
	return &models.DeleteCap{Success: true}, nil
}

func (m *mutation) CreateCap(input *models.CreateCapInput) (*models.Cap, error) {
	year, week := time.Now().ISOWeek()
	var cap = &models.Cap{}

	cap.Ctime = time.Now().UTC().Add(8 * time.Hour)
	cap.Utime = time.Now().UTC().Add(8 * time.Hour)
	cap.Product = input.Product
	cap.Desc = input.Desc
	cap.Year, cap.Week = year, week

	if err := m.GetCapStore("cap").Save(cap); err != nil {
		return nil, fmt.Errorf("cannot insert cap, error: %v", err)
	}
	return cap, nil
}

func (m *mutation) UpdateCap(id string, input models.UpdateCapInput) (*models.UpdateCap, error) {
	var update = bson.M{}
	update["$set"] = bson.M{"product": input.Product, "desc": input.Desc}

	if err := m.GetCapStore("cap").UpdateID(id, update); err != nil {
		return &models.UpdateCap{Success: false}, fmt.Errorf("cannot update cap, error: %v", err)
	}
	return &models.UpdateCap{Success: true}, nil
}

func (q query) Cap(id bson.ObjectId) (*models.Cap, error) {
	var result= models.Cap{}
	if err := q.GetCapStore("cap").FindByObjectID(id, &result); err != nil {
		return nil, fmt.Errorf("cannot find cap with id: %v, error: %v", id, err)
	}
	return &result, nil
}

func (q query) CapByYearWeek(year string, week string) (*models.Cap, error) {
	var filter = bson.M{}
	var result *models.Cap
	filter["year"] = year
	filter["week"] = week
	if err := q.GetCapStore("cap").FindOne(filter, result); err != nil {
		return nil, fmt.Errorf("cannot find cap with year: %s week: %s, error: %v", year, week, err)
	}
	return result, nil
}

func (q query) ListCaps(pageIndex int, pageSize int, filter string) (*models.CapList, error) {
	var result *models.CapList
	if err := q.GetCapStore("cap").FindAll(nil, result, pageIndex, pageSize); err != nil {
		return nil, fmt.Errorf("cannot find caps, error: %v", err)
	}
	return result, nil
}
