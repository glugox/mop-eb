package requests

import (
	"errors"
)

type LoginRequest struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func (u *LoginRequest) Validate() error {

	if u.Password == "" {
		return errors.New("missing Password")
	}
	if u.Email == "" {
		return errors.New("missing Email")
	}

	return nil
}