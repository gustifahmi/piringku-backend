package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DataDiri struct {
	IdDataDiri					uint	`gorm:"primary_key"`
	UserId						uint
	TanggalPengukuran			time.Time
	Nama						string
	FamilyRole					string
	JenisKelamin				string
	TanggalLahir				time.Time
	Provinsi					string
	Kota						string
	Kecamatan					string
	Tinggi						float64
	Berat						float64
	Posisi						string
	Odema						bool
	BMI							float64
	BMIForAgeZScore				float64
	BMIForAgePercentile			float64
	WeightForAgeZScore			float64
	WeightForAgePercentile		float64
	HeightForAgeZScore			float64
	HeightForAgePercentile		float64
	WeightForHeightZScore		float64
	WeightForHeightPercentile	float64
	WFHIsNull					bool
}

func InitTableDataDiri(DBCon *gorm.DB) {
	DBCon.AutoMigrate(&DataDiri{})
	DBCon.Model(&DataDiri{}).AddForeignKey("user_id", "users(id_user)", "CASCADE", "CASCADE")
}

func CreateDataDiri(dataDiri DataDiri, db *gorm.DB) DataDiri {
	err := db.Create(&dataDiri).Error
	if err != nil {
		panic(err)
	}

	return dataDiri
}

func GetDataDiriListByIdUser(idUser uint, db *gorm.DB) []DataDiri {
	var daftarDataDiri []DataDiri

	var dataDiriAyah DataDiri
	db.Where("user_id = ? AND family_role = ?", idUser, "Ayah").First(&dataDiriAyah)
	daftarDataDiri = append(daftarDataDiri, dataDiriAyah)

	var dataDiriIbu DataDiri
	db.Where("user_id = ? AND family_role = ?", idUser, "Ibu").First(&dataDiriIbu)
	daftarDataDiri = append(daftarDataDiri, dataDiriIbu)

	var dataDiriAnak []DataDiri
	db.Order("tanggal_lahir").
		Where("user_id = ? AND family_role = ? AND tanggal_lahir > ?",
			idUser, "Anak", time.Now().AddDate(-8, 0, 0 )).
		Find(&dataDiriAnak)

	for _, dataDiri := range dataDiriAnak {
		daftarDataDiri = append(daftarDataDiri, dataDiri)
	}

	return daftarDataDiri
}

func GetDataDiriOrtu(role string, idUser uint, db *gorm.DB) DataDiri {
	var dataDiri DataDiri

	err := db.Where("family_role = ? AND user_id = ?", role, idUser).First(&dataDiri).Error
	if err != nil {
		panic(err)
	}

	return dataDiri
}

func GetDataDiriAnak(idUser uint, nomorUrutan int, db *gorm.DB) DataDiri {
	var dataDiriAnak []DataDiri
	db.Order("tanggal_lahir").
		Where("user_id = ? AND family_role = ? AND tanggal_lahir > ?",
			idUser, "Anak", time.Now().AddDate(-8, 0, 0 )).
		Find(&dataDiriAnak)

	dataDiri := dataDiriAnak[nomorUrutan - 1]

	return dataDiri
}

func UpdateDataDiri(dataDiriLama DataDiri, dataDiriBaru DataDiri, db *gorm.DB) DataDiri {
	err := db.Model(&dataDiriLama).Updates(&dataDiriBaru).Error
	if err != nil{
		panic(err)
	}

	if !dataDiriBaru.WFHIsNull {
		db.Model(&dataDiriLama).Update("wfh_is_null", false)
	}

	if !dataDiriBaru.Odema {
		db.Model(&dataDiriLama).Update("odema", false)
	}

	return dataDiriBaru
}