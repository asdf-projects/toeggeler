package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
	"github.com/segmentio/ksuid"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

func GetRouter(env *Env) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(addContentTypeMiddleware)

	router.HandleFunc("/api/users", env.getUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users", env.createUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{id}", env.updateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{id}", env.deleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/{id}", env.getUser).Methods(http.MethodGet)

	router.HandleFunc("/api/games", env.submitGame).Methods(http.MethodPost)
	router.HandleFunc("/api/games", env.getGamesPlayed).Methods(http.MethodGet)

	return router
}

func (env *Env) getGamesPlayed(w http.ResponseWriter, r *http.Request) {
	playerId := r.URL.Query().Get("player")
	if playerId == "" {
		http.Error(w, "Please provide a player id using the query parameter 'player'", 400)
		return
	}

	id, _ := strconv.ParseInt(playerId, 10, 64)
	games, _ := models.GetGamesPlayedForPlayer(env.DB, id)

	json.NewEncoder(w).Encode(games)
}

func (env *Env) submitGame(w http.ResponseWriter, r *http.Request) {
	var err error
	var gameId ksuid.KSUID

	reqBody, _ := ioutil.ReadAll(r.Body)
	var gameEvents []models.GameEvent
	json.Unmarshal(reqBody, &gameEvents)

	gameId, err = ksuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, event := range gameEvents {
		err = event.IsValid()
		if err != nil {
			break
		}
	}

	if err != nil {
		http.Error(w, "Invalid event provided", 400)
		return
	}

	game, err := models.SubmitGame(env.DB, gameId.String(), &gameEvents)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(game)
}

func (env *Env) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers(env.DB)
	if err != nil {
		log.Print(err)
		http.Error(w, "", 500)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (env *Env) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := models.GetUser(env.DB, id)
	if err != nil {
		log.Print(err)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User with id '"+id+"' not found", 404)
		} else {
			http.Error(w, "", 500)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (env *Env) createUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userRequest models.CreateUserRequest
	json.Unmarshal(reqBody, &userRequest)

	user, err := models.CreateUser(env.DB, userRequest)
	if err != nil {
		log.Print(err)

		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			http.Error(w, "User with name '"+userRequest.Username+"' or mail '"+userRequest.Mail+"' already exists.", 400)
		} else {
			http.Error(w, "", 500)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (env *Env) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var userRequest models.UpdateUserRequest
	json.Unmarshal(reqBody, &userRequest)

	user, err := models.UpdateUser(env.DB, id, userRequest)
	if err != nil {
		log.Print(err)

		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			http.Error(w, "Email '"+userRequest.Mail+"' already in use.", 400)
		} else {
			http.Error(w, "", 500)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (env *Env) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := models.DeleteUser(env.DB, id)
	if err != nil {
		log.Print(err)

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, " User '"+id+"' not found.", 404)
		} else {
			http.Error(w, "", 500)
		}
		return
	}
}
