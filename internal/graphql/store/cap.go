package store

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
	"time"
)

func (m *mutation) DeleteCap(id bson.ObjectId) (*models.DeleteCap, error) {
	m.Logger.Infof("delete cap %s", id.Hex())

	if err := m.GetStore("cap").Remove(id); err != nil {
		m.Logger.Errorf("cannot delete cap %s, err: %s", id.Hex(), err)
		return &models.DeleteCap{Success: false}, fmt.Errorf("cannot delete cap with id: %s, error: %v", id, err)
	}
	m.Logger.Errorf(" delete cap success")
	return &models.DeleteCap{Success: true}, nil
}

func (m *mutation) CreateCap(input *models.CreateCapInput) (*models.Cap, error) {
	year, week := time.Now().ISOWeek()
	q := bson.M{}
	q["product"] = input.Product
	q["year"] = year
	q["week"] = week
	var cap = &models.Cap{}
	if err := m.GetStore("cap").FindOne(q, cap); err != nil {
		if err == mgo.ErrNotFound {
			m.Logger.Infof("can't find cap, I will create new one'")

			cap = &models.Cap{}
			cap.Ctime = time.Now().UTC()
			cap.Utime = time.Now().UTC()
			cap.Product = input.Product
			cap.Desc = input.Desc
			cap.Year, cap.Week = year, week
			// first observe this year and week whether exist or not product record

			if err := m.GetStore("cap").Save(cap); err != nil {
				m.Logger.Errorf("cannot insert cap, error: %v", err)

				return nil, fmt.Errorf("cannot insert cap, error: %v", err)
			}
			return cap, nil
		}
		m.Logger.Errorf("cannot find cap, error: %v", err)
		return nil, fmt.Errorf("cannot find cap, error: %v", err)
	}
	m.Logger.Errorf("create new cap document failed, because cap has exist")
	return nil, fmt.Errorf("create new cap document failed, because cap has exist")
}

func (m *mutation) UpdateCap(id bson.ObjectId, input models.UpdateCapInput) (*models.UpdateCap, error) {
	var update = bson.M{}
	update["$set"] = bson.M{"product": input.Product, "desc": input.Desc}

	if err := m.GetStore("cap").UpdateID(id, update); err != nil {
		m.Logger.Errorf("update cap failed, error: %v", err)
		return &models.UpdateCap{Success: false}, fmt.Errorf("cannot update cap, error: %v", err)
	}
	m.Logger.Infof("update cap success")
	return &models.UpdateCap{Success: true}, nil
}

// queries


// Cap ...
func (q query) Cap(id bson.ObjectId) (*models.Cap, error) {
	var result = models.Cap{}
	if err := q.GetStore("cap").FindByObjectID(id, &result); err != nil {
		return nil, fmt.Errorf("cannot find cap with id: %v, error: %v", id, err)
	}
	return &result, nil
}

// CapByYearWeek ...
func (q query) CapByYearWeek(year int, week int) (*models.Cap, error) {
	var filter = bson.M{}
	var result = &models.Cap{}
	filter["year"] = year
	filter["week"] = week
	if err := q.GetStore("cap").FindOne(filter, result); err != nil {
		return nil, fmt.Errorf("cannot find cap with year: %d week: %d, error: %v", year, week, err)
	}
	return result, nil
}

// ListCaps ...
func (q query) ListCaps(pageIndex int, pageSize int, filter string) (*models.CapList, error) {
	var result = new(models.CapList)
	caps := make([]*models.Cap, 0, pageSize)
	var count int
	var err error
	if count, err = q.GetStore("cap").Count(nil); err != nil {
		return nil, fmt.Errorf("cannot find caps, error: %v", err)
	}

	if err = q.GetStore("cap").FindAllWithPageSize(nil, &caps, pageIndex, pageSize); err != nil {
		return nil, fmt.Errorf("cannot find caps, error: %v", err)
	}
	result.Count = count
	result.Data = caps
	result.Code = 0
	return result, nil
}
