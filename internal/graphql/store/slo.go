package store

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math"
	"reflect"
	"report/internal/graphql/models"
	"strconv"
	"time"
)

var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf("")

// Slo
func (q *query) Slo(id bson.ObjectId) (*models.Slo, error) {
	var result = models.Slo{}
	if err := q.GetStore("slo").FindByObjectID(id, &result); err != nil {
		return nil, fmt.Errorf("cannot find slo with id: %v, error: %v", id, err)
	}
	return &result, nil
}

// SloPretty ...
func (q *query) SloPretty(id bson.ObjectId) (*models.SloPretty, error) {
	/**
	Product string     `json:"product" bson:"product"`
	Sloes   []float64 `json:"sloes" bson:"sloes"`
	Names   []string  `json:"names" bson:"names"`
	Year    int        `json:"year" bson:"year"`
	Week    int        `json:"week" bson:"week"
	*/
	// legend.data ['GPU显卡使用率', 'slurm显卡使用率(监控)', 'slurm集群显卡使用率']
	// [53, 60, 91.15],
	var pretty = &models.SloPretty{}
	var slo = &models.Slo{}

	var filter = bson.M{}
	filter["_id"] = id
	if err := q.GetStore("Slo").FindOne(filter, slo); err != nil {
		return nil, fmt.Errorf("cannot find slo with id %s, error: %v", id.Hex(), err)
	}
	pretty.Legend = make([]string, 0)
	pretty.Sloes = make([]float64, 0)
	for metric, value := range slo.Metrics {
		pretty.Legend = append(pretty.Legend, metric)

		m, err := getFloat(value)
		if err != nil {
			return nil, fmt.Errorf("cannot convert value to float64, error: %v", err)
		}
		pretty.Sloes = append(pretty.Sloes, m)
	}

	pretty.Product = slo.Product
	pretty.Year = slo.Year
	pretty.Week = slo.Week

	return pretty, nil
}

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.ParseFloat(s, 64)
		} else {
			return math.NaN(), fmt.Errorf("Can't convert %v to float64", v.Type())
		}
	}
}

// SloByYearWeek ...
func (q *query) SloByYearWeek(year int, week int) (*models.Slo, error) {
	var filter = bson.M{}
	var result = &models.Slo{}
	filter["year"] = year
	filter["week"] = week
	if err := q.GetStore("slo").FindOne(filter, result); err != nil {
		return nil, fmt.Errorf("cannot find slo with year: %d week: %d, error: %v", year, week, err)
	}
	return result, nil
}

func (q *query) ListSlos(pageIndex int, pageSize int, filter string) (*models.SloList, error) {
	var result = new(models.SloList)
	Slos := make([]*models.Slo, pageSize)
	var count int
	var err error
	if count, err = q.GetStore("slo").Count(nil); err != nil {
		return nil, fmt.Errorf("cannot find Slos, error: %v", err)
	}

	if err = q.GetStore("slo").FindAll(nil, &Slos, pageIndex, pageSize); err != nil {
		return nil, fmt.Errorf("cannot find Slos, error: %v", err)
	}
	result.Count = count
	result.Data = Slos
	result.Code = 0
	return result, nil
}

// mutation
func (m *mutation) DeleteSlo(id bson.ObjectId) (*models.DeleteSlo, error) {
	m.Logger.Infof("delete Slo %s", id.Hex())

	if err := m.GetStore("slo").Remove(id); err != nil {
		m.Logger.Errorf("cannot delete Slo %s, err: %s", id.Hex(), err)
		return &models.DeleteSlo{Success: false}, fmt.Errorf("cannot delete Slo with id: %s, error: %v", id, err)
	}
	m.Logger.Errorf(" delete slo success")
	return &models.DeleteSlo{Success: true}, nil
}

func (m *mutation) CreateSlo(input *models.CreateSloInput) (*models.Slo, error) {
	year, week := time.Now().ISOWeek()
	q := bson.M{}
	q["year"] = year
	q["week"] = week
	var Slo = &models.Slo{}
	if err := m.GetStore("slo").FindOne(q, Slo); err != nil {
		if err == mgo.ErrNotFound {
			m.Logger.Infof("can't find Slo, I will create new one'")

			Slo = &models.Slo{}
			Slo.Ctime = time.Now().UTC()
			Slo.Utime = time.Now().UTC()

			Slo.Year, Slo.Week = year, week
			// first observe this year and week whether exist or not product record

			if err := m.GetStore("slo").Save(Slo); err != nil {
				m.Logger.Errorf("cannot insert Slo, error: %v", err)

				return nil, fmt.Errorf("cannot insert Slo, error: %v", err)
			}
			return Slo, nil
		}
		m.Logger.Errorf("cannot find Slo, error: %v", err)
		return nil, fmt.Errorf("cannot find Slo, error: %v", err)
	}
	m.Logger.Errorf("create new Slo document failed, because Slo has exist")
	return nil, fmt.Errorf("create new Slo document failed, because Slo has exist")
}

func (m *mutation) UpdateSlo(id bson.ObjectId, input models.UpdateSloInput) (*models.UpdateSlo, error) {
	var update = bson.M{}
	//TODO: update
	//update["$set"] = bson.M{"product": input.Slo, "desc": input.Desc}

	if err := m.GetStore("slo").UpdateID(id, update); err != nil {
		m.Logger.Errorf("update Slo failed, error: %v", err)
		return &models.UpdateSlo{Success: false}, fmt.Errorf("cannot update slo, error: %v", err)
	}
	m.Logger.Infof("update Slo success")
	return &models.UpdateSlo{Success: true}, nil
}
