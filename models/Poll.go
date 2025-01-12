package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Title   string
	Option1 string
	Option2 string
	Option3 string
}
