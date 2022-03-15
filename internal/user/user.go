package user

import (
	"github.com/gin-gonic/gin"
	"github.com/piringku-web/piringku-backend/database"
	"github.com/piringku-web/piringku-backend/model"
	"strconv"
	"time"
)

func Auth(c *gin.Context){
	var user model.User
	var userTarget model.User

	err := c.ShouldBind(&user)
	if err != nil {
		panic(err)
	}

	err = database.DBCon.Where("email = ?", user.Email).First(&userTarget).Error
	if err != nil {
		user = model.CreateUser(user, database.DBCon)
	} else {
		user = model.Login(userTarget, user, database.DBCon)
	}

	c.JSON(200, user)
}

func CheckDataDiri(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil {
		panic(err)
	}
	check := true

	var dataDiriOrtu []model.DataDiri
	database.DBCon.Where("(family_role = ? OR family_role = ?) AND user_id = ?", "Ayah", "Ibu", idUser).
		Find(&dataDiriOrtu)

	if len(dataDiriOrtu) == 0 {
		check = false
	} else {
		var dataDiriAnak []model.DataDiri
		database.DBCon.Order("tanggal_lahir").Where("user_id = ? AND family_role = ?", idUser, "Anak").
			Find(&dataDiriAnak)
		if len(dataDiriAnak) == 0 {
			check = false
		}
	}

	c.JSON(200, check)
}

func GetSubjek(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil {
		panic(err)
	}

	var daftarSubjek []model.DataDiri

	var ayah model.DataDiri
	err = database.DBCon.Select("id_data_diri, nama").
		Where("user_id = ? AND family_role = ?", idUser, "Ayah").First(&ayah).Error
	if err == nil {
		daftarSubjek = append(daftarSubjek, ayah)
	}

	var ibu model.DataDiri
	err = database.DBCon.Select("id_data_diri, nama").
		Where("user_id = ? AND family_role = ?", idUser, "Ibu").First(&ibu).Error
	if err == nil {
		daftarSubjek = append(daftarSubjek, ibu)
	}

	var daftarAnak []model.DataDiri
	err = database.DBCon.Order("tanggal_lahir").Select("id_data_diri, nama").
		Where("user_id = ? AND family_role = ? AND tanggal_lahir > ?",
			idUser, "Anak", time.Now().AddDate(-8, 0, 0 )).Find(&daftarAnak).Error
	if err == nil {
		for _, anak := range daftarAnak {
			daftarSubjek = append(daftarSubjek, anak)
		}
	}

	c.JSON(200, gin.H{"subjek": daftarSubjek})
}