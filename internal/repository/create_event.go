package repository

func (r *Repository) Create(e *Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events[e.UserID] = append(r.events[e.UserID], *e)
	return nil
}
