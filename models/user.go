package models

import (
	"errors"

	"github.com/go-bongo/bongo"
)

type User struct {
	bongo.DocumentBase `bson:",inline"`
	Mail               string  `json:"mail"`
	Password           string  `json:"password"`
	Session            Session `json:"session"`
}

type Session struct {
	Token string `json:"token"`
}

func (user *User) isValid() (isvalid bool, err error) {
	// if email exist
	//   error

	return false, errors.New("user is not valid")
}

func (user *User) Create() (err error) {

	return nil
}
