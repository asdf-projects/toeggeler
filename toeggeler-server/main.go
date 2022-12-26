package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pelletier/go-toml"
	"github.com/steinm91/toeggeler/toeggeler-server/api"
	"github.com/steinm91/toeggeler/toeggeler-server/eval"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type EnvVars struct {
	Port    int64
	DBFile  string
	DevMode bool
}

func loadConfig() EnvVars {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		panic(err)
	}

	devMode := config.Get("common.dev").(bool)
	port := config.Get("server.port").(int64)
	dbFile := config.Get("database.file").(string)

	return EnvVars{
		Port:    port,
		DBFile:  dbFile,
		DevMode: devMode,
	}
}

func main() {
	envVars := loadConfig()

	if envVars.DevMode {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	db, err := models.Open(envVars.DBFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.Migrate(db)
	if err != nil {
		panic(err)
	}

	engineEnv := &eval.Env{DB: db}
	evalEngine := eval.NewEvalEngine(engineEnv)

	apiEnv := &api.Env{
		DB:         db,
		Port:       envVars.Port,
		EvalEngine: evalEngine,
		DevMode:    envVars.DevMode,
	}

	api.StartApiServer(apiEnv)
}
