package database

import (
	"fmt"
)

func (db *appdbimpl) GetIsBanned(userId string, bannedToCheckId string) (bool, error) {

	// get number of instances from DB where userId banned bannedToCheckId
	// if count > 0, then returns true meaning that userId banned bannedToCheckId
	var count int
	err := db.c.QueryRow(`SELECT COUNT(*) FROM banned WHERE userId = ? AND bannedId = ?;`, userId, bannedToCheckId).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error in get-isBanned.go:error retieving data from database")
	}
	if count > 0 {
		return true, nil
	}
	err = db.c.QueryRow(`SELECT COUNT(*) FROM banned WHERE userId = ? AND bannedId = ?;`, bannedToCheckId, userId).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error in get-isBanned.go:error retieving data from database")
	}
	if count > 0 {
		return true, nil
	}
	return false, nil

}
