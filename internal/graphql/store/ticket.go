package store

import "report/internal/graphql/models"

func (q query) Tickets(id *string, year *string, week *string, pageIndex int, pageSize int) (*models.TicketRes, error) {
	panic("implement me")
}
