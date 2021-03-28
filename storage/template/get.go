package template

import (
	"database/sql"
)

// Getter interface that must implement a db storage
type Getter interface {
	GetRows(rows *sql.Rows) (interface{}, error)
}

// Get used for work with template get
type Get struct {
	get Getter
}

// NewGet return a new pointer of Get
func NewGet(g Getter) *Get {
	return &Get{g}
}

// GetAll return all rows from a table
func (g *Get) GetAll(db *sql.DB, query string) (interface{}, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms, err := g.get.GetRows(rows)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}
