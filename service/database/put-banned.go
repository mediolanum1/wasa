package database

import (
	"fmt"
)

func (db *appdbimpl) PutBanned(userId string, userToBan string) error {

	// check if user that we want to ban is already banned
	var count int
	err := db.c.QueryRow(`SELECT COUNT(*) FROM banned WHERE userId = ? AND bannedId = ?;`, userId, userToBan).Scan(&count)
	if err != nil {
		return fmt.Errorf("error in put-banned.go: error retriving data from database")
	}
	if count > 0 {
		return fmt.Errorf("error in put-banned.go: user already banned")
	}

	// if user is not banned yet, we add him to banned table
	_, err = db.c.Exec("INSERT INTO banned (userId, bannedId) VALUES (?, ?);", userId, userToBan)
	if err != nil {
		return fmt.Errorf("error in put-banned.go: error putting into banned")
	}
	err = db.DeleteFollowing(userId, userToBan)
	if err != nil {
		return fmt.Errorf("error in put-banned.go: error putting into banned")
	}
	err = db.DeleteFollowing(userToBan, userId)
	if err != nil {
		return fmt.Errorf("error in put-banned.go: error putting into banned")
	}

	return nil
}
