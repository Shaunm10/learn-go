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

	createUserTable()
	createEventsTable()
	createRegistrationTable()
}

func createUserTable() {
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
}

func createEventsTable() {
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
	_, err := DatabaseInstance.Exec(createEventsTableScript)

	if err != nil {
		panicError := errors.Join(errors.New("Could not create events table"), err)

		panic(panicError)
	}
}

func createRegistrationTable() {

	createRegistrationTableScript := `
	CREATE TABLE IF NOT registrations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		event_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(event_id) REFERENCES events(id)
	)
	`
	_, err := DatabaseInstance.Exec(createRegistrationTableScript)

	if err != nil {
		panicError := errors.Join(errors.New("Could not create registration table"), err)

		panic(panicError)
	}
}
