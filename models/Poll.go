package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Title        string
	Option1      string
	Option2      string
	Option3      string
	Option1Votes int
	Option2Votes int
	Option3Votes int
}
