package address

import "time"

// Model of address
type Model struct {
	ID           uint
	PostalCode   string
	AddressLine1 string
	AddressLine2 string
	Colony       string
	State        string
	City         string
	CreateAt     time.Time
	UpdateAt     time.Time
}

type Models []*Model

// Storage interface that must implement a db storage
type Storage interface {
	Create(*Model) error
	GetAll() (Models, error)
}

// Service of address
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new address
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}
