package typePayment

import (
	"time"
)

//Model of typePayment
type Model struct {
	ID       uint
	Name     string
	CreateAt time.Time
	UpdateAt time.Time
}

// Storage infertace that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of typePayment
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Creare a new typePayment
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
