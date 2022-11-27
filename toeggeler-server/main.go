package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/steinm91/toeggeler/toeggeler-server/api"
)

type EnvVars struct {
	Port     string
	DbSource string
	DevMode  bool
}

func loadEnvVars() *EnvVars {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Could not load environment variables. Using defaults.")
		return &EnvVars{":8080", "./toeggeler.sqlite", false}
	}

	apiPort := os.Getenv("API_PORT")
	database := os.Getenv("DB_SOURCE")
	devMode, err := strconv.ParseBool(os.Getenv("DEV_MODE"))
	if err != nil {
		devMode = false
	}

	return &EnvVars{apiPort, database, devMode}
}

func connectToDatabase(dbSource string) *sql.DB {
	db, err := sql.Open("sqlite3", dbSource)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	envVars := loadEnvVars()

	if envVars.DevMode {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	db := connectToDatabase(envVars.DbSource)
	migrateDatabase(db)

	apiEnv := &api.Env{
		DB:   db,
		Port: envVars.Port,
	}

	api.StartApiServer(apiEnv)
}
