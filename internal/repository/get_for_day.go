package repository

import "time"

func (r *Repository) GetForDay(userID int, data time.Time) ([]Event, error) {

	r.mu.RLock() // только чтение
	defer r.mu.RUnlock()

	userEvents, ok := r.events[userID]
	if !ok || len(userEvents) == 0 {
		return nil, ErrNoEvents // у пользователя нет событий
	}

	var result []Event
	for _, event := range userEvents {
		if event.Date == data {
			result = append(result, event)
		}
	}
	if len(result) == 0 {
		return nil, ErrNoEvents
	}
	return result, nil
}
