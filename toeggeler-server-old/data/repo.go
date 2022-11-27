package data

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrNotExists = errors.New("row does not exist")
	ErrDuplicate = errors.New("record already exists")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) CreatePlayer(player Player) (*Player, error) {
	insertPlayerStmt := "INSERT INTO players(name, mail, password, joined) values(?, ?, ?, ?)"
	res, err := r.db.Exec(insertPlayerStmt, player.Name, player.Mail, player.Password, player.Joined)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return nil, ErrDuplicate
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	player.ID = id

	return &player, nil
}

func (r *SQLiteRepository) GetAllPlayers() ([]Player, error) {
	rows, err := r.db.Query("SELECT * from players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allPlayers []Player
	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.ID, &player.Name, &player.Mail, &player.Password, &player.Joined); err != nil {
			return nil, err
		}
		allPlayers = append(allPlayers, player)
	}
	return allPlayers, nil
}

func (r *SQLiteRepository) GetPlayerByName(name string) (*Player, error) {
	row := r.db.QueryRow("SELECT * FROM players WHERE name = ?", name)

	var player Player
	if err := row.Scan(&player.ID, &player.Name, &player.Mail, &player.Password, &player.Joined); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &player, nil
}
