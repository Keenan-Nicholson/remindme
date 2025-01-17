package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the database and creates the table if it doesn't exist
func InitDB() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS duration_reminders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT (datetime('now', 'utc')),
		username TEXT,
		duration INTEGER,
		reminder TEXT
	);

	CREATE TABLE IF NOT EXISTS datetime_reminders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT (datetime('now', 'utc')),
		username TEXT,
		targetTime DATETIME,
		reminder TEXT
	);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertReminder inserts a new reminder into the database
func InsertDurationReminder(username string, duration time.Duration, reminder string) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	durationSeconds := int(duration.Seconds())

	sqlStmt := `
	INSERT INTO duration_reminders (username, duration, reminder) VALUES (?, ?, ?)
	`

	_, err = db.Exec(sqlStmt, username, durationSeconds, reminder)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertDateTimeReminder(username string, targetTime time.Time, reminder string) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	INSERT INTO datetime_reminders (username, targetTime, reminder) VALUES (?, ?, ?)
	`

	_, err = db.Exec(sqlStmt, username, targetTime, reminder)
	if err != nil {
		log.Fatal(err)
	}
}
