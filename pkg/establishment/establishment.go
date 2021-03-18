package establishment

import "time"

// Model of establishment
type Model struct {
	ID        uint
	AddressID uint
	CreateAt  time.Time
	UpdateAt  time.Time
}

// Storage interface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of establishment
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new establishment
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
