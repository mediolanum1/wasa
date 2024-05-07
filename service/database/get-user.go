package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	UserId               string   `json:"userId"`
	Username             string   `json:"username"`
	Number_of_photos     int      `json:"number_of_photos"`
	Number_of_followers  int      `json:"number_of_followers"`
	Number_of_followings int      `json:"number_of_followings"`
	Photos               []Photo  `json:"photos"`
	Banned               []string `json:"banned"`
}

func (db *appdbimpl) GetUser(username string) (User, error) {

	var user User
	var userId string
	err := db.c.QueryRow(`SELECT userId FROM users WHERE username = ?;`, username).Scan(&userId)
	if errors.Is(err, sql.ErrNoRows) {
		return user, fmt.Errorf("error getting querry from database: user not found")
	}
	user.UserId = userId
	user.Username = username
	// now we gonna get the number of photos of user
	var count int
	err = db.c.QueryRow(`SELECT COUNT(*) FROM photos WHERE userId = ?;`, userId).Scan(&count)
	if err != nil {
		return user, fmt.Errorf("error in get-user.go:error counting number of photos")
	}
	user.Number_of_photos = count

	// now we gonna get the number of followers
	err = db.c.QueryRow(`SELECT COUNT(*) FROM followers WHERE followedId = ?;`, userId).Scan(&count)
	if err != nil {
		return user, fmt.Errorf("error in get-user.go:error counting number of followers")
	}
	user.Number_of_followers = count

	// now we gonna get the number of photos of followings
	err = db.c.QueryRow(`SELECT COUNT(*) FROM followers WHERE followerId = ?;`, userId).Scan(&count)
	if err != nil {
		return user, fmt.Errorf("error in get-user.go:error counting number of followings")
	}
	user.Number_of_followings = count

	// init photos array and appending photos that were posted by userId
	user.Photos, _, err = db.GetPhotos(userId)
	if err != nil {
		return user, fmt.Errorf("error in get-user.go:error while getting photos")
	}

	return user, nil
}
