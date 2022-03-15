package model

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Menu struct {
	IdMenu			uint			`gorm:"primary_key"`
	NamaMenu		string
	Ingredients		postgres.Jsonb
	CaraMasak		postgres.Jsonb
	FotoMenu		postgres.Jsonb
	Sarapan			bool
	MakanSiang		bool
	MakanMalam		bool
	Selingan		bool
	KelompokMakanan	postgres.Jsonb
	Pertanyaan		postgres.Jsonb
	TotalLike		int
	Kalori			float64
	Karbohidrat		float64
	Lemak			float64
	Protein			float64
	VitaminA		float64
	VitaminD		float64
	Folat			float64
	ZatBesi			float64
	Iodin			float64
	Zinc			float64
}

func InitTableMenu(DBCon *gorm.DB){
	DBCon.AutoMigrate(&Menu{})
}

func CreateMenu(menu Menu, db *gorm.DB) Menu {
	err := db.Create(&menu).Error
	if err != nil{
		panic(err)
	}

	return menu
}

func GetMenuList(page int, db *gorm.DB) []Menu {
	var daftarMenu []Menu

	err := db.Order("id_menu desc").Offset((page-1)*10).Limit(10).Find(&daftarMenu).Error
	if err != nil {
		panic(err)
	}

	return daftarMenu
}

func GetSarapanMenuList(page int, db *gorm.DB) []Menu {
	var daftarMenu []Menu

	err := db.Where("sarapan = ?", true).Order("id_menu desc").Offset((page-1)*10).Limit(10).Find(&daftarMenu).Error
	if err != nil {
		panic(err)
	}

	return daftarMenu
}

func GetMakanSiangMenuList(page int, db *gorm.DB) []Menu {
	var daftarMenu []Menu

	err := db.Where("makan_siang = ?", true).Order("id_menu desc").Offset((page-1)*10).Limit(10).Find(&daftarMenu).Error
	if err != nil {
		panic(err)
	}

	return daftarMenu
}

func GetMakanMalamMenuList(page int, db *gorm.DB) []Menu {
	var daftarMenu []Menu

	err := db.Where("makan_malam = ?", true).Order("id_menu desc").Offset((page-1)*10).Limit(10).Find(&daftarMenu).Error
	if err != nil {
		panic(err)
	}

	return daftarMenu
}

func GetSelinganMenuList(page int, db *gorm.DB) []Menu {
	var daftarMenu []Menu

	err := db.Where("selingan = ?", true).Order("id_menu desc").Offset((page-1)*10).Limit(10).Find(&daftarMenu).Error
	if err != nil {
		panic(err)
	}

	return daftarMenu
}

func GetBookmarkedMenuList(idUser uint, db *gorm.DB) []Menu {
	var daftarMenu []Menu

	db.Table("menus").Joins("left join bookmark_menus on bookmark_menus.menu_id = menus.id_menu").
		Where("bookmark_menus.user_id = ?", idUser).Order("id_bookmark desc").Scan(&daftarMenu)

	return daftarMenu
}

func GetDetailMenu(idMenu uint, db *gorm.DB) Menu {
	var menu Menu

	err := db.Where("id_menu = ?", idMenu).First(&menu).Error
	if err != nil {
		panic(err)
	}

	return menu
}

func UpdateMenu(menuLama Menu, menuBaru Menu, db *gorm.DB) Menu {
	err := db.Model(&menuLama).Updates(&menuBaru).Error
	if err != nil{
		panic(err)
	}

	if !menuBaru.Sarapan {
		db.Model(&menuLama).Update("sarapan", false)
	}

	if !menuBaru.MakanSiang {
		db.Model(&menuLama).Update("makan_siang", false)
	}

	if !menuBaru.MakanMalam {
		db.Model(&menuLama).Update("makan_malam", false)
	}

	if !menuBaru.Selingan {
		db.Model(&menuLama).Update("selingan", false)
	}

	if menuBaru.VitaminA == 0 {
		db.Model(&menuLama).Update("vitamin_a", 0)
	}

	if menuBaru.VitaminD == 0 {
		db.Model(&menuLama).Update("vitamin_d", 0)
	}

	if menuBaru.Folat == 0 {
		db.Model(&menuLama).Update("folat", 0)
	}

	if menuBaru.ZatBesi == 0 {
		db.Model(&menuLama).Update("zat_besi", 0)
	}

	if menuBaru.Iodin == 0 {
		db.Model(&menuLama).Update("iodin", 0)
	}

	if menuBaru.Zinc == 0 {
		db.Model(&menuLama).Update("zinc", 0)
	}

	return menuBaru
}

func DeleteMenu(idMenu uint, db *gorm.DB) Menu {
	menu := GetDetailMenu(idMenu, db)

	err := db.Delete(&menu).Error
	if err != nil {
		panic(err)
	}

	return menu
}