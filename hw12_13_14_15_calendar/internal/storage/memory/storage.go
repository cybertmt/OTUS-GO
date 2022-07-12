package memorystorage

import (
	"sync"
	"time"

	"github.com/cybertmt/OTUS-GO/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type Storage struct {
	mu     sync.RWMutex
	events map[uuid.UUID]storage.Event
}

func (s *Storage) Create(e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[e.ID]; ok {
		return storage.ErrEventAlreadyExists
	}

	s.events[e.ID] = e
	return nil
}

func (s *Storage) Update(e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events[e.ID] = e
	return nil
}

func (s *Storage) Delete(id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[id]; !ok {
		return storage.ErrEventDoesNotExists
	}

	delete(s.events, id)
	return nil
}

func (s *Storage) FindAll() ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	events := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		events = append(events, event)
	}
	return events, nil
}

func (s *Storage) FindAtMonth(dayStart time.Time) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	interval := dayStart.AddDate(0, 1, 0).Sub(dayStart)
	var events []storage.Event
	for _, event := range s.events {
		diff := event.StartedAt.Sub(dayStart)
		if diff >= 0 && diff <= interval {
			events = append(events, event)
		}
	}
	return events, nil
}

func (s *Storage) Find(id uuid.UUID) (*storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if event, ok := s.events[id]; ok {
		return &event, nil
	}

	return nil, storage.ErrorEventNotFound
}

func New() *Storage {
	return &Storage{
		events: make(map[uuid.UUID]storage.Event),
	}
}
