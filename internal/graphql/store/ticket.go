package store

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"report/internal/graphql/models"
	"time"
)

// Ticket
func (q *query) Ticket(id bson.ObjectId) (*models.Ticket, error) {
	var result = models.Ticket{}
	if err := q.GetStore("ticket").FindByObjectID(id, &result); err != nil {
		return nil, fmt.Errorf("cannot find cap with id: %v, error: %v", id, err)
	}
	return &result, nil
}

// TicketPretty ...
func (q *query) TicketPretty(id bson.ObjectId) (*models.TicketPretty, error) {
	/**
	Orders  []*int    `json:"orders" bson:"orders"`
	Legend []*string `json:"legend" bson:"legend"`
	Week    int       `json:"week" bson:"week"`
	Year    int       `json:"year" bson:"year"`
	*/
	// legend.data ["常规", "非常规", "常规 2h 内", "常规 2h 以上", "非常规 2h 内", "非常规 2h 以上"]
	// data: [{value: 69, name: "常规"}, }
	var pretty = &models.TicketPretty{}
	var ticket = &models.Ticket{}

	var filter = bson.M{}
	filter["_id"] = id
	if err := q.GetStore("ticket").FindOne(filter, ticket); err != nil {
		return nil, fmt.Errorf("cannot find cap with id %s, error: %v", id.Hex(), err)
	}
	pretty.Legend = []string{"常规", "非常规", "常规 2h 内", "常规 2h 以上", "非常规 2h 内", "非常规 2h 以上"}

	// build orders
	orders := ticket.Order
	chartData := make([]*models.TicketEchartData, 0)
	chartData = append(chartData, &models.TicketEchartData{
		Name:  "常规",
		Value: orders.Normal,
	})

	chartData = append(chartData, &models.TicketEchartData{
		Name:  "非常规",
		Value: orders.Abnormal,
	})

	chartData = append(chartData, &models.TicketEchartData{
		Name:  "常规 2h 内",
		Value: orders.NormalLt2h,
	})

	chartData = append(chartData, &models.TicketEchartData{
		Name:  "常规 2h 以上",
		Value: orders.NormalGt2h,
	})

	chartData = append(chartData, &models.TicketEchartData{
		Name:  "非常规 2h 内",
		Value: orders.AbnormalLt2h,
	})

	chartData = append(chartData, &models.TicketEchartData{
		Name:  "非常规 2h 以上",
		Value: orders.AbnormalGt2h,
	})
	pretty.Orders = chartData
	pretty.Year = ticket.Year
	pretty.Week = ticket.Week

	return pretty, nil
}

// TicketByYearWeek ...
func (q *query) TicketByYearWeek(year int, week int) (*models.Ticket, error) {
	var filter = bson.M{}
	var result = &models.Ticket{}
	filter["year"] = year
	filter["week"] = week
	if err := q.GetStore("ticket").FindOne(filter, result); err != nil {
		return nil, fmt.Errorf("cannot find cap with year: %d week: %d, error: %v", year, week, err)
	}
	return result, nil
}

func (q *query) ListTickets(pageIndex int, pageSize int, filter string) (*models.TicketList, error) {
	var result = new(models.TicketList)
	tickets := make([]*models.Ticket, pageSize)
	var count int
	var err error
	if count, err = q.GetStore("ticket").Count(nil); err != nil {
		return nil, fmt.Errorf("cannot find tickets, error: %v", err)
	}

	if err = q.GetStore("ticket").FindAllWithPageSize(nil, &tickets, pageIndex, pageSize); err != nil {
		return nil, fmt.Errorf("cannot find tickets, error: %v", err)
	}
	result.Count = count
	result.Data = tickets
	result.Code = 0
	return result, nil
}

// mutation
func (m *mutation) DeleteTicket(id bson.ObjectId) (*models.DeleteTicket, error) {
	m.Logger.Infof("delete ticket %s", id.Hex())

	if err := m.GetStore("ticket").Remove(id); err != nil {
		m.Logger.Errorf("cannot delete ticket %s, err: %s", id.Hex(), err)
		return &models.DeleteTicket{Success: false}, fmt.Errorf("cannot delete ticket with id: %s, error: %v", id, err)
	}
	m.Logger.Errorf(" delete cap success")
	return &models.DeleteTicket{Success: true}, nil
}

func (m *mutation) CreateTicket(input *models.CreateTicketInput) (*models.Ticket, error) {
	year, week := time.Now().ISOWeek()
	q := bson.M{}
	q["year"] = year
	q["week"] = week
	var ticket = &models.Ticket{}
	if err := m.GetStore("ticket").FindOne(q, ticket); err != nil {
		if err == mgo.ErrNotFound {
			m.Logger.Infof("can't find ticket, I will create new one'")

			ticket = &models.Ticket{}
			ticket.Ctime = time.Now().UTC()
			ticket.Utime = time.Now().UTC()

			ticket.Order = new(models.Order)
			ticket.Order.NormalLt2h = input.NormalLt2h
			ticket.Order.NormalGt2h = input.NormalGt2h
			ticket.Order.AbnormalLt2h = input.AbnormalLt2h
			ticket.Order.AbnormalGt2h = input.AbnormalGt2h
			ticket.Order.Normal = input.NormalLt2h + input.NormalGt2h
			ticket.Order.Abnormal = input.AbnormalLt2h + input.AbnormalGt2h

			ticket.Year, ticket.Week = year, week
			// first observe this year and week whether exist or not product record

			if err := m.GetStore("ticket").Save(ticket); err != nil {
				m.Logger.Errorf("cannot insert ticket, error: %v", err)

				return nil, fmt.Errorf("cannot insert ticket, error: %v", err)
			}
			return ticket, nil
		}
		m.Logger.Errorf("cannot find ticket, error: %v", err)
		return nil, fmt.Errorf("cannot find ticket, error: %v", err)
	}
	m.Logger.Errorf("create new ticket document failed, because ticket has exist")
	return nil, fmt.Errorf("create new ticket document failed, because ticket has exist")
}

func (m *mutation) UpdateTicket(id bson.ObjectId, input models.UpdateTicketInput) (*models.UpdateTicket, error) {
	var update = bson.M{}
	//update["$set"] = bson.M{"product": input.ticket, "desc": input.Desc}

	if err := m.GetStore("ticket").UpdateID(id, update); err != nil {
		m.Logger.Errorf("update ticket failed, error: %v", err)
		return &models.UpdateTicket{Success: false}, fmt.Errorf("cannot update cap, error: %v", err)
	}
	m.Logger.Infof("update ticket success")
	return &models.UpdateTicket{Success: true}, nil
}
