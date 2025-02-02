package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database and creates tables if they don't exist.
func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./trainer.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create tables if they don't exist
	queries := []string{
		`CREATE TABLE IF NOT EXISTS trainers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		);`,
		`CREATE TABLE IF NOT EXISTS clients (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			trainer_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			FOREIGN KEY(trainer_id) REFERENCES trainers(id)
		);`,
		`CREATE TABLE IF NOT EXISTS exercises (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS client_exercises (
			client_id INTEGER NOT NULL,
			exercise_id INTEGER NOT NULL,
			FOREIGN KEY(client_id) REFERENCES clients(id),
			FOREIGN KEY(exercise_id) REFERENCES exercises(id),
			PRIMARY KEY(client_id, exercise_id)
		);`,
	}

	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}
