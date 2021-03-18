package table

import "time"

//Model of table
type Model struct {
	EstablishmentID uint
	Avalaible       bool
	Capacity        int
	CreateAt        time.Time
	UpdateAt        time.Time
}

// Storage iterface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of table
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new table
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
