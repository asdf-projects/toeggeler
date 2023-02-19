package eval

import (
	"database/sql"
	"log"

	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type Env struct {
	DB *sql.DB
}

type EvalEngine struct {
	userService models.UserService
	gameService models.GameService
}

type Stats struct {
	PlayerId int64 `json:"playerId"`
	Rating   int   `json:"rating"`
	Wins     int   `json:"wins"`
	Losses   int   `json:"losses"`
	Goals    int   `json:"goals"`
	Foetelis int   `json:"foetelis"`
	OwnGoals int   `json:"ownGoals"`
}

func NewEvalEngine(env *Env) EvalEngine {
	us := models.UserService{DB: env.DB}
	gs := models.GameService{DB: env.DB}

	return EvalEngine{userService: us, gameService: gs}
}

func (e EvalEngine) Eval() []PlayerStats {
	events, err := e.gameService.GetGameEvents()
	if err != nil {
		log.Println(err)
	}

	games := EvalGames(events)
	playerStats := EvalPlayerStats(games)

	//e.gameService.InsertGames(games)
	//e.gameService.InsertPlayerStats(games)

	return playerStats
}

func (e EvalEngine) GetStats() []Stats {
	events, err := e.gameService.GetGameEvents()
	if err != nil {
		log.Println(err)
	}

	games := EvalGames(events)
	playerStats := EvalPlayerStats(games)

	statsMap := map[int64]*Stats{}

	for _, stats := range playerStats {
		playerId := stats.ID

		_, exists := statsMap[playerId]
		if exists == false {
			statsMap[playerId] = &Stats{
				PlayerId: playerId,
				Wins:     0,
				Losses:   0,
				Goals:    0,
				OwnGoals: 0,
				Foetelis: 0,
			}
		}

		statsMap[playerId].Rating = stats.RatingAfter

		statsMap[playerId].Goals = statsMap[playerId].Goals + stats.Goals + stats.Foetelis
		statsMap[playerId].OwnGoals = statsMap[playerId].OwnGoals + stats.OwnGoals
		statsMap[playerId].Foetelis = statsMap[playerId].Foetelis + stats.Foetelis

		if stats.Won {
			statsMap[playerId].Wins++
		} else {
			statsMap[playerId].Losses++
		}
	}

	statsList := []Stats{}
	for _, entry := range statsMap {
		statsList = append(statsList, *entry)
	}

	return statsList
}

func (e EvalEngine) GetGames() []Game {
	events, err := e.gameService.GetGameEvents()
	if err != nil {
		log.Println(err)
	}

	games := EvalGames(events)

	return games
}
