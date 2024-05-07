package database

import (
	"fmt"
)

func (db *appdbimpl) GetPhotos(userId string) ([]Photo, []string, error) {

	var photos []Photo
	var photoIds []string
	rows, err := db.c.Query(`SELECT photoId FROM photos WHERE userId = ?;`, userId)
	if err != nil {
		return photos, photoIds, fmt.Errorf("error in get-photos.go:error counting number of comments")
	}

	for rows.Next() {
		var photoId string
		err := rows.Scan(&photoId)
		if err != nil {
			return photos, photoIds, fmt.Errorf("error in get-photos.go:error scanning thought rows")
		}
		photoIds = append(photoIds, photoId)
	}
	if err = rows.Err(); err != nil {
		return photos, photoIds, fmt.Errorf("error in get-photos.go: error while checking rows")
	}

	for _, val := range photoIds {
		photo, err := db.GetPhoto(val)
		if err != nil {
			return photos, photoIds, fmt.Errorf("error in get-photos.go")
		}
		photos = append(photos, photo)
	}

	return photos, photoIds, nil
}
