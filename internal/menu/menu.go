package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/piringku-web/piringku-backend/database"
	"github.com/piringku-web/piringku-backend/model"
	"strconv"
)

func CreateNewMenu(c *gin.Context){
	var menu model.Menu

	if err := c.ShouldBind(&menu); err != nil {
		panic(err)
	}

	menu = model.CreateMenu(menu, database.DBCon)

	c.JSON(201, gin.H{"Menu telah berhasil dibuat": &menu})
}

func GetCount(c *gin.Context) {
	kategori := c.Param("kategori")
	var count int

	if kategori == "all" {
		err := database.DBCon.Table("menus").Count(&count).Error
		if err != nil {
			panic(err)
		}
	} else if kategori == "sarapan" {
		err := database.DBCon.Table("menus").Where("sarapan = ?", true).Count(&count).Error
		if err != nil {
			panic(err)
		}
	} else if kategori == "makansiang" {
		err := database.DBCon.Table("menus").Where("makan_siang = ?", true).Count(&count).Error
		if err != nil {
			panic(err)
		}
	} else if kategori == "makanmalam" {
		err := database.DBCon.Table("menus").Where("makan_malam = ?", true).Count(&count).Error
		if err != nil {
			panic(err)
		}
	} else if kategori == "selingan" {
		err := database.DBCon.Table("menus").Where("selingan = ?", true).Count(&count).Error
		if err != nil {
			panic(err)
		}
	}

	c.JSON(200, count)
}

func GetDaftarSemuaMenu(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Param("page"), 10, 64)
	daftarMenu := model.GetMenuList(int(page), database.DBCon)

	c.JSON(200, daftarMenu)
}

func GetDaftarMenuSarapan(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Param("page"), 10, 64)
	daftarMenu := model.GetSarapanMenuList(int(page), database.DBCon)

	c.JSON(200, daftarMenu)
}

func GetDaftarMenuMakanSiang(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Param("page"), 10, 64)
	daftarMenu := model.GetMakanSiangMenuList(int(page), database.DBCon)

	c.JSON(200, daftarMenu)
}

func GetDaftarMenuMakanMalam(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Param("page"), 10, 64)
	daftarMenu := model.GetMakanMalamMenuList(int(page), database.DBCon)

	c.JSON(200, daftarMenu)
}

func GetDaftarMenuSelingan(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Param("page"), 10, 64)
	daftarMenu := model.GetSelinganMenuList(int(page), database.DBCon)

	c.JSON(200, daftarMenu)
}

func GetMenuDisimpan(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		panic(err)
	}

	daftarMenu := model.GetBookmarkedMenuList(uint(idUser),database.DBCon)

	c.JSON(200, daftarMenu)
}

func GetMenuById(c *gin.Context){
	var bookmark model.BookmarkMenu
	var isBookmarked bool
	isLogin := true

	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		isBookmarked = false
		isLogin = false
	}

	idMenu, err := strconv.ParseUint(c.Param("idMenu"), 10, 64)

	if err != nil{
		panic(err)
	}

	if isLogin {
		err = database.DBCon.Where("user_id = ? AND menu_id = ?", idUser, idMenu).First(&bookmark).Error
		if err != nil {
			isBookmarked = false
		} else {
			isBookmarked = true
		}
	}

	menu := model.GetDetailMenu(uint(idMenu), database.DBCon)

	c.JSON(200, gin.H{"Menu": menu, "IsBookmarked": isBookmarked})
}

func GetDetailMenu(c *gin.Context) {
	idMenu, _ := strconv.ParseUint(c.Param("idMenu"), 10, 64)

	menu := model.GetDetailMenu(uint(idMenu), database.DBCon)

	c.JSON(200, menu)
}

func UpdateMenuById(c *gin.Context){
	idMenu, err := strconv.ParseUint(c.Param("idMenu"), 10, 64)

	if err != nil {
		panic(err)
	}

	var menuBaru model.Menu

	if err := c.ShouldBind(&menuBaru); err != nil{
		panic(err)
	}

	menuLama := model.GetDetailMenu(uint(idMenu), database.DBCon)

	menuBaru = model.UpdateMenu(menuLama, menuBaru, database.DBCon)

	c.JSON(200, gin.H{"Menu telah berhasil diupdate" : &menuBaru})
}

func DeleteMenuById(c *gin.Context){
	idMenu, err := strconv.ParseUint(c.Param("idMenu"), 10, 64)

	if err != nil{
		panic(err)
	}

	menu := model.DeleteMenu(uint(idMenu), database.DBCon)

	c.JSON(200, gin.H{"Menu telah berhasil dihapus" : &menu})
}

func BookmarkMenu(c *gin.Context){
	var bookmark model.BookmarkMenu

	if err := c.ShouldBind(&bookmark); err != nil {
		panic(err)
	}

	bookmark = model.CreateBookmarkMenu(bookmark, database.DBCon)

	c.JSON(201, gin.H{"Bookmark telah berhasil dibuat": &bookmark})
}

func DeleteBookmarkMenu(c *gin.Context){
	var bookmark model.BookmarkMenu

	if err := c.ShouldBind(&bookmark); err != nil {
		panic(err)
	}

	bookmark = model.DeleteBookmark(bookmark.UserId, bookmark.MenuId, database.DBCon)

	c.JSON(200, gin.H{"Bookmark telah berhasil dihapus" : &bookmark})
}
