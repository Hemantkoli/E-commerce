package models

import (
	"log"
)

func (user *User) FetchAllUsers() ([]User, error) {
	var users []User
	err := DBInstance().Find(&users).Error
	return users, err
}

func (user *User) AddUser() error {
	err := DBInstance().Create(user).Error
	return err
}

func (user *User) RemoveUser(id uint) error {
	removeUser := &User{}
	err := DBInstance().Where("id=?", id).Find(removeUser).Error
	if err != nil {
		log.Println(err)
	}
	return DBInstance().Delete(removeUser).Error
}

func (user *User) UpdateUser() error {
	return DBInstance().Save(user).Error
}