package database

import (
	"fmt"
)

func (db *appdbimpl) ChangeUsername(userId string, newUsername string) error {

	// Prepare the update statement
	stmt, err := db.c.Prepare(`UPDATE users SET username = ? WHERE userId = ?;`)
	if err != nil {
		return fmt.Errorf("error preparing the update statement in change-username.go")
	}
	defer stmt.Close()

	// Execute the update statement with specific values
	_, err1 := stmt.Exec(newUsername, userId)
	if err1 != nil {

		return fmt.Errorf(" username already exists, choose a new one")
	}

	return nil
}
