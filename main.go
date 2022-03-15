package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/piringku-web/piringku-backend/database"
	"github.com/piringku-web/piringku-backend/internal/data"
	"github.com/piringku-web/piringku-backend/internal/datadiri"
	"github.com/piringku-web/piringku-backend/internal/keragaman"
	"github.com/piringku-web/piringku-backend/internal/makanan"
	"github.com/piringku-web/piringku-backend/internal/menu"
	"github.com/piringku-web/piringku-backend/internal/user"
	"os"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	database.InitDB()

	r.GET("/", index)
	bookmarkPath := r.Group("/bookmark")
	{
		bookmarkPath.POST("/create", menu.BookmarkMenu)
		bookmarkPath.DELETE("/delete", menu.DeleteBookmarkMenu)
	}
	dataPath := r.Group("/data")
	{
		dataPath.GET("/bmi/get", data.BMIData)
		dataPath.GET("/bmi-for-age/get", data.GetBMIForAge)
		dataPath.GET("/weight-for-age/get", data.GetWeightForAge)
		dataPath.GET("/height-for-age/get", data.GetHeightForAge)
		dataPath.GET("/weight-for-height/get", data.GetWeightForHeight)
	}
	dataDiriPath := r.Group("/datadiri")
	{
		dataDiriPath.POST("/create", datadiri.CreateDataDiri)
		dataDiriPath.GET("/list/:idUser", datadiri.GetDaftarDataDiri)
		dataDiriPath.GET("/get/ayah/:idUser", datadiri.GetDataDiriAyah)
		dataDiriPath.GET("/get/ibu/:idUser", datadiri.GetDataDiriIbu)
		dataDiriPath.GET("/get/anak/:idUser/:nomorUrutan", datadiri.GetAnakDataDiri)
		dataDiriPath.PUT("/update/ayah/:idUser", datadiri.UpdateDataDiriAyah)
		dataDiriPath.PUT("/update/ibu/:idUser", datadiri.UpdateDataDiriIbu)
		dataDiriPath.PUT("/update/anak/:idUser/:nomorUrutan", datadiri.UpdateDataDiriAnak)
	}
	makananPath := r.Group("/makanan")
	{
		makananPath.GET("/get", makanan.GetDaftarMakanan)
	}
	menuPath := r.Group("/menu")
	{
		menuPath.GET("/count/:kategori", menu.GetCount)
		menuPath.GET("/all/get/:page", menu.GetDaftarSemuaMenu)
		menuPath.GET("/sarapan/get/:page", menu.GetDaftarMenuSarapan)
		menuPath.GET("/makansiang/get/:page", menu.GetDaftarMenuMakanSiang)
		menuPath.GET("/makanmalam/get/:page", menu.GetDaftarMenuMakanMalam)
		menuPath.GET("/selingan/get/:page", menu.GetDaftarMenuSelingan)
		menuPath.GET("/disimpan/:idUser", menu.GetMenuDisimpan)
		menuPath.GET("/detail/:idMenu", menu.GetDetailMenu)
		menuPath.GET("/get/:idUser/:idMenu", menu.GetMenuById)
		menuPath.POST("/create", menu.CreateNewMenu)
		menuPath.PUT("/update/:idMenu", menu.UpdateMenuById)
		menuPath.DELETE("/delete/:idMenu", menu.DeleteMenuById)
	}
	userPath := r.Group("/user")
	{
		userPath.GET("/check/:idUser", user.CheckDataDiri)
		userPath.GET("/subjek/:idUser", user.GetSubjek)
		userPath.PUT("/login", user.Auth)
	}
	keragamanPath := r.Group("/keragaman")
	{
		keragamanPath.POST("/create", keragaman.CreateNewKeragaman)
		keragamanPath.GET("/get/:idUser/:idSubjek", keragaman.GetKeragamanMakananBySubjek)
		keragamanPath.GET("/get/:idUser", keragaman.GetKeragamanMakanan)
		keragamanPath.GET("/hari-ini/get/:idDataDiri", keragaman.GetKeragamanHariIni)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}

	err = database.DBCon.Close()
	if err != nil {
		panic(err)
	}
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}