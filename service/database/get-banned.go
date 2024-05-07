package database

import (
	"fmt"
)

func (db *appdbimpl) GetBanned(userId string) ([]string, error) {

	rows, err := db.c.Query(`SELECT bannedId FROM banned WHERE userId = ?;`, userId)
	if err != nil {
		return nil, fmt.Errorf("error getting querry from database")
	}

	var banned []string

	for rows.Next() {
		var col1 string
		err := rows.Scan(&col1)
		if err != nil {
			return banned, fmt.Errorf("error in get-banned.go: error with scanning through")
		}
		banned = append(banned, col1)
	}
	if err = rows.Err(); err != nil {
		return banned, fmt.Errorf("error in get-photo.go: error while checking rows")
	}

	return banned, nil
}
