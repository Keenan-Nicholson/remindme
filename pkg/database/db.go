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
	CREATE TABLE IF NOT EXISTS reminders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT (datetime('now', 'utc')),
		username TEXT,
		duration INTEGER,
		reminder TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertReminder inserts a new reminder into the database
func InsertReminder(username string, duration time.Duration, reminder string) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	durationSeconds := int(duration.Seconds())

	sqlStmt := `
	INSERT INTO reminders (username, duration, reminder) VALUES (?, ?, ?)
	`

	_, err = db.Exec(sqlStmt, username, durationSeconds, reminder)
	if err != nil {
		log.Fatal(err)
	}
}
