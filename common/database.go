package common

import (
	"fmt"
	"gblog-server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := "root"
	password := "1234"
	host := "localhost"
	port := "3306"
	database := "gblog"
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, database, charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
