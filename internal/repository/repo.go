package repository

import "sync"

type Repository struct {
	mu     sync.RWMutex
	events map[int][]Event
}

func NewRepository() *Repository {
	return &Repository{
		events: make(map[int][]Event),
	}
}
