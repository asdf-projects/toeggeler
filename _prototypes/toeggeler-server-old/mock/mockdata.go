package mock

import (
	"database/sql"
	"log"
	"time"

	"github.com/steinm91/toeggeler/data"
	"github.com/steinm91/toeggeler/security"
)

func CreatePlayers() []data.Player {
	hashedPassword, _ := security.HashPassword("555-Nase")

	return []data.Player{
		{ID: 1, Name: "Raffi1", Mail: "raffi1@test.ch", Password: hashedPassword},
		{ID: 2, Name: "Safti", Mail: "safti@test.ch", Password: hashedPassword},
		{ID: 3, Name: "Knafti", Mail: "knafti@test.ch", Password: hashedPassword},
		{ID: 4, Name: "Minelaos", Mail: "minelaos@test.ch", Password: hashedPassword},
	}
}

func CreateGame() data.Game {
	players := CreatePlayers()

	gameStart := data.GameEvent{
		Event: data.GAME_START,
		Time:  100,
		Setup: &data.GameSetup{
			Settings: data.GameSettings{
				GoalLimit: 8,
			},
			Side1: data.Side{
				OnOffense: players[0].ID,
				OnDefense: players[1].ID,
				Color:     "red",
			},
			Side2: data.Side{
				OnOffense: players[2].ID,
				OnDefense: players[3].ID,
				Color:     "blue",
			},
		},
	}

	goalScored := data.GameEvent{
		Event:    data.GOAL,
		Time:     112,
		ScoredBy: gameStart.Setup.Side1.OnOffense,
	}

	ownGoal := data.GameEvent{
		Event:    data.OWN_GOAL,
		Time:     124,
		ScoredBy: gameStart.Setup.Side2.OnDefense,
	}

	gameEnded := data.GameEvent{
		Event: data.GAME_END,
		Time:  180,
	}

	gameEvents := []data.GameEvent{
		gameStart,
		goalScored,
		ownGoal,
		gameEnded,
	}

	game := data.Game{
		ID:    1,
		Setup: *gameStart.Setup,
		Score: data.Score{
			Side1: 3,
			Side2: 4,
		},
		StartTime: 100,
		EndTime:   180,
		Log:       gameEvents,
	}

	return game
}

func FillWithMockData(db *sql.DB) {
	for _, player := range CreatePlayers() {
		InsertPlayer(db, player)
	}
}

func InsertPlayer(db *sql.DB, player data.Player) {
	insertPlayerStmt := "INSERT INTO players(name, mail, password, joined) values(?, ?, ?, ?)"
	res, err := db.Exec(
		insertPlayerStmt,
		player.Name,
		player.Mail,
		player.Password,
		time.Now().Unix(),
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
}
