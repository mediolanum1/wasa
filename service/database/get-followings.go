package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetFollowings(userId string) ([]string, error) {

	rows, err := db.c.Query(`SELECT followedId FROM followers WHERE followerId = ?;`, userId)
	if err != nil {
		return nil, fmt.Errorf("error getting querry from database")
	}

	var followings []string
	var username string
	for rows.Next() {
		var col1 string
		err := rows.Scan(&col1)
		if err != nil {
			return followings, fmt.Errorf("error in get-followings.go: error with scanning through")
		}
		err = db.c.QueryRow(`SELECT username FROM users WHERE userId = ?;`, col1).Scan(&username)
		if errors.Is(err, sql.ErrNoRows) {
			return followings, fmt.Errorf("error getting querry from database: user not found")
		}
		followings = append(followings, username)
	}
	if err = rows.Err(); err != nil {
		return followings, fmt.Errorf("error in get-followings.go: error while checking rows")
	}
	return followings, nil
}
