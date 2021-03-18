package user

import (
	"time"
)

//Model of user
type Model struct {
	ID       uint
	Username string
	Email    string
	Password string
	Name     string
	LastName string
	RoleID   uint
	CreateAt time.Time
	UpdateAt time.Time
}

// Storage infertace that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of user
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Creare a new user
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
