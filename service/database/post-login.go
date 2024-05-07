package database

import (
	"fmt"
)

func (db *appdbimpl) PostLogin(username string) (string, error) {
	var userId string
	err := db.c.QueryRow(`SELECT userId FROM users WHERE username =  ?;`, username).Scan(&userId)
	if err != nil {
		return "", fmt.Errorf("error getting querry from database")
	}
	return userId, nil
}
