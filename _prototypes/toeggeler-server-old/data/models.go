package data

import (
	"encoding/json"
	"errors"
)

type GameEventType string

const (
	GAME_START   GameEventType = "GAME_START"
	GAME_END     GameEventType = "GAME_END"
	GAME_TIMEOUT GameEventType = "GAME_TIMEOUT"
	GOAL         GameEventType = "GOAL"
	OWN_GOAL     GameEventType = "OWN_GOAL"
	COUNTER_GOAL GameEventType = "COUNTER_GOAL"
)

func (g *GameEventType) UnmarshalJSON(b []byte) error {
	var s string
	json.Unmarshal(b, &s)
	eventType := GameEventType(s)
	switch eventType {
	case GAME_START, GAME_END, GAME_TIMEOUT, GOAL, OWN_GOAL, COUNTER_GOAL:
		*g = eventType
		return nil
	}
	return errors.New("invalid game event type")
}

func (g GameEventType) IsValid() error {
	switch g {
	case GAME_START, GAME_END, GAME_TIMEOUT, GOAL, OWN_GOAL, COUNTER_GOAL:
		return nil
	}
	return errors.New("invalid game event type")
}

type Player struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"-"`
	Joined   int64  `json:"joined"`
}

type GameEvent struct {
	Event    GameEventType `json:"event"`
	Time     int64         `json:"time"`
	Setup    *GameSetup    `json:"setup,omitempty"`
	ScoredBy int64         `json:"scoredBy,omitempty"`
}

type Side struct {
	OnOffense int64  `json:"onOffense"`
	OnDefense int64  `json:"onDefense"`
	Color     string `json:"color"`
}

type GameSettings struct {
	TimeLimitInMS int64 `json:"timeLimitInMs"`
	GoalLimit     int64 `json:"goalLimit"`
}

type GameSetup struct {
	Settings GameSettings `json:"settings"`
	Side1    Side         `json:"side1"`
	Side2    Side         `json:"side2"`
}

type Score struct {
	Side1 byte `json:"side1"`
	Side2 byte `json:"side2"`
}

type Game struct {
	ID        int64       `json:"id"`
	Setup     GameSetup   `json:"setup"`
	Score     Score       `json:"score"`
	StartTime int64       `json:"startTime"`
	EndTime   int64       `json:"endTime"`
	Log       []GameEvent `json:"log"`
}
