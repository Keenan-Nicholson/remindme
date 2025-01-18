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
	// Open the database connection globally
	db, err = sql.Open("sqlite3", "./database.db")
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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at DATETIME DEFAULT (datetime('now', 'utc')),
		username TEXT,
		duration INTEGER,
		reminder TEXT
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
	// Convert duration to seconds
	durationSeconds := int(duration.Seconds())

	// Insert the reminder into the database
	sqlStmt := `INSERT INTO reminders (username, duration, reminder) VALUES (?, ?, ?)`
	result, err := db.Exec(sqlStmt, username, durationSeconds, reminder)
	if err != nil {
		return 0, fmt.Errorf("failed to insert reminder: %w", err)
	}

	// Get the ID of the newly inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last inserted ID: %w", err)
	}
	return int(id), nil
}

// DeleteReminder deletes a reminder from the database by its unique row ID
func DeleteReminder(rowID int) error {
	// Delete the reminder with the given ID
	query := "DELETE FROM reminders WHERE id = ?"
	result, err := db.Exec(query, rowID)
	if err != nil {
		log.Println("Error executing DELETE query:", err)
		return fmt.Errorf("failed to delete reminder: %w", err)
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	// If no rows were affected, it means no reminder was found with that ID
	if rowsAffected == 0 {
		log.Printf("No reminder found with ID %d", rowID)
	}
	return nil // Return nil to indicate success
}
