package product

import (
	"time"
)

//Model of product
type Model struct {
	ID          uint
	Name        string
	Price       float64
	Description string
	CreateAt    time.Time
	UpdateAt    time.Time
}

// Storage infertace that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of product
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Creare a new product
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
