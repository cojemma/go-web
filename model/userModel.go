package model

import (
	"go-web/database"

	"gorm.io/gorm"
)

var sqldb *gorm.DB

type User struct {
	gorm.Model
	UserName string
}

type UserScope struct {
	ID       int
	UserName string
}

func init() {
	sqldb = database.GetDB()
	sqldb.AutoMigrate(&User{})
}

func GetUser(scope *UserScope) (User, error) {
	user := &User{}
	err := userScopes(sqldb, scope).Take(user).Error

	return *user, err
}

func GetUsers() ([]User, error) {
	var users []User
	err := sqldb.Find(&users).Error

	return users, err
}

func CreateUser(user *User) error {
	return sqldb.Create(user).Error
}

func DeleteUser(scope *UserScope) error {
	user := &User{}
	err := userScopes(sqldb, scope).Delete(user).Error

	return err
}

func UpdateUser(user *User) error {
	return sqldb.Updates(user).Error
}

func userScopes(db *gorm.DB, scope *UserScope) *gorm.DB {
	return db.Scopes(
		findByID(scope.ID),
	)
}

func findByID(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
