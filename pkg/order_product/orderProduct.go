package orderProduct //Tabla pivote entre ordenes y productos

import (
	"time"
)

//Model of orderProduct
type Model struct {
	ID        uint
	ProductID uint
	OrderID   uint
	Amount    uint
	CreateAt  time.Time
	UpdateAt  time.Time
}

// Storage infertace that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of orderProduct
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Creare a new orderProduct
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
