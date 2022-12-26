package models

import (
	"encoding/json"
	"errors"
)

type GameEventType string

const (
	TEAM_1 int = 1
	TEAM_2 int = 2

	OFF int = 1
	DEF int = 2
)

const (
	GAME_START GameEventType = "GAME_START"
	GAME_END   GameEventType = "GAME_END"
	GOAL       GameEventType = "GOAL"
	OWN_GOAL   GameEventType = "OWN_GOAL"
	FOETELI    GameEventType = "FOETELI"
)

type GameEvent struct {
	Id        int64         `json:"id"`
	GameId    string        `json:"gameId"`
	Timestamp int64         `json:"timestamp"`
	Event     GameEventType `json:"event"`

	// GAME_START
	Team1 *Team `json:"team1,omitempty"`
	Team2 *Team `json:"team2,omitempty"`

	// GOAL | OWN_GOAL | FOETELI
	Player *int64 `json:"player,omitempty"`
}

type Team struct {
	Offense int64 `json:"offense"`
	Defense int64 `json:"defense"`
}

type Player struct {
	Goals    int
	OwnGoals int
	Foetelis int
}

type Game struct {
	GameId      string
	PlayerStats map[int64]*Player
	Offense1    int64
	Defense1    int64
	Offense2    int64
	Defense2    int64
	GameStart   int64
	GameEnd     int64
	Team1Goals  int
	Team2Goals  int
}

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
