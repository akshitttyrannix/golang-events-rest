package users

import (
	"errors"

	"example.com/events/common/database"

	"example.com/events/common/utils"
)

func (user User) save() (*User, error) {
	query := `INSERT INTO users (user_id, username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	_, err = database.DB.Exec(query, user.UserID, user.Username, user.Email, hashedPassword, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}
	return &user, nil
}

func (user *User) checkCredentials() error {
	query := `SELECT user_id, password FROM users WHERE email = ? AND is_deleted = false`
	row := database.DB.QueryRow(query, user.Email)

	var dbPassword string
	err := row.Scan(&user.UserID, &dbPassword)
	if err != nil {
		return errors.New("Unauthorized user")
	}

	if !utils.VerifyPassword(user.Password, dbPassword) {
		return errors.New("Credentials invalid")
	}

	return nil
}
