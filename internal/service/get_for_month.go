package service

import (
	"time"
)

func (s *Service) GetForMonth(userID int, month string) ([]Event, error) {
	date, err := time.Parse("2006-01-02", month)
	if err != nil {
		return nil, err
	}
	monthRepo := date.Month()
	year := date.Year()
	dbEvent, err := s.repo.GetForMonth(userID, monthRepo, year)
	if err != nil {
		return nil, err
	}
	events := make([]Event, 0, len(dbEvent))
	for _, event := range dbEvent {
		var e Event
		e.FillFromDB(&event)
		events = append(events, e)
	}
	return events, nil
}
