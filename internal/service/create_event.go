package service

import (
	"L2.18-calendar/internal/repository"
	"errors"
)

func (s *Service) Create(e *Event) error {
	if e.Text == "" {
		return errors.New("empty text")
	}
	toRepo := repository.Event{
		UserID: e.UserID,
		Date:   e.Date,
		Text:   e.Text,
	}

	err := s.repo.Create(&toRepo)
	if err != nil {
		return err
	}
	return nil
}
