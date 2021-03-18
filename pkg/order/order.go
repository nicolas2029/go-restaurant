package order

import (
	"time"
)

//Model of order
type Model struct {
	ID              uint
	EstablishmentID uint
	TypeOrderID     uint
	AddressID       uint
	CreateAt        time.Time
	UpdateAt        time.Time
}

// Storage infertace that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of order
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Creare a new order
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
