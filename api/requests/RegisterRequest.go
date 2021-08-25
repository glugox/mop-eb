package requests

import (
	"errors"
)

type RegisterRequest struct {
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	FirstName string    `gorm:"size:100;not null" json:"first_name"`
	LastName  string    `gorm:"size:100;not null" json:"last_name"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
}

func (u *RegisterRequest) Validate() error {

	if u.Email == "" {
		return errors.New("missing Email")
	}
	if u.FirstName == "" {
		return errors.New("missing FirstName")
	}
	if u.LastName == "" {
		return errors.New("missing LastName")
	}
	if u.Password == "" {
		return errors.New("missing Password")
	}

	return nil
}