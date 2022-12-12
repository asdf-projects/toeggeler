package models

import (
	"database/sql"
	"log"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"-"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Mail string `json:"mail"`
}

const (
	readUsersStmt  = "SELECT user_id, user_name, user_mail FROM users"
	readUserStmt   = readUsersStmt + " WHERE user_id = ?"
	createUserStmt = "INSERT INTO users(user_name, user_mail, user_password) values(?, ?, ?)"
	updateUserStmt = "UPDATE users SET user_mail = ? WHERE user_id = ?"
	deleteUserStmt = "DELETE FROM users WHERE user_id = ?"
)

func CreateUser(db *sql.DB, userRequest CreateUserRequest) (User, error) {
	var user User
	res, err := db.Exec(createUserStmt, userRequest.Username, userRequest.Mail, userRequest.Password)
	if err != nil {
		return user, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return user, err
	}

	user = User{Id: id, Username: userRequest.Username, Mail: userRequest.Mail}

	return user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(readUsersStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username, &user.Mail); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(db *sql.DB, id string) (User, error) {
	row := db.QueryRow(readUserStmt, id)

	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Mail)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(db *sql.DB, id string, userRequest UpdateUserRequest) (User, error) {
	_, err := db.Exec(updateUserStmt, userRequest.Mail, id)

	log.Println(err)

	if err != nil {
		return User{}, err
	}

	return GetUser(db, id)
}

func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Exec(deleteUserStmt, id)
	return err
}
