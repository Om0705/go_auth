package models

import (
	"database/sql"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) GetPassword(db *sql.DB) (string, error) {
	var storedPassword string
	query := "SELECT password FROM users WHERE username=$1"
	err := db.QueryRow(query, u.Username).Scan(&storedPassword)
	return storedPassword, err
}
