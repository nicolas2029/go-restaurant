package storage

import (
	"database/sql"

	"github.com/nicolas2029/go-restaurant/pkg/address"
	"github.com/nicolas2029/go-restaurant/storage/template"
)

// Postgresql statements
const (
	psqlGetAllAddress = `SELECT * FROM addresses;`
)

// scanner interface
type scanner interface {
	Scan(dest ...interface{}) error
}

// GetterRows interface that must implement a get template
type GetterRows interface {
	GetRows(rows *sql.Rows) (interface{}, error)
}

// PsqlAddress used for work with postgres - addresses
type PsqlAddress struct {
	db  *sql.DB
	get *template.Get
}

// NewPsqlAddress return a new pointer of PsqlAddress
func NewPsqlAddress(db *sql.DB) *PsqlAddress {
	psql := &PsqlAddress{db, nil}
	psql.get = template.NewGet(psql)
	return psql
}

// scanRowProduct scan a row of the addresses table
func scanRowProduct(s scanner) (*address.Model, error) {
	m := &address.Model{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.PostalCode,
		&m.AddressLine1,
		&m.AddressLine2,
		&m.Colony,
		&m.State,
		&m.City,
		&m.CreateAt,
		&updatedAtNull,
	)

	if err != nil {
		return &address.Model{}, err
	}

	m.UpdateAt = updatedAtNull.Time
	return m, nil
}

// GetRows implement the interface template.Getter
func (s *PsqlAddress) GetRows(rows *sql.Rows) (interface{}, error) {
	ms := make(address.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
}

// GetAll implement the interface address.Storage
func (s *PsqlAddress) GetAll() (address.Models, error) {
	//gets := template.NewGet(s) // Necesito moverla a la estrcutura
	ms, err := s.get.GetAll(s.db, psqlGetAllAddress)
	if v, ok := ms.(address.Models); ok {
		return v, err
	}
	return nil, err
}

// Create implement the interface address.Storage
func (s *PsqlAddress) Create(m *address.Model) error {
	return nil
}
