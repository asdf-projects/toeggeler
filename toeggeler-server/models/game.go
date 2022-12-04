package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type GameEventType string

const (
	createGameEvent = "INSERT INTO events(event_game_id, event_data) values(?, ?)"
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
	ID string `json:"id"`
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

	return &Game{ID: gameId}, nil
}
