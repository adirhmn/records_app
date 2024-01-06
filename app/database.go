package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"os"
	"time"
)

func NewDB() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres",connStr)
	helper.PanicifError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
