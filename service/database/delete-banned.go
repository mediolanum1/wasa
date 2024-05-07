package database

import (
	"fmt"
)

func (db *appdbimpl) DeleteBanned(userId string, userToUnban string) error {

	_, err := db.c.Exec("DELETE FROM banned WHERE userId=? AND bannedId=?;", userId, userToUnban)
	if err != nil {
		return fmt.Errorf("error in delete-banned.go: coudnt find banned record")
	}
	return nil
}
