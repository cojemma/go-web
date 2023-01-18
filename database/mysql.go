package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open("root:db@tcp(db:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	mysqlDB = db
}

func connectMysql() error {
	db, err := gorm.Open(mysql.Open("root:db@tcp(db:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return err
	}

	mysqlDB = db

	return nil
}

func GetDB() *gorm.DB {
	if mysqlDB == nil {
		connectMysql()
	}

	return mysqlDB
}
