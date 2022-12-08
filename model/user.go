package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string `gorm:"varchar(32);not null"`
	Password  string `gorm:"size:255;not null"`
	Email     string `gorm:"varchar(32);not null"`
	Avatar    string `gorm:"size:255;not null"`
	Collects  Array  `gorm:"type:longtext"`
	Following Array  `gorm:"type:longtext"`
	Fans      int    `gorm:"AUTO_INCREMENT"`
}

type UserInfo struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar"`
	UserName string `json:"userName"`
}
