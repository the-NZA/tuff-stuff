package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	hashCost = 16 // Hash hashCost for bcrypt
)

// User represents each user in application.
type User struct {
	ID             string `json:"id" db:"id"`
	Login          string `json:"login" db:"login"`
	HashedPassword string `json:"password" db:"password"`
	Password       string `json:"-"`
}

// Validate user's fields.
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Login, validation.Required, validation.Length(minUserNameLength, maxUserNameLength)),
		validation.Field(&u.Password, validation.Required, validation.Length(minPasswordLength, maxPasswordLength)),
	)
}

// HashPassword hashes and store given password.
func (u *User) HashPassword(password string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		return err
	}

	u.HashedPassword = string(hashedBytes)

	return nil
}

// CheckPassword compares hashed and raw passwords.
func (u User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
}
