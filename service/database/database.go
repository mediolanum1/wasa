/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: ", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUser(userId string) (User, error)
	GetFollowings(username string) ([]string, error)
	PutFollowing(username string, followToUsername string) error
	ChangeUsername(username string, newUsername string) error
	DeleteFollowing(username string, usernameToUnfollow string) error
	PostToken(token string) (bool, error)
	PostLogin(username string) (string, error)
	InsertUser(username string) error
	PutBanned(userId string, userToBan string) error
	DeleteBanned(userId string, userToUnban string) error
	PostPhoto(userId string, photo multipart.File, description string, date string) error
	GetPhoto(photoId string) (Photo, error)
	PutLike(userId string, photoId string) error
	DeleteLike(userId string, photoId string) error
	PostComment(userId string, photoId string, comment string, date string) error
	GetComments(photoId string) ([]Comment, error)
	DeleteComment(commentId string, userId string) error
	GetIsBanned(userId string, bannedToCheckId string) (bool, error)
	GetStream(userId string) ([]Photo, error)
	GetPhotos(userId string) ([]Photo, []string, error)
	GetFollowers(userId string) ([]string, error)
	GetBanned(userId string) ([]string, error)
	DeletePhoto(userId string, photoId string) error
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

const (
	createUsersTable = ` 
	CREATE TABLE IF NOT EXISTS users (
		userId             INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT  UNIQUE,
		username        TEXT NOT NULL UNIQUE
	);
`
	createCommentsTable = ` 
	CREATE TABLE IF NOT EXISTS comments (
	commentId		    INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
	userId              INTEGER NOT NULL,
	photoId				INTEGER NOT NULL,
	comment 			TEXT NOT NULL,
	date 				TEXT NOT NULL
);
`
	createLikesTable = ` 
	CREATE TABLE IF NOT EXISTS likes (
		userId              INTEGER NOT NULL,
		photoId				INTEGER NOT NULL
		
	);
`

	createFollowingsTable = ` 
	CREATE TABLE IF NOT EXISTS followers (
		id              INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
		followerId       INTEGER NOT NULL,
		followedId		INTEGER NOT NULL

	);
`

	createPhotosTable = ` 
	CREATE TABLE IF NOT EXISTS photos (
		userId             INTEGER NOT NULL,
		photoId				INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
		photoDate  			TEXT NOT NULL ,
		photoCaption		TEXT
	);
`
	createBannedTable = ` 
CREATE TABLE IF NOT EXISTS banned (
	id              INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
	userId             INTEGER NOT NULL,
	bannedId             INTEGER NOT NULL
);
`
)

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	tables := map[string]string{
		"users":      createUsersTable,
		"followings": createFollowingsTable,
		"photos":     createPhotosTable,
		"banned":     createBannedTable,
		"likes":      createLikesTable,
		"comments":   createCommentsTable,
	}

	for tableName, tableCreationQuerry := range tables {
		err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name=?;`, tableName).Scan(&tableName)
		if errors.Is(err, sql.ErrNoRows) {
			_, err = db.Exec(tableCreationQuerry)
			if err != nil {
				return nil, fmt.Errorf("error creating database structure")
			}
		}
	}
	// Check if table exists. If not, the database is empty, and we need to create the structure

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
