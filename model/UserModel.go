package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	IdUser					uint	`gorm:"primary_key"`
	Email					string
	RoleId					int
	Role					Role
	LastQuestionnaire		time.Time
}

func InitTableUser(DBCon *gorm.DB){
	if !DBCon.HasTable(&User{}) {
		DBCon.AutoMigrate(&User{})
		DBCon.Model(&User{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")
		user := User {Email: "piringku.adm@gmail.com", RoleId:  1}
		DBCon.Create(&user)
	} else {
		DBCon.AutoMigrate(&User{})
	}
}

func CreateUser(user User, db *gorm.DB) User {
	newUser := User{
		Email: user.Email,
		RoleId: 3,
	}

	err := db.Create(&newUser).Error
	if err != nil{
		panic(err)
	}

	return newUser
}

func Login(userTarget User, userBaru User, db *gorm.DB) User {
	err := db.Where("id_user = ?", userTarget.IdUser).First(&userBaru).Error
	if err != nil{
		panic(err)
	}

	return userBaru
}