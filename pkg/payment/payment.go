package payment

import (
	"time"
)

//Model of payment
type Model struct {
	ID            uint
	Paid          bool
	Total         float64
	TypePaymentID uint
	CreateAt      time.Time
	UpdateAt      time.Time
}

// Storage infertace that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of payment
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Creare a new payment
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
