package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"size:255;not null;" json:"first_name"`
	LastName  string `gorm:"size:255;not null;" json:"last_name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Bio       string `gorm:"type:text;not null;" json:"bio"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func (u *User) BeforeSave() error {
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}
