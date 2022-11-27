package main

import (
	"database/sql"
	"log"
)

func migrateDatabase(db *sql.DB) {
	createTables := `
	CREATE TABLE IF NOT EXISTS users(
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_name TEXT NOT NULL UNIQUE,
		user_mail TEXT NOT NULL UNIQUE,
		user_password TEXT NOT NULL,
		user_created DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createTables)
	if err != nil {
		log.Fatal(err)
	}
}
