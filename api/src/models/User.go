package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_At time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("The name is required and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("The nick is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("The email is required and cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("The email provided is not valid")
	}

	if stage == "create" && user.Password == "" {
		return errors.New("The password is required and cannot be empty")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "create" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return nil
}
