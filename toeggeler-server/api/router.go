package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

func GetRouter(env *Env) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(addContentTypeMiddleware)

	router.HandleFunc("/api/users", env.getUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users", env.createUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{name}", env.updateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{name}", env.deleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/{name}", env.getUser).Methods(http.MethodGet)

	return router
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
	name := vars["name"]

	user, err := models.GetUser(env.DB, name)
	if err != nil {
		log.Print(err)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User '"+name+"' not found", 404)
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
	name := vars["name"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var userRequest models.UpdateUserRequest
	json.Unmarshal(reqBody, &userRequest)

	user, err := models.UpdateUser(env.DB, name, userRequest)
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
	name := vars["name"]

	err := models.DeleteUser(env.DB, name)
	if err != nil {
		log.Print(err)

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, " User '"+name+"' not found.", 404)
		} else {
			http.Error(w, "", 500)
		}
		return
	}
}
