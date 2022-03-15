package makanan

import (
	"github.com/gin-gonic/gin"
	"github.com/piringku-web/piringku-backend/database"
	"github.com/piringku-web/piringku-backend/model"
)

func GetDaftarMakanan(c *gin.Context) {
	daftarMakanan := model.GetMakananList(database.DBCon)

	c.JSON(200, daftarMakanan)
}