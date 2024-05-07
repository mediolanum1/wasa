package database

import (
	"fmt"
)

func (db *appdbimpl) DeleteComment(commentId string, userId string) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE commentId=? AND userId = ?;", commentId, userId)
	if err != nil {
		return fmt.Errorf("error in delete-comment.go: you can only delete your own comments")
	}
	return nil
}
