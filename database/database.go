package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/piringku-web/piringku-backend/model"
	"os"
)

//const (
//	host     = "localhost"
//	port     = "5432"
//	user     = "postgres"
//	password = ""
//	dbname   = "piringku"
//)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//DBCon, err = gorm.Open("postgres", psqlInfo)
	DBCon, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	model.InitTableBookmark(DBCon)
	model.InitTableDataDiri(DBCon)
	model.InitTableKeragamanMakanan(DBCon)
	model.InitTableMakanan(DBCon)
	model.InitTableMenu(DBCon)
	model.InitTableRole(DBCon)
	model.InitTableUser(DBCon)
}
