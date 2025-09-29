package service

import (
	"time"
)

func (s *Service) GetForWeek(userID int, date string) ([]Event, error) {
	dateRepo, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	dbEvent, err := s.repo.GetForWeek(userID, dateRepo)
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
