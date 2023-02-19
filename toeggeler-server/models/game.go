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

// Game event description
// @Description Game event description
type GameEvent struct {
	// ID of the event
	Id int64 `json:"id"`
	// ID of the game
	GameId string `json:"gameId"`
	// Unix timestamp
	Timestamp int64 `json:"timestamp"`

	// Event type
	Event GameEventType `json:"event"`

	// Required when event is "GAME_START"
	Team1 *Team `json:"team1,omitempty"`

	// Required when event is "GAME_START"
	Team2 *Team `json:"team2,omitempty"`

	// Required when is is "GOAL", "OWN_GOAL" or "FOETELI"
	Player *int64 `json:"player,omitempty"`
}

// @Description A team consists of an offensive and defensive player
type Team struct {
	// ID of the player on offense
	Offense int64 `json:"offense"`
	// ID of the player on defense
	Defense int64 `json:"defense"`
}

type Player struct {
	Goals    int
	OwnGoals int
	Foetelis int
}

type Game struct {
	GameId         string
	GameStart      int64
	GameEnd        int64
	Duration       int
	team1_offense  int64
	team1_defense  int64
	team1_score    int
	team2_offense  int64
	team2_defense  int64
	team2_score    int
	winner_offense int64
	winner_defense int64
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
