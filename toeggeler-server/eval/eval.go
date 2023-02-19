package eval

import (
	"time"

	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type PlayerStats struct {
	ID           int64
	GameDate     int64
	GameId       string
	Team         int
	Position     string
	Won          bool
	Goals        int
	Foetelis     int
	OwnGoals     int
	RatingBefore int
	RatingAfter  int
}

type Team struct {
	Off      int64
	Def      int64
	OwnGoals int
}

type Game struct {
	GameId    string
	GameStart string
	GameEnd   string
	Duration  int64
	Stats     map[int64]*PlayerStats
	Team1     Team
	Team2     Team
	Winner    int
	Score     struct {
		Team1 int
		Team2 int
	}
}

func EvalGames(events *[]models.GameEvent) []Game {
	gamesPlayed := []Game{}

	var currentGame Game

	for _, event := range *events {
		switch event.Event {
		case models.GAME_START:
			currentGame = NewGame(event)

		case models.GOAL, models.FOETELI, models.OWN_GOAL:
			currentGame.OnGoal(event)

		case models.GAME_END:
			currentGame.OnGameEnd(event)
			gamesPlayed = append(gamesPlayed, currentGame)
		}
	}

	return gamesPlayed
}

func EvalPlayerStats(games []Game) []PlayerStats {
	playerStats := []PlayerStats{}
	latestRatings := map[int64]int{}

	for _, game := range games {
		statsMap := game.Stats

		for player := range statsMap {
			latestRating, exists := latestRatings[player]
			if exists == false {
				latestRatings[player] = 400
				latestRating = 400
			}

			statsMap[player].RatingBefore = latestRating
		}

		game.updatePlayerRatings()

		latestRatings[game.Team1.Off] = game.Stats[game.Team1.Off].RatingAfter
		latestRatings[game.Team1.Def] = game.Stats[game.Team1.Def].RatingAfter
		latestRatings[game.Team2.Off] = game.Stats[game.Team2.Off].RatingAfter
		latestRatings[game.Team2.Def] = game.Stats[game.Team2.Def].RatingAfter

		playerStats = append(playerStats, game.getPlayerStats()...)
	}

	return playerStats
}

func getLatestRatingForPlayer(player int64) int {
	return 400
}

func NewGame(event models.GameEvent) Game {
	t1 := Team{
		Off:      event.Team1.Offense,
		Def:      event.Team1.Defense,
		OwnGoals: 0,
	}

	t2 := Team{
		Off:      event.Team2.Offense,
		Def:      event.Team2.Defense,
		OwnGoals: 0,
	}

	gameStart := time.UnixMilli(event.Timestamp).Format(time.RFC3339)

	stats := map[int64]*PlayerStats{}
	stats[t1.Off] = newStats(t1.Off, 1, "Offense", event.Timestamp)
	stats[t1.Def] = newStats(t1.Def, 1, "Defense", event.Timestamp)
	stats[t2.Off] = newStats(t2.Off, 2, "Offense", event.Timestamp)
	stats[t2.Def] = newStats(t2.Def, 2, "Defense", event.Timestamp)

	return Game{
		GameId:    event.GameId,
		GameStart: gameStart,
		Stats:     stats,
		Team1:     t1,
		Team2:     t2,
	}
}

func newStats(playerId int64, team int, position string, date int64) *PlayerStats {
	return &PlayerStats{
		ID:       playerId,
		GameDate: date,
		Team:     team,
		Position: position,
		Won:      false,
		Goals:    0,
		Foetelis: 0,
		OwnGoals: 0,
	}
}

func (g *Game) OnGameEnd(event models.GameEvent) {
	gameEnd := time.UnixMilli(event.Timestamp).Format(time.RFC3339)
	g.GameEnd = gameEnd

	if g.Score.Team1 > g.Score.Team2 {
		g.Winner = 1
		g.Stats[g.Team1.Off].Won = true
		g.Stats[g.Team1.Def].Won = true
	} else {
		g.Winner = 2
		g.Stats[g.Team2.Off].Won = true
		g.Stats[g.Team2.Def].Won = true
	}
}

func (g *Game) OnGoal(event models.GameEvent) {
	player := *event.Player
	isOwnGoal := false

	switch event.Event {
	case models.GOAL:
		g.Stats[player].Goals++
	case models.OWN_GOAL:
		g.Stats[player].OwnGoals++
		isOwnGoal = true
	case models.FOETELI:
		g.Stats[player].Goals++
		g.Stats[player].Foetelis++
	}

	if player == g.Team1.Off || player == g.Team1.Def {
		if isOwnGoal {
			g.Score.Team2++
			g.Team1.OwnGoals++

		} else {
			g.Score.Team1++
		}
	}

	if player == g.Team2.Off || player == g.Team2.Def {
		if isOwnGoal {
			g.Score.Team1++
			g.Team2.OwnGoals++
		} else {
			g.Score.Team2++
		}
	}
}

func (g *Game) getPlayerStats() []PlayerStats {
	players := g.getPlayers()
	playerStats := []PlayerStats{}

	for _, player := range players {
		playerStat := *g.Stats[player]
		playerStat.GameId = g.GameId
		playerStats = append(playerStats, playerStat)
	}

	return playerStats
}

func (g *Game) getPlayers() []int64 {
	return []int64{
		g.Team1.Off, g.Team1.Def, g.Team2.Off, g.Team2.Def,
	}
}

func (g *Game) updatePlayerRatings() {
	elo := NewDefaultElo()

	stats := g.Stats

	ratingT1 := (stats[g.Team1.Off].RatingBefore + stats[g.Team1.Def].RatingBefore) / 2
	ratingT2 := (stats[g.Team2.Off].RatingBefore + stats[g.Team2.Def].RatingBefore) / 2

	newRatingT1 := elo.Rating(ratingT1, ratingT2, stats[g.Team1.Off].Won)
	newRatingT2 := elo.Rating(ratingT2, ratingT1, stats[g.Team2.Off].Won)

	deltaT1 := newRatingT1 - ratingT1
	deltaT2 := newRatingT2 - ratingT2

	stats[g.Team1.Off].RatingAfter = stats[g.Team1.Off].RatingBefore + deltaT1
	stats[g.Team1.Def].RatingAfter = stats[g.Team1.Def].RatingBefore + deltaT1
	stats[g.Team2.Off].RatingAfter = stats[g.Team2.Off].RatingBefore + deltaT2
	stats[g.Team2.Def].RatingAfter = stats[g.Team2.Def].RatingBefore + deltaT2
}
