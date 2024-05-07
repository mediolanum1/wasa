package database

import (
	"fmt"
)

func (db *appdbimpl) DeleteLike(userId string, photoId string) error {

	_, err1 := db.c.Exec("DELETE FROM likes WHERE userId=? AND photoId=?;", userId, photoId)
	if err1 != nil {
		return fmt.Errorf("error in delete-like: you can only remove your own likes")
	}
	return nil
}
