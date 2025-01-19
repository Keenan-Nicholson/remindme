package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB // Declare a global db connection

// InitDB initializes the database and creates the table if it doesn't exist
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "file:./database.db?cache=shared&mode=rwc&_loc=UTC")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Check if the database connection is alive
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	// Create the table if it doesn't exist
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS reminders (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		created_at DATETIME DEFAULT (datetime('now', 'utc')) NOT NULL,
		username TEXT NOT NULL,
		duration INTEGER NOT NULL,
		reminder TEXT NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	log.Println("Database initialized successfully!")
}

// InsertReminder inserts a new reminder into the database and returns its ID.
func InsertReminder(username string, duration time.Duration, reminder string) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			log.Println("transaction rolled back due to error:", err)
		}
	}()

	_, err = tx.Exec("INSERT INTO reminders (username, duration, reminder) VALUES (?, ?, ?)", username, int(duration.Seconds()), reminder)
	if err != nil {
		return 0, fmt.Errorf("failed to insert reminder: %w", err)
	}

	var id int
	err = tx.QueryRow("SELECT last_insert_rowid()").Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return id, nil

}

func GetReminders() (*sql.Rows, error) {
	// Query the database for all reminders
	rows, err := db.Query("SELECT id, created_at, username, duration, reminder FROM reminders")
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}
	return rows, nil
}

func DeleteReminder(id int) error {
	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Ensure to rollback the transaction in case of an error
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Println("transaction rolled back due to error:", err)
		}
	}()

	// Mutate the database within the transaction
	_, err = tx.Exec("DELETE FROM reminders WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete reminder: %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Success
	log.Println("Reminder deleted successfully!")
	return nil
}
