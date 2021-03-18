package typeOrder

import "time"

// Model of typeOrder
type Model struct {
	ID       uint
	name     string
	CreateAt time.Time
	UpdateAt time.Time
}

// Storage iterface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of type_order
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new type_order
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
