package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

// User model info
// @Description User information
type User struct {
	// ID of the user
	Id int64 `json:"id"`
	// Username
	Username string `json:"username"`
	// Email address
	Mail     string `json:"mail"`
	Password string `json:"-"`
}

type UpdateUserRequest struct {
	Mail string `json:"mail"`
}

var ErrUserExists = errors.New("User already exists")
var ErrMailExists = errors.New("Mail already exists")
var ErrUserNotFound = errors.New("User not found")

const (
	readUsersStmt      = "SELECT user_id, user_name, user_mail FROM users"
	readUserStmt       = readUsersStmt + " WHERE user_id = ?"
	readUserByNameStmt = readUsersStmt + " WHERE user_name = ?"
	readUserPwd        = "SELECT user_password FROM users WHERE user_name = ?"
	createUserStmt     = "INSERT INTO users(user_name, user_mail, user_password) values($1, $2, $3)"
	updateUserStmt     = "UPDATE users SET user_mail = ? WHERE user_id = ?"
	deleteUserStmt     = "DELETE FROM users WHERE user_id = ?"
)

func (us *UserService) Create(name, mail, password string) (*User, error) {
	mail = strings.ToLower(mail)
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	passwordHash := string(hashedBytes)

	user := User{
		Username: name,
		Mail:     mail,
	}

	res, err := us.DB.Exec(createUserStmt, name, mail, passwordHash)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return nil, ErrUserExists
		}

		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.Id = id

	return &user, nil
}

func (us *UserService) GetUsers() (*[]User, error) {
	rows, err := us.DB.Query(readUsersStmt)
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

	return &users, nil
}

func (us *UserService) GetUser(id string) (*User, error) {
	row := us.DB.QueryRow(readUserStmt, id)

	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Mail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GetUserByName(username string) (*User, error) {
	row := us.DB.QueryRow(readUserByNameStmt, username)

	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Mail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, err
}

func (us *UserService) GetUserPassword(username string) (*string, error) {
	row := us.DB.QueryRow(readUserPwd, username)

	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &hashedPassword, nil
}

func (us *UserService) UpdateUser(id string, mail string) (*User, error) {
	_, err := us.DB.Exec(updateUserStmt, mail, id)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return nil, ErrUserExists
		}
	}

	return us.GetUser(id)
}

func (us *UserService) DeleteUser(id string) error {
	_, err := us.DB.Exec(deleteUserStmt, id)

	if errors.Is(err, sql.ErrNoRows) {
		return ErrUserNotFound
	}

	return err
}
