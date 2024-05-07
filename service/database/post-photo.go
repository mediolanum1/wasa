package database

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func (db *appdbimpl) PostPhoto(userId string, photo multipart.File, description string, date string) error {

	// Insert the photo data and description into the photos table
	_, err := db.c.Exec("INSERT INTO photos (userId, photoCaption, photoDate) VALUES (?, ?, ?)", userId, description, date)
	if err != nil {
		return fmt.Errorf("error in post-photo.go: error saving into database")
	}

	var photoId string
	// err = db.c.QueryRow(`SELECT photoId FROM photos WHERE id = ?;`, newId).Scan(&photoId)
	err = db.c.QueryRow("SELECT last_insert_rowid()").Scan(&photoId)
	if err != nil {
		return fmt.Errorf("error getting querry from database")
	}

	// saving photo
	out, err := os.Create(filepath.Join("/tmp/photos/", photoId) + ".png")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, photo)
	if err != nil {
		return err
	}

	return nil
}
