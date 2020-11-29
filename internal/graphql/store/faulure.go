package store

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
	"time"
)

// Failure
func (q *query) Failure(id bson.ObjectId) (*models.Failure, error) {
	var result = models.Failure{}
	if err := q.GetStore("failure").FindByObjectID(id, &result); err != nil {
		return nil, fmt.Errorf("cannot find failure with id: %v, error: %v", id, err)
	}
	return &result, nil
}

// FailurePretty ...
func (q *query) FailurePretty(id bson.ObjectId) (*models.FailurePretty, error) {
	/**
		Data     [][]*int  `json:"data" bson:"data"`
		Products []*string `json:"products" bson:"products"`
		Levels   []*string `json:"levels" bson:"levels"`
		Year     int       `json:"year" bson:"year"`
		Week     int       `json:"week" bson:"week"
	*/
	// legend.data ['GPU显卡使用率', 'slurm显卡使用率(监控)', 'slurm集群显卡使用率']
	// [53, 60, 91.15],
	var pretty = &models.FailurePretty{}
	var failure = &models.Failure{}

	var filter = bson.M{}
	filter["_id"] = id
	if err := q.GetStore("Failure").FindOne(filter, failure); err != nil {
		return nil, fmt.Errorf("cannot find failure with id %s, error: %v", id.Hex(), err)
	}

	pretty.Year = failure.Year
	pretty.Week = failure.Week

	return pretty, nil
}

// FailureByYearWeek ...
func (q *query) FailureByYearWeek(year int, week int) (*models.Failure, error) {
	var filter = bson.M{}
	var result = &models.Failure{}
	filter["year"] = year
	filter["week"] = week
	if err := q.GetStore("failure").FindOne(filter, result); err != nil {
		return nil, fmt.Errorf("cannot find failure with year: %d week: %d, error: %v", year, week, err)
	}
	return result, nil
}

func (q *query) ListFailures(pageIndex int, pageSize int, filter string) (*models.FailureList, error) {
	var result = new(models.FailureList)
	Failures := make([]*models.Failure, pageSize)
	var count int
	var err error
	if count, err = q.GetStore("failure").Count(nil); err != nil {
		return nil, fmt.Errorf("cannot find Failures, error: %v", err)
	}

	if err = q.GetStore("failure").FindAll(nil, &Failures, pageIndex, pageSize); err != nil {
		return nil, fmt.Errorf("cannot find Failures, error: %v", err)
	}
	result.Count = count
	result.Data = Failures
	result.Code = 0
	return result, nil
}

// mutation
func (m *mutation) DeleteFailure(id bson.ObjectId) (*models.DeleteFailure, error) {
	m.Logger.Infof("delete Failure %s", id.Hex())

	if err := m.GetStore("failure").Remove(id); err != nil {
		m.Logger.Errorf("cannot delete Failure %s, err: %s", id.Hex(), err)
		return &models.DeleteFailure{Success: false}, fmt.Errorf("cannot delete Failure with id: %s, error: %v", id, err)
	}
	m.Logger.Errorf(" delete failure success")
	return &models.DeleteFailure{Success: true}, nil
}

func (m *mutation) CreateFailure(input *models.CreateFailureInput) (*models.Failure, error) {
	year, week := time.Now().ISOWeek()
	q := bson.M{}
	q["year"] = year
	q["week"] = week
	var failure = &models.Failure{}
	if err := m.GetStore("failure").FindOne(q, failure); err != nil {
		if err == mgo.ErrNotFound {
			m.Logger.Infof("can't find Failure, I will create new one'")

			failure = &models.Failure{}

			failure.Year, failure.Week = year, week
			// first observe this year and week whether exist or not product record

			if err := m.GetStore("failure").Save(failure); err != nil {
				m.Logger.Errorf("cannot insert Failure, error: %v", err)

				return nil, fmt.Errorf("cannot insert Failure, error: %v", err)
			}
			return failure, nil
		}
		m.Logger.Errorf("cannot find Failure, error: %v", err)
		return nil, fmt.Errorf("cannot find Failure, error: %v", err)
	}
	m.Logger.Errorf("create new Failure document failed, because Failure has exist")
	return nil, fmt.Errorf("create new Failure document failed, because Failure has exist")
}

func (m *mutation) UpdateFailure(id bson.ObjectId, input models.UpdateFailureInput) (*models.UpdateFailure, error) {
	var update = bson.M{}

	// TODO: update failure
	//update["$set"] = bson.M{"product": input.Failure, "desc": input.Desc}

	if err := m.GetStore("failure").UpdateID(id, update); err != nil {
		m.Logger.Errorf("update Failure failed, error: %v", err)
		return &models.UpdateFailure{Success: false}, fmt.Errorf("cannot update failure, error: %v", err)
	}
	m.Logger.Infof("update Failure success")
	return &models.UpdateFailure{Success: true}, nil
}
