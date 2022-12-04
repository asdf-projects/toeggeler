package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type GameEventType string

const (
	createGameEvent  = "INSERT INTO events(event_game_id, event_data) values(?, ?)"
	createGamePlayed = "INSERT INTO games_played(game_id, player_id, team, position) values(?, ?, ?, ?)"
	getGamesPlayed   = "SELECT game_id FROM games_played where player_id = ?"
	getEventsForGame = "SELECT event_data FROM events where event_game_id = ?"
)

const (
	GAME_START GameEventType = "GAME_START"
	GAME_END   GameEventType = "GAME_END"
	GOAL       GameEventType = "GOAL"
	OWN_GOAL   GameEventType = "OWN_GOAL"
	FOETELI    GameEventType = "FOETELI"
)

func (g *GameEventType) UnmarshalJSON(b []byte) error {
	var s string
	json.Unmarshal(b, &s)
	eventType := GameEventType(s)
	switch eventType {
	case GAME_START, GAME_END, GOAL, OWN_GOAL, FOETELI:
		*g = eventType
		return nil
	}
	return errors.New("Invalid game event type")
}

func (g GameEventType) IsValid() error {
	switch g {
	case GAME_START, GAME_END, GOAL, OWN_GOAL, FOETELI:
		return nil
	}
	return errors.New("Invalid game event type.")
}

type GameEvent struct {
	Id        int64         `json:"id"`
	Timestamp int64         `json:"timestamp"`
	Event     GameEventType `json:"event"`

	// GAME_START
	Team1 *Team `json:"team1,omitempty"`
	Team2 *Team `json:"team2,omitempty"`

	// GOAL | OWN_GOAL | FOETELI
	Player *int64 `json:"player,omitempty"`
}

type Game struct {
	ID     string       `json:"id"`
	Events *[]GameEvent `json:"events,omitempty"`
}

func (e GameEvent) IsValid() error {
	var err = errors.New("Invalid event")
	switch e.Event {
	case GAME_START:
		if e.Player != nil {
			return err
		}

		if e.Team1 != nil && e.Team2 != nil {
			return nil
		}

	case GOAL, OWN_GOAL, FOETELI:
		if e.Team1 != nil || e.Team2 != nil {
			return err
		}

		if e.Player != nil {
			return nil
		}

	case GAME_END:
		if e.Player != nil || e.Team1 != nil || e.Team2 != nil {
			return err
		}

		return nil
	}

	return err
}

type Team struct {
	Offense int64 `json:"offense"`
	Defense int64 `json:"defense"`
}

func GetGamesPlayedForPlayer(db *sql.DB, playerId int64) (*[]Game, error) {
	rows, err := db.Query(getGamesPlayed, playerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	games := []Game{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		games = append(games, Game{ID: id})
	}

	enhancedGames := []Game{}
	for _, game := range games {
		rows, err = db.Query(getEventsForGame, game.ID)
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

			fmt.Println(event)
			events = append(events, event)
		}

		g := Game{ID: game.ID, Events: &events}
		enhancedGames = append(enhancedGames, g)
	}

	return &enhancedGames, nil
}

func SubmitGame(db *sql.DB, gameId string, gameEvents *[]GameEvent) (*Game, error) {
	var err error
	for _, event := range *gameEvents {
		eventJson, err := json.Marshal(event)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		_, err = db.Exec(createGameEvent, gameId, eventJson)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	if err != nil {
		return nil, err
	}

	updateGamesPlayed(db, gameId, gameEvents)

	return &Game{ID: gameId}, nil
}

func updateGamesPlayed(db *sql.DB, gameId string, gameEvents *[]GameEvent) {
	var startEvent GameEvent
	for _, event := range *gameEvents {
		if event.Event == GAME_START {
			startEvent = event
			break
		}
	}

	addGamePlayedForPlayer(db, gameId, startEvent.Team1.Offense, 1, "offense")
	addGamePlayedForPlayer(db, gameId, startEvent.Team1.Defense, 1, "defense")
	addGamePlayedForPlayer(db, gameId, startEvent.Team2.Offense, 2, "offense")
	addGamePlayedForPlayer(db, gameId, startEvent.Team2.Defense, 2, "defense")

	fmt.Println("Updated games played")
}

func addGamePlayedForPlayer(db *sql.DB, gameId string, playerId int64, team int, position string) {
	_, err := db.Exec(createGamePlayed, gameId, playerId, team, position)
	if err != nil {
		fmt.Println("Could not update games_played for player with ID ", playerId)
		fmt.Println(err)
	}
}
