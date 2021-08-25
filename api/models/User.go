package models

import (
	"errors"
	"html"
	"net/mail"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	FirstName string    `gorm:"size:100;not null" json:"first_name"`
	LastName  string    `gorm:"size:100;not null" json:"last_name"`
	Password  string    `gorm:"size:100;not null;" json:"-"` // Dont expose
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {

	switch strings.ToLower(action) {

		case "register":

			if u.Password == "" {
				return errors.New("Password Empty")
			}
			if u.Email == "" {
				return errors.New("Email Empty")
			}
			if u.FirstName == "" {
				return errors.New("FirstName Empty")
			}
			if u.LastName == "" {
				return errors.New("LastName Empty")
			}
			_, err := mail.ParseAddress(u.Email)
			if err != nil {
				return errors.New("Invalid Email")
			}

			return nil

		case "login":
			if u.Password == "" {
				return errors.New("Password Empty")
			}
			if u.Email == "" {
				return errors.New("Email Empty")
			}
			return nil

		default:
			//
	}



	return  nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAll(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error // .Limit(20)
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}