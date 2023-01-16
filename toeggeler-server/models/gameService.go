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

	InsertGames       = "INSERT INTO games(game_id, game_start, game_end, team1_offense, team1_defense, team2_offense, team2_defense, team1_goals, team2_goals) VALUES "
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

func (gs *GameService) GetGamesPlayed() (*[]Game, error) {
	rows, err := gs.DB.Query(getGamesPlayed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.GameId, &game.GameStart, &game.GameEnd, &game.Offense1, &game.Defense1, &game.Offense2, &game.Defense2, &game.Team1Goals, &game.Team2Goals); err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return &games, nil
}

func (gs *GameService) InsertGames(games []Game) error {
	insertGamesQuery := InsertGames

	for i, game := range games {
		values := fmt.Sprintf("(\"%s\", %d, %d, %d, %d, %d, %d, %d, %d)", game.GameId, game.GameStart, game.GameEnd, game.Offense1, game.Defense1, game.Offense2, game.Defense2, game.Team1Goals, game.Team2Goals)
		if i < len(games)-1 {
			values += ", "
		}
		insertGamesQuery += values
	}

	_, err := gs.DB.Exec(insertGamesQuery)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (gs *GameService) InsertPlayerStats(games []Game) error {
	insertPlayerStatsQuery := InsertPlayerStats

	for i, game := range games {

		values := getValuesString(game, game.Offense1) + ","
		values += getValuesString(game, game.Defense1) + ","
		values += getValuesString(game, game.Offense2) + ","
		values += getValuesString(game, game.Defense2)

		if i < len(games)-1 {
			values += ","
		}

		insertPlayerStatsQuery += values
	}

	log.Println(insertPlayerStatsQuery)

	_, err := gs.DB.Exec(insertPlayerStatsQuery)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func getValuesString(game Game, player int64) string {
	valuesTemplate := "(\"%s\", %d, %d, %d, %d, %t, %d, %d, %d, %d)"

	gameId := game.GameId
	goals := game.PlayerStats[player].Goals
	foetelis := game.PlayerStats[player].Foetelis
	ownGoals := game.PlayerStats[player].OwnGoals
	rating := 0

	var teamMate int64
	var hasWon bool
	var team int
	var position int

	if player == game.Offense1 {
		team = TEAM_1
		position = OFF
		teamMate = game.Defense1
	} else if player == game.Defense1 {
		team = TEAM_1
		position = DEF
		teamMate = game.Offense1
	} else if player == game.Offense2 {
		team = TEAM_2
		position = OFF
		teamMate = game.Defense2
	} else {
		team = TEAM_2
		position = DEF
		teamMate = game.Offense2
	}

	if team == TEAM_1 {
		hasWon = game.Team1Goals > game.Team2Goals
	} else {
		hasWon = game.Team2Goals > game.Team1Goals
	}

	return fmt.Sprintf(valuesTemplate, gameId, player, team, position, teamMate, hasWon, goals, foetelis, ownGoals, rating)
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
