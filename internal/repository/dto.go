package repository

import (
	"time"
)

type Event struct {
	UserID int
	Date   time.Time
	Text   string
}

// WeekRange код вычисляющий начало и конец недели
func WeekRange(t time.Time) (time.Time, time.Time) {
	// День недели (0 = воскресенье, 1 = понедельник и т.д.)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // делаем воскресенье = 7
	}

	// Начало недели — понедельник
	start := time.Date(t.Year(), t.Month(), t.Day()-weekday+1, 0, 0, 0, 0, t.Location())
	// Конец недели — воскресенье
	end := start.AddDate(0, 0, 7)

	return start, end
}
