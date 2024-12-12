package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(100);not null;"`
	Username     string `gorm:"type:varchar(50);not null;"`
	Bio          string `gorm:"type:text;"`
	Image        string `gorm:"type:varchar(255);"`
	PasswordHash string `gorm:"type:varchar(200);not null;"`
}

func (u *User) TableName() string {
	return "users"
}
