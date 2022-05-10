package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	by, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil
	}

	user.Password = string(by)
	return nil
}

func (user *User) CheckPassword(provePassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(provePassword))
	if err != nil {
		return err
	}
	return nil
}
