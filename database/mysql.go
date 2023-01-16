package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func init() {
	mysqlDB, _ = gorm.Open(mysql.Open("root:db@tcp(db:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
}

func GetDB() *gorm.DB {
	return mysqlDB
}
