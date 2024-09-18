package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func Connect() *sql.DB {
	// Connect to the database
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		zap.L().Fatal("Error connecting to the database", zap.Any("error", err))
	}

	err = db.Ping()
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	return db
}
