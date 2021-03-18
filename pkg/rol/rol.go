package rol

import (
	"time"
)

//Model of rol
type Model struct {
	ID       uint
	Name     string
	Level    uint
	CreateAt time.Time
	UpdateAt time.Time
}

// Storage infertace that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of rol
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Creare a new rol
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
