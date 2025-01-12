package database

import "polling-app.com/m/models"

func AutoMigrate() {
	db := GetInstance()
	err := db.AutoMigrate(&models.Poll{})
	if err != nil {
		panic("failed to migrate database")
	}
}
