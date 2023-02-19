package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open(dbSource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbSource)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}

func Migrate(db *sql.DB) error {

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
		event_game_id TEXT NOT NULL,
		event_timestamp DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS games(
		game_id TEXT NOT NULL UNIQUE PRIMARY KEY,
		game_type INTEGER,
		game_start DATETIME NOT NULL,
		game_end DATETIME NOT NULL,
		team1_offense INTEGER NOT NULL,
		team1_defense INTEGER NOT NULL,
		team2_offense INTEGER NOT NULL,
		team2_defense INTEGER NOT NULL,
		team1_goals INTEGER NOT NULL,
		team2_goals INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS mat_view_games(
		game_id TEXT NOT NULL UNIQUE PRIMARY KEY,
		game_start DATETIME NOT NULL,
		game_end DATETIME NOT NULL,
		duration_ms INTEGER NOT NULL,
		team1_offense INTEGER NOT NULL,
		team1_defense INTEGER NOT NULL,
		team1_score INTEGER NOT NULL,
		team2_offense INTEGER NOT NULL,
		team2_defense INTEGER NOT NULL,
		team2_score INTEGER NOT NULL,
		winner_defense INTEGER NOT NULL,
		winner_offense INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS player_stats(
		player_id INTEGER NOT NULL,
		rating INTEGER NOT NULL,
		wins INTEGER NOT NULL,
		losses INTEGER NOT NULL,
		goals INTEGER NOT NULL,
		foetelis INTEGER NOT NULL,
		own_goals INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS game_player_stats(
		game_id TEXT NOT NULL,
		player_id INTEGER NOT NULL,
		team INTEGER NOT NULL,
		position INTEGER NOT NULL,
		teammate_id INTEGER NOT NULL,
		won BOOLEAN NOT NULL,
		goals INTEGER,
		foetelis INTEGER,
		own_goals INTEGER,
		rating INTEGER,
		PRIMARY KEY (game_id, player_id),
		UNIQUE(game_id, player_id)
	);
	`

	_, err := db.Exec(createTables)
	return err
}
