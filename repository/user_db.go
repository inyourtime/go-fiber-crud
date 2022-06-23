package repository

import "gorm.io/gorm"

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetByEmail(email string) (*User, error) {

	user := User{}

	tx := r.db.Where("email=?", email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) Create(email, name, password string) (*User, error) {

	user := User{Email: email, Name: name, Password: password}

	tx := r.db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
