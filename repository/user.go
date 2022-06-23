package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Name     string
	Password string
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	Create(email, name, password string) (*User, error)
}
