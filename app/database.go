package app

import (
	"database/sql"
	"time"

	"github.com/MrBista/crud-go/helper"
)

func NewDB() *sql.DB {
	// db, err = sql.Open("mysql", "root@tcp(localhost:3001)/go_learn")
	// helper.PanicIfError(err)

	// db.SetMaxIdleConns()
	db, err := sql.Open("mysql", "bisma:2308@tcp(127.0.0.1:3001)/go_learn")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)
	return db
}
