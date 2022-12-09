package common

import (
	"fmt"
	"gblog-server/model"
	"net/url"

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
	loc := "Asia/Shanghai"

	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=%s", user, password, host, port, database, charset, url.QueryEscape(loc))
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
