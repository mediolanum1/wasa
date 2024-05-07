package database

import (
	"fmt"
)

func (db *appdbimpl) InsertUser(username string) error {

	_, err := db.c.Exec(`INSERT iNTO users (username) VALUES (?);`, username)
	if err != nil {
		return fmt.Errorf("error getting querry from database")
	}

	return nil
}
