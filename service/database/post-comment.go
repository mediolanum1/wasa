package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) PostComment(userId string, photoId string, comment string, date string) error {

	// get the userId of the author of the photo
	var photoCreatorId string
	err := db.c.QueryRow(`SELECT userId FROM photos WHERE photoId = ?;`, photoId).Scan(&photoCreatorId)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("error getting querry from database in post-comment.go")
	}

	// check if the user is banned by author of the photo
	if a, err := db.GetIsBanned(userId, photoCreatorId); !a && err == nil {
		// Insert the comment data and description into the photos table
		_, err := db.c.Exec("INSERT INTO comments (userId, photoId, comment, date) VALUES (?, ?, ?, ?)", userId, photoId, comment, date)
		if err != nil {
			return fmt.Errorf("error in post-comment.go: error saving into database")
		}
		return nil
	} else if err != nil {
		return fmt.Errorf("error in post-comment.go: error retriving data from database")
	} else if a && err == nil {
		return fmt.Errorf("error in post-comment.go: you cant post under photo of user that banned you")
	}
	return nil
}
