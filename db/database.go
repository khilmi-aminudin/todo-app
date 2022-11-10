package db

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {

	connStr := "postgresql://root:secret@localhost:5432/todo?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
