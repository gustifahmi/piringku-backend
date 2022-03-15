package datadiri

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/piringku-web/piringku-backend/database"
	"github.com/piringku-web/piringku-backend/model"
	"math"
	"strconv"
	"time"
)

func CreateDataDiri(c *gin.Context){
	var dataDiri model.DataDiri

	if err := c.ShouldBind(&dataDiri); err != nil {
		panic(err)
	}
	var familyRole = dataDiri.FamilyRole

	if familyRole == "Ayah" || familyRole == "Ibu" {
		var existed model.DataDiri
		err := database.DBCon.Where("user_id = ? AND family_role = ?", dataDiri.UserId, familyRole).First(&existed).Error
		if err != nil {
			dataDiriWithBMI := model.DataDiri{
				UserId:				dataDiri.UserId,
				Nama:				dataDiri.Nama,
				FamilyRole:			dataDiri.FamilyRole,
				JenisKelamin:		dataDiri.JenisKelamin,
				TanggalLahir:		dataDiri.TanggalLahir,
				Tinggi:				dataDiri.Tinggi,
				Berat:				dataDiri.Berat,
				Provinsi:			dataDiri.Provinsi,
				Kota:				dataDiri.Kota,
				Kecamatan:			dataDiri.Kecamatan,
				BMI:				dataDiri.Berat / math.Pow(dataDiri.Tinggi/100, 2),
			}

			dataDiriWithBMI = model.CreateDataDiri(dataDiriWithBMI, database.DBCon)

			c.JSON(201, gin.H{"Data diri telah berhasil dibuat": &dataDiriWithBMI})
		} else {
			c.JSON(401, "Tidak dapat menambah data diri")
		}
	} else if familyRole == "Anak" {
		var dataDiriAnak []model.DataDiri
		err := database.DBCon.Where("user_id = ? AND family_role = ? AND tanggal_lahir > ?",
			dataDiri.UserId, "Anak", time.Now().AddDate(-8, 0, 0)).
			Find(&dataDiriAnak).Error
		if err != nil {
			panic(err)
		}

		if len(dataDiriAnak) < 3 {
			dataDiriWithBMI := ProcessDataDiriAnak(dataDiri)

			dataDiriWithBMI = model.CreateDataDiri(dataDiriWithBMI, database.DBCon)

			c.JSON(201, gin.H{"Data diri telah berhasil dibuat": &dataDiriWithBMI})
		} else {
			c.JSON(401, "Tidak dapat menambah data diri")
		}
	}
}

func GetDaftarDataDiri(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		panic(err)
	}

	daftarDataDiri := model.GetDataDiriListByIdUser(uint(idUser),database.DBCon)
	c.JSON(200, gin.H{"list": daftarDataDiri})
}

func GetDataDiriAyah(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		panic(err)
	}

	dataDiri := model.GetDataDiriOrtu("Ayah", uint(idUser),database.DBCon)

	c.JSON(200, dataDiri)
}

func GetDataDiriIbu(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		panic(err)
	}

	dataDiri := model.GetDataDiriOrtu("Ibu", uint(idUser),database.DBCon)

	c.JSON(200, dataDiri)
}

func GetAnakDataDiri(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		panic(err)
	}

	nomorUrutan, err := strconv.ParseUint(c.Param("nomorUrutan"), 10, 64)
	if err != nil{
		panic(err)
	}

	dataDiri := model.GetDataDiriAnak(uint(idUser), int(nomorUrutan), database.DBCon)

	c.JSON(200, dataDiri)
}

func UpdateDataDiriAyah(c *gin.Context){
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil {
		panic(err)
	}

	var dataDiriBaru model.DataDiri

	if err := c.ShouldBind(&dataDiriBaru); err != nil{
		panic(err)
	}
	dataDiriWithBMI := model.DataDiri{
		UserId:				dataDiriBaru.UserId,
		Nama:				dataDiriBaru.Nama,
		FamilyRole:			dataDiriBaru.FamilyRole,
		JenisKelamin:		dataDiriBaru.JenisKelamin,
		TanggalLahir:		dataDiriBaru.TanggalLahir,
		Tinggi:				dataDiriBaru.Tinggi,
		Berat:				dataDiriBaru.Berat,
		Provinsi:			dataDiriBaru.Provinsi,
		Kota:				dataDiriBaru.Kota,
		Kecamatan:			dataDiriBaru.Kecamatan,
		BMI:				dataDiriBaru.Berat / math.Pow(dataDiriBaru.Tinggi/100, 2),
	}
	var dataDiriLama model.DataDiri

	err = database.DBCon.Where("user_id = ? AND family_role = ?", idUser, "Ayah").First(&dataDiriLama).Error
	if err != nil{
		panic(err)
	}

	dataDiriWithBMI = model.UpdateDataDiri(dataDiriLama, dataDiriWithBMI, database.DBCon)

	c.JSON(200, gin.H{"Data diri telah berhasil diupdate" : &dataDiriBaru})
}

func UpdateDataDiriIbu(c *gin.Context){
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)

	if err != nil {
		panic(err)
	}

	var dataDiriBaru model.DataDiri

	if err := c.ShouldBind(&dataDiriBaru); err != nil{
		panic(err)
	}
	dataDiriWithBMI := model.DataDiri{
		UserId:				dataDiriBaru.UserId,
		Nama:				dataDiriBaru.Nama,
		FamilyRole:			dataDiriBaru.FamilyRole,
		JenisKelamin:		dataDiriBaru.JenisKelamin,
		TanggalLahir:		dataDiriBaru.TanggalLahir,
		Tinggi:				dataDiriBaru.Tinggi,
		Berat:				dataDiriBaru.Berat,
		Provinsi:			dataDiriBaru.Provinsi,
		Kota:				dataDiriBaru.Kota,
		Kecamatan:			dataDiriBaru.Kecamatan,
		BMI:				dataDiriBaru.Berat / math.Pow(dataDiriBaru.Tinggi/100, 2),
	}

	var dataDiriLama model.DataDiri
	err = database.DBCon.Where("user_id = ? AND family_role = ?", idUser, "Ibu").First(&dataDiriLama).Error
	if err != nil{
		panic(err)
	}

	dataDiriWithBMI = model.UpdateDataDiri(dataDiriLama, dataDiriWithBMI, database.DBCon)

	c.JSON(200, gin.H{"Data diri telah berhasil diupdate" : &dataDiriBaru})
}

func UpdateDataDiriAnak(c *gin.Context) {
	idUser, err := strconv.ParseUint(c.Param("idUser"), 10, 64)
	if err != nil{
		panic(err)
	}

	nomorUrutan, err := strconv.ParseUint(c.Param("nomorUrutan"), 10, 64)
	if err != nil{
		panic(err)
	}

	var dataDiriBaru model.DataDiri

	if err := c.ShouldBind(&dataDiriBaru); err != nil{
		panic(err)
	}

	dataDiriBaruProcessed := ProcessDataDiriAnak(dataDiriBaru)

	dataDiriLama := model.GetDataDiriAnak(uint(idUser), int(nomorUrutan), database.DBCon)

	dataDiriBaruProcessed = model.UpdateDataDiri(dataDiriLama, dataDiriBaruProcessed, database.DBCon)

	c.JSON(200, gin.H{"Data diri telah berhasil diupdate" : &dataDiriBaruProcessed})
}

func ProcessDataDiriAnak(dataDiri model.DataDiri) model.DataDiri {
	var tinggi = dataDiri.Tinggi
	if dataDiri.Posisi == "Berbaring" {
		tinggi -= 0.7
	}

	BMI := dataDiri.Berat / math.Pow(tinggi/100, 2)

	y1, M1, d1 := time.Now().Date()
	y2, M2, d2 := dataDiri.TanggalLahir.Date()

	usiaTahun := y1-y2
	usiaBulan := int(M1-M2)
	usiaBulan += usiaTahun * 12

	if d2 > d1 {
		usiaBulan -= 1
	}
	axisBulanB := fmt.Sprintf("B%d", usiaBulan-22)
	axisBulanC := fmt.Sprintf("C%d", usiaBulan-22)
	axisBulanD := fmt.Sprintf("D%d", usiaBulan-22)

	var bmiForAgeZ, bmiForAgePercentile float64
	var weightForAgeZ, weightForAgePercentile float64
	var heightForAgeZ, heightForAgePercentile float64
	var weightForHeightZ, weightForHeightPercentile float64
	var L, M, S float64
	wfhIsNull := true

	if dataDiri.JenisKelamin == "Laki-laki" {
		if !dataDiri.Odema {
			bmiForAgeBoysFile, err := excelize.OpenFile("xlsx/bmi-for-age-boys.xlsx")
			if err != nil {
				panic(err)
			}
			L, _ = strconv.ParseFloat(bmiForAgeBoysFile.GetCellValue("z-score", axisBulanB), 64)
			M, _ = strconv.ParseFloat(bmiForAgeBoysFile.GetCellValue("z-score", axisBulanC), 64)
			S, _ = strconv.ParseFloat(bmiForAgeBoysFile.GetCellValue("z-score", axisBulanD), 64)
			bmiForAgeZ = (math.Pow(BMI/M, L) - 1) / (L * S)
			bmiForAgePercentile = (0.5 * (1 + math.Erf(bmiForAgeZ / math.Sqrt(2)))) * 100

			weightForAgeBoysFile, err := excelize.OpenFile("xlsx/weight-for-age-boys.xlsx")
			if err != nil {
				panic(err)
			}
			L, _ = strconv.ParseFloat(weightForAgeBoysFile.GetCellValue("z-score", axisBulanB), 64)
			M, _ = strconv.ParseFloat(weightForAgeBoysFile.GetCellValue("z-score", axisBulanC), 64)
			S, _ = strconv.ParseFloat(weightForAgeBoysFile.GetCellValue("z-score", axisBulanD), 64)
			weightForAgeZ = (math.Pow(dataDiri.Berat/M, L) - 1) / (L * S)
			weightForAgePercentile = (0.5 * (1 + math.Erf(weightForAgeZ / math.Sqrt(2)))) * 100
		}

		heightForAgeBoysFile, err := excelize.OpenFile("xlsx/height-for-age-boys.xlsx")
		if err != nil {
			panic(err)
		}
		L, _ = strconv.ParseFloat(heightForAgeBoysFile.GetCellValue("z-score", axisBulanB), 64)
		M, _ = strconv.ParseFloat(heightForAgeBoysFile.GetCellValue("z-score", axisBulanC), 64)
		S, _ = strconv.ParseFloat(heightForAgeBoysFile.GetCellValue("z-score", axisBulanD), 64)
		heightForAgeZ = (math.Pow(tinggi / M, L) - 1) / (L * S)
		heightForAgePercentile = (0.5 * (1 + math.Erf(heightForAgeZ / math.Sqrt(2)))) * 100

		if usiaBulan <= 60 && tinggi > 65 && tinggi < 120 && !dataDiri.Odema {
			wfhIsNull = false
			tinggi = math.Round(tinggi * 2) / 2
			weightForHeightBoysFile, err := excelize.OpenFile("xlsx/weight-for-height-boys.xlsx")
			if err != nil {
				panic(err)
			}

			axisTinggiB := fmt.Sprintf("B%d", int(tinggi * 2-128))
			axisTinggiC := fmt.Sprintf("C%d", int(tinggi * 2-128))
			axisTinggiD := fmt.Sprintf("D%d", int(tinggi * 2-128))

			L, _ = strconv.ParseFloat(weightForHeightBoysFile.GetCellValue("z-score", axisTinggiB), 64)
			M, _ = strconv.ParseFloat(weightForHeightBoysFile.GetCellValue("z-score", axisTinggiC), 64)
			S, _ = strconv.ParseFloat(weightForHeightBoysFile.GetCellValue("z-score", axisTinggiD), 64)

			weightForHeightZ = (math.Pow(dataDiri.Berat / M, L) - 1) / (L * S)
			weightForHeightPercentile = (0.5 * (1 + math.Erf(weightForHeightZ / math.Sqrt(2)))) * 100
		}
	} else if dataDiri.JenisKelamin == "Perempuan" {
		if !dataDiri.Odema {
			bmiForAgeGirlsFile, err := excelize.OpenFile("xlsx/bmi-for-age-girls.xlsx")
			if err != nil {
				panic(err)
			}
			L, _ = strconv.ParseFloat(bmiForAgeGirlsFile.GetCellValue("z-score", axisBulanB), 64)
			M, _ = strconv.ParseFloat(bmiForAgeGirlsFile.GetCellValue("z-score", axisBulanC), 64)
			S, _ = strconv.ParseFloat(bmiForAgeGirlsFile.GetCellValue("z-score", axisBulanD), 64)
			bmiForAgeZ = (math.Pow(BMI / M, L) - 1) / (L * S)
			bmiForAgePercentile = (0.5 * (1 + math.Erf(bmiForAgeZ / math.Sqrt(2)))) * 100

			weightForAgeGirlsFile, err := excelize.OpenFile("xlsx/weight-for-age-girls.xlsx")
			if err != nil {
				panic(err)
			}
			L, _ = strconv.ParseFloat(weightForAgeGirlsFile.GetCellValue("z-score", axisBulanB), 64)
			M, _ = strconv.ParseFloat(weightForAgeGirlsFile.GetCellValue("z-score", axisBulanC), 64)
			S, _ = strconv.ParseFloat(weightForAgeGirlsFile.GetCellValue("z-score", axisBulanD), 64)
			weightForAgeZ = (math.Pow(dataDiri.Berat / M, L) - 1) / (L * S)
			weightForAgePercentile = (0.5 * (1 + math.Erf(weightForAgeZ / math.Sqrt(2)))) * 100
		}

		heightForAgeGirlsFile, err := excelize.OpenFile("xlsx/height-for-age-girls.xlsx")
		if err != nil {
			panic(err)
		}
		L, _ = strconv.ParseFloat(heightForAgeGirlsFile.GetCellValue("z-score", axisBulanB), 64)
		M, _ = strconv.ParseFloat(heightForAgeGirlsFile.GetCellValue("z-score", axisBulanC), 64)
		S, _ = strconv.ParseFloat(heightForAgeGirlsFile.GetCellValue("z-score", axisBulanD), 64)
		heightForAgeZ = (math.Pow(tinggi / M, L) - 1) / (L * S)
		heightForAgePercentile = (0.5 * (1 + math.Erf(heightForAgeZ / math.Sqrt(2)))) * 100

		if usiaBulan <= 60 && tinggi > 65 && tinggi < 120 && !dataDiri.Odema {
			wfhIsNull = false
			tinggi = math.Round(tinggi * 2) / 2
			weightForHeightGirlsFile, err := excelize.OpenFile("xlsx/weight-for-height-girls.xlsx")
			if err != nil {
				panic(err)
			}

			axisTinggiB := fmt.Sprintf("B%d", int(tinggi * 2-128))
			axisTinggiC := fmt.Sprintf("C%d", int(tinggi * 2-128))
			axisTinggiD := fmt.Sprintf("D%d", int(tinggi * 2-128))

			L, _ = strconv.ParseFloat(weightForHeightGirlsFile.GetCellValue("z-score", axisTinggiB), 64)
			M, _ = strconv.ParseFloat(weightForHeightGirlsFile.GetCellValue("z-score", axisTinggiC), 64)
			S, _ = strconv.ParseFloat(weightForHeightGirlsFile.GetCellValue("z-score", axisTinggiD), 64)

			weightForHeightZ = (math.Pow(dataDiri.Berat / M, L) - 1) / (L * S)
			weightForHeightPercentile = (0.5 * (1 + math.Erf(weightForHeightZ / math.Sqrt(2)))) * 100
		}
	}

	dataDiriWithBMI := model.DataDiri{
		UserId:						dataDiri.UserId,
		TanggalPengukuran:			dataDiri.TanggalPengukuran,
		Nama:						dataDiri.Nama,
		FamilyRole:					dataDiri.FamilyRole,
		JenisKelamin:				dataDiri.JenisKelamin,
		TanggalLahir:				dataDiri.TanggalLahir,
		Tinggi:						dataDiri.Tinggi,
		Berat:						dataDiri.Berat,
		Provinsi:					dataDiri.Provinsi,
		Kota:						dataDiri.Kota,
		Kecamatan:					dataDiri.Kecamatan,
		Odema:						dataDiri.Odema,
		Posisi:						dataDiri.Posisi,
		BMIForAgeZScore:			bmiForAgeZ,
		BMIForAgePercentile:		bmiForAgePercentile,
		WeightForAgeZScore:			weightForAgeZ,
		WeightForAgePercentile:		weightForAgePercentile,
		HeightForAgeZScore:			heightForAgeZ,
		HeightForAgePercentile:		heightForAgePercentile,
		WeightForHeightZScore:		weightForHeightZ,
		WeightForHeightPercentile:	weightForHeightPercentile,
		WFHIsNull:					wfhIsNull,
	}

	return dataDiriWithBMI
}