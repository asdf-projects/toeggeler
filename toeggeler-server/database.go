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
	
	CREATE TABLE IF NOT EXISTS events(
		event_id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_data TEXT NOT NULL,
		event_game_id TEXT NOT NULL
	);
	`
	_, err := db.Exec(createTables)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM events;")
	if err != nil {
		log.Fatal(err)
	}
}
