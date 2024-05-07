package database

import (
	"fmt"
)

type Photo struct {
	UserId       int    `json:"userId"`
	PhotoId      int    `json:"photoId"`
	Username     string `json:"username"`
	PhotoDate    string `json:"photoDate"`
	PhotoCaption string `json:"photoCaption"`

	Number_of_comments int       `json:"number_of_comments"`
	Number_of_likes    int       `json:"number_of_likes"`
	Comments           []Comment `json:"comments"`
	Likes              []string  `json:"likes"`
}

func (db *appdbimpl) GetPhoto(photoId string) (Photo, error) {

	// getting information about the photo: who created, photoId, when was uploaded, caption
	var photo Photo
	err := db.c.QueryRow(`SELECT * FROM photos WHERE photoId = ?;`, photoId).Scan(&photo.UserId, &photo.PhotoId, &photo.PhotoDate, &photo.PhotoCaption)
	if err != nil {
		return photo, fmt.Errorf("error getting querry from database")
	}
	// get username
	var userId string
	err = db.c.QueryRow(`SELECT userId FROM photos WHERE photoId = ?;`, photoId).Scan(&userId)
	if err != nil {
		return photo, fmt.Errorf("error in get-photo.go:error error getting userId of author of photo")
	}
	// get username
	var username string
	err = db.c.QueryRow(`SELECT username FROM users WHERE userId = ?;`, userId).Scan(&username)
	if err != nil {
		return photo, fmt.Errorf("error in get-photo.go:error counting number of likes")
	}
	photo.Username = username
	// now we are getting all comments under the photo
	var count int
	err = db.c.QueryRow(`SELECT COUNT(*) FROM comments WHERE photoId = ?;`, photoId).Scan(&count)
	if err != nil {
		return photo, fmt.Errorf("error in get-photo.go:error counting number of comments")
	}
	photo.Number_of_comments = count

	var comments []Comment
	rows, err := db.c.Query(`SELECT * FROM comments WHERE photoId = ?;`, photoId)
	if err != nil {
		return photo, fmt.Errorf("error in get-photo.go: getting comments from database")
	}

	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.CommentId, &comment.UserId, &comment.PhotoId, &comment.Comment, &comment.Date)
		if err != nil {
			return photo, fmt.Errorf("error in get-photo.go: scanning comments")
		}
		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return photo, fmt.Errorf("error in get-photo.go: error while checking rows")
	}
	photo.Comments = comments

	// now we are getting all likes for the photo
	err = db.c.QueryRow(`SELECT COUNT(*) FROM likes WHERE photoId = ?;`, photoId).Scan(&count)
	if err != nil {
		return photo, fmt.Errorf("error in get-photo.go:error counting number of likes")
	}
	photo.Number_of_likes = count
	var likes []string
	rows, err = db.c.Query(`SELECT userId FROM likes WHERE photoId = ?;`, photoId)
	if err != nil {
		return photo, fmt.Errorf("error in get-photo.go:error getting likes")
	}
	for rows.Next() {
		var userId string
		// Scan the userId from the row
		err = rows.Scan(&userId)
		if err != nil {
			return photo, fmt.Errorf("error in get-photo.go: error in scanning likes")
		}

		// Append userId to the slice
		likes = append(likes, userId)
	}
	if err = rows.Err(); err != nil {
		return photo, fmt.Errorf("error in get-photo.go: error while checking rows")
	}
	photo.Likes = likes

	return photo, nil
}
