package handler

import "L2.18-calendar/internal/service"

type Service interface {
	Create(event *service.Event) error
	GetForMonth(userID int, month string) ([]service.Event, error)
	GetForDay(userID int, date string) ([]service.Event, error)
	GetForWeek(userID int, data string) ([]service.Event, error)
}
