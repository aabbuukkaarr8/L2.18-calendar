package repository

import (
	"errors"
	"time"
)

var ErrNoEvents = errors.New("no events found for user")

func (r *Repository) GetForMonth(userID int, month time.Month, year int) ([]Event, error) {

	r.mu.RLock() // только чтение
	defer r.mu.RUnlock()

	userEvents, ok := r.events[userID]
	if !ok || len(userEvents) == 0 {
		return nil, ErrNoEvents // у пользователя нет событий
	}

	var result []Event
	for _, event := range userEvents {
		if event.Date.Month() == month && event.Date.Year() == year {
			result = append(result, event)
		}
	}
	if len(result) == 0 {
		return nil, ErrNoEvents
	}
	return result, nil
}
