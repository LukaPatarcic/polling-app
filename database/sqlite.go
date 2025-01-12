package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func CheckIfDatabaseFileExists() bool {
	_, err := os.Stat("test.db")
	if os.IsNotExist(err) {
		file, err := os.Create("test.db")
		if err != nil {
			panic("failed to create database file")
		}
		file.Close()
		return false
	}
	return true
}

func GetInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
