package storage

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

//NewPostgresDB make a conection to Postgres
func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://admin_restaurant:RestAuraNt_pgsql.561965697@localhost:5433/restaurant?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}
	})
}

//Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}
