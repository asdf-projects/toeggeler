package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type GameService struct {
	DB *sql.DB
}

const (
	insertGameEvents = "INSERT INTO events(event_game_id, event_timestamp, event_data) VALUES(?, ?, ?)"
	getEventsForGame = "SELECT event_data FROM events WHERE event_game_id = ?"
	getAllEvents     = "SELECT event_data FROM events ORDER BY event_game_id, event_timestamp ASC"

	getGamesPlayed = "SELECT game_id, cast(game_start AS INTEGER), cast(game_end AS INTEGER), team1_offense, team1_defense, team2_offense, team2_defense, team1_goals, team2_goals FROM games ORDER BY game_end ASC"

	InsertGames       = "INSERT INTO mat_view_games(game_id, game_start, game_end, duration, team1_defense, team2_offense, team2_defense, team1_goals, team2_goals) VALUES "
	InsertPlayerStats = "INSERT INTO game_player_stats(game_id, player_id, team, position, teammate_id, won, goals, foetelis, own_goals, rating) VALUES "
)

func (gs *GameService) GetGameEvents() (*[]GameEvent, error) {
	rows, err := gs.DB.Query(getAllEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []GameEvent
	for rows.Next() {
		var eventString string
		var event GameEvent
		if err := rows.Scan(&eventString); err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(eventString), &event)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return &events, nil
}

func (gs *GameService) InsertGameEvents(gameId string, gameEvents *[]GameEvent) (*Game, error) {
	var err error
	for _, event := range *gameEvents {
		eventJson, err := json.Marshal(event)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		_, err = gs.DB.Exec(insertGameEvents, gameId, event.Timestamp, eventJson)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	if err != nil {
		return nil, err
	}

	return &Game{GameId: gameId}, nil
}

func (gs *GameService) ClearGamesTable() error {
	_, err := gs.DB.Exec("DELETE FROM games")
	if err != nil {
		log.Println(err)
	}
	return err
}

func (gs *GameService) ClearGamePlayerStats() error {
	_, err := gs.DB.Exec("DELETE FROM game_player_stats")
	if err != nil {
		log.Println(err)
	}
	return err
}
