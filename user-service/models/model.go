package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique" json:"email"`
	Name  string `json:"name"`
}