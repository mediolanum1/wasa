package database

import (
	"fmt"
)

type Comment struct {
	CommentId int
	UserId    int
	PhotoId   string
	Comment   string
	Date      string
}

func (db *appdbimpl) GetComments(photoId string) ([]Comment, error) {
	var comments []Comment
	rows, err := db.c.Query(`SELECT * FROM comments WHERE photoId = ?;`, photoId)
	if err != nil {
		return comments, fmt.Errorf("error getting querry from database")
	}

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment)
		if err != nil {
			return comments, fmt.Errorf("error in get-comments.go: error with scanning through")
		}
		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return comments, fmt.Errorf("error in get-comments.go: error while checking rows")
	}
	return comments, nil

}
