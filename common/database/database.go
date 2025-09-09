package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("sqlite", "events.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		user_id TEXT PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at INTEGER NOT NULL,
		updated_at INTEGER NOT NULL,
		is_deleted BOOLEAN NOT NULL DEFAULT FALSE
	)`

	if _, err := DB.Exec(createUsersTable); err != nil {
		panic(err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		event_id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		start_date INTEGER NOT NULL,
		end_date INTEGER NOT NULL,
		location TEXT NOT NULL,
		created_at INTEGER NOT NULL,
		updated_at INTEGER NOT NULL,
		is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
		user_id TEXT NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(user_id)
	)`

	if _, err := DB.Exec(createEventsTable); err != nil {
		panic(err)
	}
}
