package database

import (
	"fmt"
)

func (db *appdbimpl) DeleteFollowing(userId string, userToUnfollow string) error {

	_, err1 := db.c.Exec("DELETE FROM followers WHERE followerId=? AND followedId=?;", userId, userToUnfollow)
	if err1 != nil {
		return fmt.Errorf("error delete following")
	}
	return nil
}
