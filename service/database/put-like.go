package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) PutLike(userId string, photoId string) error {

	// check if there is already a follower entry, if not then we g

	var count int
	err := db.c.QueryRow(`SELECT COUNT(*) FROM likes WHERE userId = ? AND photoId = ?;`, userId, photoId).Scan(&count)
	if err != nil {
		return fmt.Errorf("error in put-like.go:error retieving data from database")
	}
	if count > 0 {
		return fmt.Errorf("error in put-like.go: already liked the photo")
	}

	// get the userId of the author of the photo
	var photoCreatorId string
	err = db.c.QueryRow(`SELECT userId FROM photos WHERE photoId = ?;`, photoId).Scan(&photoCreatorId)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("error getting querry from database in put-like.go")
	}

	if a, err := db.GetIsBanned(userId, photoCreatorId); !a && err == nil {

		_, err1 := db.c.Exec("INSERT INTO likes (userId, photoId) VALUES (?, ?);", userId, photoId)
		if err1 != nil {
			return fmt.Errorf("error in put-like.go: error putting like")
		}
		return nil
	} else if err != nil {
		return fmt.Errorf("error in putt-like.go: error retriving data from database")
	} else if a && err == nil {
		return fmt.Errorf("error in put-like.go: you cant put like under photo of user that banned you")
	}
	return nil
}
