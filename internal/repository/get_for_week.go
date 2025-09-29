package repository

import "time"

func (r *Repository) GetForWeek(userID int, date time.Time) ([]Event, error) {

	r.mu.RLock() // только чтение
	defer r.mu.RUnlock()

	userEvents, ok := r.events[userID]
	if !ok {
		return nil, ErrNoEvents
	}

	start, end := WeekRange(date)
	var result []Event

	for _, event := range userEvents {
		if !event.Date.Before(start) && event.Date.Before(end) {
			result = append(result, event)
		}
	}

	if len(result) == 0 {
		return nil, ErrNoEvents
	}

	return result, nil
}
