package api

import (
	"database/sql"
	"log"
	"net/http"
)

type Env struct {
	DB   *sql.DB
	Port string
}

func StartApiServer(env *Env) {
	router := GetRouter(env)
	log.Fatal(http.ListenAndServe(env.Port, router))
}
