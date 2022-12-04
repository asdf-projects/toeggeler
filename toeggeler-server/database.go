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

	CREATE TABLE IF NOT EXISTS games_played(
		game_id TEXT NOT NULL,
		player_id INTEGER NOT NULL,
		team INTEGER NOT NULL,
		position TEXT NOT NULL
	)

	`
	_, err := db.Exec(createTables)
	if err != nil {
		log.Fatal(err)
	}

	//_, _ = db.Exec("DELETE FROM events")
	//_, _ = db.Exec("DELETE FROM games_played")
}
