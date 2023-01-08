package eval

import (
	"database/sql"
	"fmt"
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

func NewEvalEngine(env *Env) EvalEngine {
	us := models.UserService{DB: env.DB}
	gs := models.GameService{DB: env.DB}

	return EvalEngine{userService: us, gameService: gs}
}

func (e EvalEngine) Eval() {
	events, err := e.gameService.GetGameEvents()
	if err != nil {
		log.Println(err)
	}

	games := e.EvalGames(events)

	e.gameService.InsertGames(games)
	e.gameService.InsertPlayerStats(games)

	fmt.Println(len(games))
}

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
				PlayerStats: playerStats,
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
				stats.OwnGoals++
			case models.FOETELI:
				stats.Foetelis++
			}
		}

	}
	/*
		for _, game := range games {
			log.Println("============", game.gameId, "============")
			log.Println("Team1:", game.team1Goals)
			log.Println("Team2:", game.team2Goals)
			log.Println("---")
			log.Printf("Player %d (Team 1), Goals: %d, Own Goals: %d, Foetelis: %d", game.offense1, game.playerStats[game.offense1].goals, game.playerStats[game.offense1].ownGoals, game.playerStats[game.offense1].foetelis)
			log.Printf("Player %d (Team 1), Goals: %d, Own Goals: %d, Foetelis: %d", game.defense1, game.playerStats[game.defense1].goals, game.playerStats[game.defense1].ownGoals, game.playerStats[game.defense1].foetelis)
			log.Printf("Player %d (Team 2), Goals: %d, Own Goals: %d, Foetelis: %d", game.offense2, game.playerStats[game.offense2].goals, game.playerStats[game.offense2].ownGoals, game.playerStats[game.offense2].foetelis)
			log.Printf("Player %d (Team 2), Goals: %d, Own Goals: %d, Foetelis: %d", game.defense2, game.playerStats[game.defense2].goals, game.playerStats[game.defense2].ownGoals, game.playerStats[game.defense2].foetelis)
			log.Println("=====================================================")
			log.Println("")
		}
	*/

	return games
}