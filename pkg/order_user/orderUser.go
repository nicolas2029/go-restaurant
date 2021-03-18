package orderUser

import (
	"time"
)

//Model of orderUser
type Model struct {
	ID       uint
	OrderID  uint
	UserID   uint
	CreateAt time.Time
	UpdateAt time.Time
}

// Storage infertace that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of orderUser
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Creare a new orderUser
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
