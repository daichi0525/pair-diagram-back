package db

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func GetDbInstantce() *gorm.DB {
	return db
}
