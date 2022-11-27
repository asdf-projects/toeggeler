package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func connectToDatabase(name string) (db *sql.DB) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	return db
}

func migrateDatabase(db *sql.DB) {
	createTables := `
	CREATE TABLE IF NOT EXISTS players(
		player_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		mail TEXT,
		password TEXT NOT NULL,
		joined INTEGER
	);

	CREATE TABLE IF NOT EXISTS games(
		games_id INTEGER PRIMARY KEY AUTOINCREMENT,
		time_limit INTEGER,
		goal_limit INTEGER,
		side1_off_player INTEGER NOT NULL,
		side1_def_player INTEGER NOT NULL,
		side1_color TEXT,
		side1_goals INTEGER NOT NULL,
		side2_off_player INTEGER NOT NULL,
		side2_def_player INTEGER NO NULL,
		side2_color TEXT,
		side2_goals INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS gamelog(
		gamelog_id INTEGER PRIMARY KEY AUTOINCREMENT,
		game_id INTEGER NOT NULL,
		event INTEGER NOT NULL,
		time INTEGER NOT NULL,
		time_limit_ms INTEGER,
		goal_limit INTEGER,
		side1_off_player INTEGER,
		side1_def_player INTEGER,
		side1_color TEXT,
		side2_off_player,
		side2_def_player,
		side2_color TEXT,
		scored_by_player INTEGER
	);

	CREATE TABLE IF NOT EXISTS playerstats (
		playerstats_id INTEGER PRIMARY KEY AUTOINCREMENT,
		evaluation_date INTEGER NOT NULL,
		player_id INTEGER NOT NULL,
		from_date INTEGER,
		to_date INTEGER,
		elo INTEGER,
		ovr_wins INTEGER,
		ovr_loses INTEGER,
		ovr_goals_for INTEGER,
		ovr_goals_against INTEGER
	);
	`
	_, err := db.Exec(createTables)
	if err != nil {
		log.Fatal(err)
	}
}
