package dao

import "gorm.io/gorm"

type UserDAO struct {
	db *gorm.DB
}

func (u *UserDAO) CreateUser(username, password string) error {
	return nil
}

func (u *UserDAO) UpdateUser(id int, username, password string) error {
    return nil
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db}
}
