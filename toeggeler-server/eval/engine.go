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

/*
func (e EvalEngine) EvalGames(events *[]models.GameEvent) []models.Game {
	games := []models.Game{}

	var currentGame *models.Game
	for _, event := range *events {
		switch event.Event {
		case models.GAME_START:
			playerStats := map[int64]*models.Player{}

			playerStats[event.Team1.Offense] = &models.Player{Goals: 0, OwnGoals: 0, Foetelis: 0}
			playerStats[event.Team1.Defense] = &models.Player{Goals: 0, OwnGoals: 0, Foetelis: 0}
			playerStats[event.Team2.Offense] = &models.Player{Goals: 0, OwnGoals: 0, Foetelis: 0}
			playerStats[event.Team2.Defense] = &models.Player{Goals: 0, OwnGoals: 0, Foetelis: 0}

			currentGame = &models.Game{
				GameId:      event.GameId,
				GameStart:   event.Timestamp,
				PlayerStats: playerStats,	rA_1 := 561
	rA_2 := 249
	rB_1 := 402
	rB_2 := 800

	rA := (rA_1 + rA_2) / 2
	rB := (rB_1 + rB_2) / 2

	elo := NewDefaultElo()
	eA := elo.ExpectedWinPercentage(rA, rB)
	eB := elo.ExpectedWinPercentage(rB, rA)
	newRatingA := elo.Rating(rA, rB, 1.0)
	newRatingB := elo.Rating(rB, rA, 0.0)

	deltaA := newRatingA - rA
	deltaB := newRatingB - rB

	log.Printf("Win Chance A: %f", eA)
	log.Printf("%d -> %d, %d -> %d", rA_1, rA_1+deltaA, rA_2, rA_2+deltaA)

	log.Printf("Win Chance B: %f", eB)
	log.Printf("%d -> %d, %d -> %d", rB_1, rB_1+deltaB, rB_2, rB_2+deltaB)
				Offense1:    event.Team1.Offense,
				Defense1:    event.Team1.Defense,
				Offense2:    event.Team2.Offense,
				Defense2:    event.Team2.Defense,
				Team1Goals:  0,
				Team2Goals:  0,
			}

		case models.GAME_END:
			if currentGame == nil || currentGame.GameId != event.GameId {
				continue
			}
			currentGame.GameEnd = event.Timestamp
			games = append(games, *currentGame)

		case models.GOAL, models.OWN_GOAL, models.FOETELI:
			if currentGame == nil || currentGame.GameId != event.GameId {
				continue
			}

			playerId := *event.Player
			eventType := event.Event

			stats, exists := currentGame.PlayerStats[playerId]
			if exists == false {
				log.Printf("Eval Error: Goal was scored by player (id=%d) not registered for the game (gameId=%s)", playerId, event.GameId)
				currentGame = nil
				continue
			}

			if playerId == currentGame.Offense1 || playerId == currentGame.Defense1 {
				if eventType == models.OWN_GOAL {
					currentGame.Team2Goals++
				} else {
					currentGame.Team1Goals++
				}
			} else {
				if eventType == models.OWN_GOAL {
					currentGame.Team1Goals++
				} else {
					currentGame.Team2Goals++
				}
			}

			switch eventType {
			case models.GOAL:
				stats.Goals++
			case models.OWN_GOAL:
				stats.OwnGoals++	rA_1 := 561
	rA_2 := 249
	rB_1 := 402
	rB_2 := 800

	rA := (rA_1 + rA_2) / 2
	rB := (rB_1 + rB_2) / 2

	elo := NewDefaultElo()
	eA := elo.ExpectedWinPercentage(rA, rB)
	eB := elo.ExpectedWinPercentage(rB, rA)
	newRatingA := elo.Rating(rA, rB, 1.0)
	newRatingB := elo.Rating(rB, rA, 0.0)

	deltaA := newRatingA - rA
	deltaB := newRatingB - rB

	log.Printf("Win Chance A: %f", eA)
	log.Printf("%d -> %d, %d -> %d", rA_1, rA_1+deltaA, rA_2, rA_2+deltaA)

	log.Printf("Win Chance B: %f", eB)
	log.Printf("%d -> %d, %d -> %d", rB_1, rB_1+deltaB, rB_2, rB_2+deltaB)
			case models.FOETELI:
				stats.Foetelis++
			}
		}

	}

	return games
}

/*
func getPlayerStats(game models.Game, playerId int64) *models.Player {
	playerStats, exists := game.PlayerStats[playerId]

	if exists {
		return playerStats
	} else {
		return &models.Player{Goals: 0, Foetelis: 0, OwnGoals: 0}
	}
}
*/
