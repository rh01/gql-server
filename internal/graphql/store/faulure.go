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
func (q *query) FailurePretty() (*models.FailurePretty, error) {
	year, week := time.Now().ISOWeek()
	/**
	series: [FailureItem!]
	xAxis: [String!]
	year: Int!
	week: Int!

	series: [{name: "P1", type: "bar", data: [ç]},…]
	0: {name: "P1", type: "bar", data: [0, 0, 0, 0, 0, 0]}
	1: {name: "P2", type: "bar", data: [1, 0, 0, 1, 0, 0]}
	2: {name: "P3", type: "bar", data: [0, 0, 0, 0, 0, 0]}
	3: {name: "P4", type: "bar", data: [0, 0, 0, 0, 0, 0]}

	data: ["ke", "ai", "ead", "dict", "zhiyun", "fanyi"]
	*/
	// legend.data ['GPU显卡使用率', 'slurm显卡使用率(监控)', 'slurm集群显卡使用率']
	// [53, 60, 91.15],
	legends := []string{"ke", "ai", "ead", "dict", "zhiyun", "fanyi"}
	levels := []string{"P1", "P2", "P3", "P4"}
	series := []*models.FailureItem{
		{"P1", []int{0, 0, 0, 0, 0, 0}, "bar"},
		{"P2", []int{0, 0, 0, 0, 0, 0}, "bar"},
		{"P3", []int{0, 0, 0, 0, 0, 0}, "bar"},
		{"P4", []int{0, 0, 0, 0, 0, 0}, "bar"},
	}

	var pretty = &models.FailurePretty{}

	var filter = bson.M{}
	filter["year"] = year
	filter["week"] = week
	var failures = []*models.Failure{}
	if err := q.GetStore("failure").FindAll(filter, &failures); err != nil {
		return nil, fmt.Errorf("cannot find failure  error: %v", err)
	}

	fmt.Println(len(failures))
	pretty.Year = year
	pretty.Week = week
	pretty.XAxis = legends
	for k1, level := range levels {
		for k2, product := range legends {
			for _, fail := range failures {
				if fail.Level == level && fail.Product == product {
					series[k1].Data[k2] += 1
				}
			}
		}
	}
	pretty.Series = series

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

	if err = q.GetStore("failure").FindAllWithPageSize(nil, &Failures, pageIndex, pageSize); err != nil {
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
	year, week := input.StartTime.ISOWeek()
	q := bson.M{}
	q["year"] = year
	q["week"] = week
	q["product"] = input.Product
	q["desc"] = input.Desc
	var failure = &models.Failure{}
	// first observe this year and week whether exist or not product record
	if err := m.GetStore("failure").FindOne(q, failure); err != nil {
		if err == mgo.ErrNotFound {
			m.Logger.Infof("can't find Failure, I will create new one'")

			failure = &models.Failure{}
			failure.Year, failure.Week = year, week
			failure.Product = input.Product
			failure.Desc = input.Desc
			failure.Ctime = time.Now().UTC()
			failure.Utime = time.Now().UTC()
			failure.Level = input.Level
			failure.StartTime = input.StartTime
			failure.EndTime = input.EndTime
			failure.Duration = input.EndTime.Sub(input.StartTime).Seconds()
			failure.Recorder = input.Recorder
			failure.Title = input.Title

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
