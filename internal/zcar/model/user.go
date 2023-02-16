package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"index;not null"`
	Password string
	Nickname string `gorm:"unique"`
	Email    string `gorm:"type:varchar(100);unique"`
	Mobile   string `gorm:"size:11;unique"`
	Active   int    `gorm:"default:1"`
}
