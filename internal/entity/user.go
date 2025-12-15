package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string `gorm:"not null; size:100"`
	Username string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100)"`
	Password string `gorm:"size:255"`
}
