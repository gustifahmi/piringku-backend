package keragaman

import (
	"github.com/gin-gonic/gin"
	"github.com/piringku-web/piringku-backend/database"
	"github.com/piringku-web/piringku-backend/model"
	"strconv"
	"time"
)

func CreateNewKeragaman(c *gin.Context){
	var keragaman model.KeragamanMakanan

	if err := c.ShouldBind(&keragaman); err != nil {
		panic(err)
	}

	keragaman = model.CreateKeragaman(keragaman, database.DBCon)

	c.JSON(201, gin.H{"Menu telah berhasil dibuat": &keragaman})
}

func GetKeragamanMakanan(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		panic(err)
	}

	keragaman := model.GetKeragaman(uint(idUser), database.DBCon)
	c.JSON(200, gin.H{"list": keragaman})
}

func GetKeragamanMakananBySubjek(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		panic(err)
	}

	idSubjek, err := strconv.ParseUint(c.Param("idSubjek"), 10, 64)
	if err != nil{
		panic(err)
	}

	keragaman := model.GetKeragamanBySubjek(uint(idUser), uint(idSubjek), database.DBCon)
	c.JSON(200, gin.H{"list": keragaman})
}

func GetKeragamanHariIni(c *gin.Context) {
	idDataDiri, err := strconv.ParseUint(c.Param("idDataDiri"), 10, 64)
	if err != nil{
		panic(err)
	}

	var keragaman model.KeragamanMakanan
	err = database.DBCon.Order("tanggal desc").Where("id_subjek = ?", idDataDiri).First(&keragaman).Error
	if err != nil {
		c.JSON(200, false)
	} else {
		if keragaman.Tanggal.Year() == time.Now().Year() && keragaman.Tanggal.Month() == time.Now().Month() &&
			keragaman.Tanggal.Day() == time.Now().Day() {
			c.JSON(200, true)
		} else {
			c.JSON(200, false)
		}
	}
}