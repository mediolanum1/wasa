package database

import (
	"fmt"
)

func (db *appdbimpl) PutFollowing(userId string, followToUserId string) error {

	// check if we already following user
	var count int
	err := db.c.QueryRow(`SELECT COUNT(*) FROM followers WHERE followerId = ? AND followedId = ?;`, userId, followToUserId).Scan(&count)
	if err != nil {
		return fmt.Errorf("error in put-following.go: error retieving data from database")

	}
	if count > 0 {
		return fmt.Errorf("error in put-following.go: already following user")
	}

	// check if we are banned by user, in case of ban we cannot follow user
	if bannedStatus, err := db.GetIsBanned(userId, followToUserId); !bannedStatus && err == nil {
		_, err1 := db.c.Exec("INSERT INTO followers (followerId, followedId) VALUES (?, ?);", userId, followToUserId)
		if err1 != nil {
			return fmt.Errorf("error in put-following.go: error putting following in database")
		}
		return nil
	} else if err != nil {
		return fmt.Errorf("error in put-following.go: error retriving data from the database")
	} else if bannedStatus && err == nil {
		return fmt.Errorf("error in put-following.go: you can not follow user that banned you")
	}
	return nil
}
