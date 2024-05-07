package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) GetStream(userId string) ([]Photo, error) {

	// init the array of photos
	var streamNot [][]Photo
	var stream []Photo
	// we need to get all followings of the user
	followings, err := db.GetFollowings(userId)
	if err != nil {
		return stream, fmt.Errorf("error in get-stream.go:error getting followings")
	}

	var followingId string
	for _, value := range followings {
		err = db.c.QueryRow(`SELECT userId FROM users WHERE username = ?;`, value).Scan(&followingId)
		if errors.Is(err, sql.ErrNoRows) {
			return stream, fmt.Errorf("error getting querry from database: user not found")
		}
		photos, _, err := db.GetPhotos(followingId)
		if err != nil {
			return stream, fmt.Errorf("error in get-stream.go:error getting photos from following users")
		}

		streamNot = append(streamNot, photos)
	}
	for _, subArray := range streamNot {
		stream = append(stream, subArray...)
	}

	// now we need to extract photos of users from followings one by one
	return stream, nil
}
