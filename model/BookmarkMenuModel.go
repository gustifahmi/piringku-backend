package model

import "github.com/jinzhu/gorm"

type BookmarkMenu struct {
	IdBookmark	uint `gorm:"primary_key"`
	MenuId		int
	UserId		int
}

func InitTableBookmark(DBCon *gorm.DB) {
	DBCon.AutoMigrate(&BookmarkMenu{})
	DBCon.Model(&BookmarkMenu{}).AddForeignKey("menu_id", "menus(id_menu)", "CASCADE", "CASCADE")
	DBCon.Model(&BookmarkMenu{}).AddForeignKey("user_id", "users(id_user)", "CASCADE", "CASCADE")
}

func CreateBookmarkMenu(bookmark BookmarkMenu, db *gorm.DB) BookmarkMenu {
	err := db.Create(&bookmark).Error
	if err != nil{
		panic(err)
	}

	return bookmark
}

func DeleteBookmark(idUser int, idMenu int, db *gorm.DB) BookmarkMenu {
	var bookmark BookmarkMenu

	err := db.Where("user_id = ? AND menu_id = ?", idUser, idMenu).First(&bookmark).Error
	if err != nil {
		panic(err)
	}

	err = db.Delete(&bookmark).Error
	if err != nil {
		panic(err)
	}

	return bookmark
}