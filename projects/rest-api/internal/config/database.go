package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
func InitDB() *gorm.DB{
	Db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return Db

}