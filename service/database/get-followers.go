package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetFollowers(userId string) ([]string, error) {

	rows, err := db.c.Query(`SELECT followerId FROM followers WHERE followedId = ?;`, userId)
	if err != nil {
		return nil, fmt.Errorf("error getting querry from database")
	}

	var followers []string
	var username string
	for rows.Next() {
		var col1 string
		err := rows.Scan(&col1)
		if err != nil {
			return followers, fmt.Errorf("error in get-followings.go: error with scanning through")
		}

		err = db.c.QueryRow(`SELECT username FROM users WHERE userId = ?;`, col1).Scan(&username)
		if errors.Is(err, sql.ErrNoRows) {
			return followers, fmt.Errorf("error getting querry from database: user not found")
		}

		followers = append(followers, username)
	}
	if err = rows.Err(); err != nil {
		return followers, fmt.Errorf("error in get-followers.go: error while checking rows")
	}

	return followers, nil
}
