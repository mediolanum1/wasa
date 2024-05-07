package database

import (
	"fmt"
	"os"
	"path/filepath"
)

func (db *appdbimpl) DeletePhoto(userId string, photoId string) error {

	_, err1 := db.c.Exec("DELETE FROM photos WHERE userId=? AND photoId=?;", userId, photoId)
	if err1 != nil {
		return fmt.Errorf("error in delete-photo: you can only remove your own photos")
	}

	_, err1 = db.c.Exec("DELETE FROM likes WHERE  photoId=?;", photoId)
	if err1 != nil {
		return fmt.Errorf("error in delete-photo: you can only remove your own photos")
	}

	_, err1 = db.c.Exec("DELETE FROM comments WHERE  photoId=?;", photoId)
	if err1 != nil {
		return fmt.Errorf("error in delete-photo: you can only remove your own photos")
	}

	err := os.Remove(filepath.Join("/tmp/photos/", photoId) + ".png")
	if err != nil {
		return err
	}

	return nil
}
