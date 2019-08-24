package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	First string `gorm:"type:varchar(100)"`
	Last  string `gorm:"type:varchar(100)"`
}
