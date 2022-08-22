package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserId uint `gorm:"not null" json:"userId"`
}