package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/data"
	"github.com/steinm91/toeggeler/mock"
	"github.com/steinm91/toeggeler/routes"
)

func main() {
	// only for dev, show line when errors occur
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// re-create database and fill with mock-data for development and testing
	os.Remove("toeggeler.db")

	db := connectToDatabase("toeggeler.db")
	repo := data.NewSQLiteRepository(db)

	migrateDatabase(db)

	mock.FillWithMockData(db)
	mock.CreateGame()

	controller := routes.New(db, repo)

	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.POST("login", controller.Login)

	v1.GET("players", controller.GetAllPlayers)
	v1.GET("players/:name", controller.GetPlayerByName)
	v1.GET("games/latest", controller.GetGame)
	v1.POST("players", controller.CreatePlayer)

	// just to test enum parsing, since that's kinda whack in go
	v1.POST("event/parse", controller.ParseGameEvent)

	r.Run()
}
