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

func checkCredentials(email string, password string) error {
	query := `SELECT password FROM users WHERE email = ? AND is_deleted = false`
	row := database.DB.QueryRow(query, email)

	var dbPassword string
	err := row.Scan(&dbPassword)
	if err != nil {
		return errors.New("Unauthorized user")
	}

	if !utils.VerifyPassword(password, dbPassword) {
		return errors.New("Credentials invalid")
	}

	return nil
}
