package service

import (
	"L2.18-calendar/internal/repository"
	"time"
)

type Repo interface {
	Create(e *repository.Event) error
	GetForMonth(userID int, month time.Month, year int) ([]repository.Event, error)
	GetForDay(userID int, data time.Time) ([]repository.Event, error)
	GetForWeek(userID int, date time.Time) ([]repository.Event, error)
}
