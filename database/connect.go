package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB {
	db, _ := gorm.Open("mysql", "root:@/go-rest")

	db.LogMode(true)

	return db
}
