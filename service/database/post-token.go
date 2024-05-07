package database

import (
	"fmt"
)

func (db *appdbimpl) PostToken(token string) (bool, error) {
	var valid bool
	err := db.c.QueryRow(`SELECT TRUE FROM users WHERE userId = ?;`, token).Scan(&valid)
	if err != nil {
		return false, fmt.Errorf("error getting querry from database")
	}

	return valid, nil
}
