package models

import (
	"database/sql"
	"errors"
	"rest-api/configs"
)

type User struct {
	ID       int
	Username string
	Password string // In a real application, make sure to store the hash of the password
}

func CreateUser(username, password string) error {
	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New(`The user exist, please use another username!`)
	}

	_, err = configs.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	return err
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	row := configs.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
