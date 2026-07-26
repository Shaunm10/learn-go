package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3" // underscore will keep this import
)

// static variables
var DatabaseInstance *sql.DB

func InitDB() {
	databasePath := "./api.db"
	var err error
	DatabaseInstance, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		panic("Could not connect to database.")
	}

	DatabaseInstance.SetMaxOpenConns(10)
	DatabaseInstance.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTableScript := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT null UNIQUE,
		password TEXT NULL,
		createDate DATETIME NOT NULL
		)
	`

	_, err := DatabaseInstance.Exec(createUsersTableScript)

	if err != nil {
		panicError := errors.Join(errors.New("Could not create users table"), err)

		panic(panicError)
	}

	createEventsTableScript := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DatabaseInstance.Exec(createEventsTableScript)

	if err != nil {
		panicError := errors.Join(errors.New("Could not create events table"), err)

		panic(panicError)
	}
}
