package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Telphone string `gorm:"type:varchar(110);not null;unique"`
	Passwd   string `gorm:"type:varchar(255);not null"`
}
