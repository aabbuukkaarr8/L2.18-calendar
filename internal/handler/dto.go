package handler

import (
	"L2.18-calendar/internal/service"
	"time"
)

type Response struct {
	UserID int    `json:"user_id" validate:"required"`
	Date   string `json:"date" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

type Event struct {
	UserID int
	Date   time.Time
	Text   string
}

func (r *Event) ToSrv() service.Event {
	return service.Event{
		UserID: r.UserID,
		Date:   r.Date,
		Text:   r.Text,
	}
}

func (r *Event) FromSrv(event service.Event) {
	r.UserID = event.UserID
	r.Date = event.Date
	r.Text = event.Text
}
