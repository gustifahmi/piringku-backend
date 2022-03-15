package model

import "github.com/jinzhu/gorm"

type Role struct {
	IDRole		uint	`gorm:"primary_key"`
	NamaRole	string	`gorm:"unique"`
}

func InitTableRole(DBCon *gorm.DB){
	if !DBCon.HasTable(&Role{}) {
		DBCon.AutoMigrate(&Role{})

		var insert []Role
		insert = append(insert, Role{
			NamaRole: "Admin",
		})
		insert = append(insert, Role{
			NamaRole: "Maintainer",
		})
		insert = append(insert, Role{
			NamaRole: "Visitor",
		})
		for _, ins := range insert {
			DBCon.Create(&ins)
		}
	} else {
		DBCon.AutoMigrate(&Role{})
	}
}