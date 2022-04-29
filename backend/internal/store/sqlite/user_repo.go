package sqlite

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

type UserRepo struct {
	db *sqlx.DB
}

func (u UserRepo) Login(login, password string) (model.User, error) {
	var (
		user model.User
		err  error
	)

	// Find user in database
	err = u.db.Get(&user, "SELECT * FROM users WHERE login = ?", login)
	if err != nil {
		return user, err
	}

	// Compare passwords
	err = user.CheckPassword(password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepo) Register(login, password string) (model.User, error) {
	user := model.User{
		Login:    login,
		Password: strings.TrimSpace(password),
	}

	// Hash password
	if err := user.HashPassword(password); err != nil {
		return user, err
	}

	// Validate user
	if err := user.Validate(); err != nil {
		return user, err
	}

	// Insert new user
	result, err := u.db.Exec(
		"INSERT INTO users (login, password) VALUES (?, ?)",
		user.Login,
		user.HashedPassword,
	)
	if err != nil {
		return user, err
	}

	// Get inserted ID
	id, err := parseID(result)
	if err != nil {
		return user, err
	}

	// Convert ID to string
	user.ID = id

	return user, nil
}
