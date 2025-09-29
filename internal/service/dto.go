package service

import (
	"L2.18-calendar/internal/repository"
	"time"
)

type Event struct {
	UserID int
	Date   time.Time
	Text   string
}

func (e *Event) FillFromDB(dbe *repository.Event) {
	e.UserID = dbe.UserID
	e.Text = dbe.Text
	e.Date = dbe.Date
}
