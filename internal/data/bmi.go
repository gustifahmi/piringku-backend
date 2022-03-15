package data

import (
	"github.com/gin-gonic/gin"
	"github.com/piringku-web/piringku-backend/database"
	"math"
)

func BMIData(c *gin.Context) {
	bmiTotalAyah, countTotalAyah := TotalBMIDataAyah()
	bmiTotalAyah = math.Round(bmiTotalAyah * 10) / 10

	bmiTotalIbu, countTotalIbu := TotalBMIDataIbu()
	bmiTotalIbu = math.Round(bmiTotalIbu * 10) / 10

	bmiUnderweightAyah, countUnderweightAyah := UnderweightBMIDataAyah()
	bmiUnderweightAyah = math.Round(bmiUnderweightAyah * 10) / 10

	bmiUnderweightIbu, countUnderweightIbu := UnderweightBMIDataIbu()
	bmiUnderweightIbu = math.Round(bmiUnderweightIbu * 10) / 10

	bmiNormalAyah, countNormalAyah := NormalBMIDataAyah()
	bmiNormalAyah = math.Round(bmiNormalAyah * 10) / 10

	bmiNormalIbu, countNormalIbu := NormalBMIDataIbu()
	bmiNormalIbu = math.Round(bmiNormalIbu * 10) / 10

	bmiOverweightAyah, countOverweightAyah := OverweightBMIDataAyah()
	bmiOverweightAyah = math.Round(bmiOverweightAyah * 10) / 10

	bmiOverweightIbu, countOverweightIbu := OverweightBMIDataIbu()
	bmiOverweightIbu = math.Round(bmiOverweightIbu * 10) / 10

	bmiObeseAyah, countObeseAyah := ObeseBMIDataAyah()
	bmiObeseAyah = math.Round(bmiObeseAyah * 10) / 10

	bmiObeseIbu, countObeseIbu := ObeseBMIDataIbu()
	bmiObeseIbu = math.Round(bmiObeseIbu * 10) / 10

	c.JSON(200, gin.H{
		"bmi_total_ayah":			bmiTotalAyah,
		"count_total_ayah":			countTotalAyah,
		"bmi_total_ibu":			bmiTotalIbu,
		"count_total_ibu":			countTotalIbu,
		"bmi_underweight_ayah":		bmiUnderweightAyah,
		"count_underweight_ayah":	countUnderweightAyah,
		"bmi_underweight_ibu":		bmiUnderweightIbu,
		"count_underweight_ibu":	countUnderweightIbu,
		"bmi_normal_ayah":			bmiNormalAyah,
		"count_normal_ayah":		countNormalAyah,
		"bmi_normal_ibu":			bmiNormalIbu,
		"count_normal_ibu":			countNormalIbu,
		"bmi_overweight_ayah":		bmiOverweightAyah,
		"count_overweight_ayah":	countOverweightAyah,
		"bmi_overweight_ibu":		bmiOverweightIbu,
		"count_overweight_ibu":		countOverweightIbu,
		"bmi_obese_ayah":			bmiObeseAyah,
		"count_obese_ayah":			countObeseAyah,
		"bmi_obese_ibu":			bmiObeseIbu,
		"count_obese_ibu":			countObeseIbu,
	})
}

func TotalBMIDataAyah() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").Where("bmi > ? AND family_role = ?", 0, "Ayah").
		Count(&count)

	err := database.DBCon.Table("data_diris").Where("bmi > ? AND family_role = ?", 0, "Ayah").
		Select("sum(bmi)").Row().Scan(&sumBMI)

	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func TotalBMIDataIbu() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").Where("bmi > ? AND family_role = ?", 0, "Ibu").
		Count(&count)

	err := database.DBCon.Table("data_diris").Where("bmi > ? AND family_role = ?", 0, "Ibu").
		Select("sum(bmi)").Row().Scan(&sumBMI)

	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func UnderweightBMIDataAyah() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").
		Where("bmi > ? AND bmi < ? AND family_role = ?", 0, 18.5, "Ayah").
		Count(&count)

	err := database.DBCon.Table("data_diris").
		Where("bmi > ? AND bmi < ? AND family_role = ?", 0, 18.5, "Ayah").
		Select("sum(bmi)").Row().Scan(&sumBMI)

	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func UnderweightBMIDataIbu() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").
		Where("bmi > ? AND bmi < ? AND family_role = ?", 0, 18.5, "Ibu").
		Count(&count)

	err := database.DBCon.Table("data_diris").
		Where("bmi > ? AND bmi < ? AND family_role = ?", 0, 18.5, "Ibu").
		Select("sum(bmi)").Row().Scan(&sumBMI)

	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func NormalBMIDataAyah() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").
		Where("bmi >= ? AND bmi < ? AND family_role = ?", 18.5, 23, "Ayah").
		Count(&count)

	err := database.DBCon.Table("data_diris").
		Where("bmi >= ? AND bmi < ? AND family_role = ?", 18.5, 23, "Ayah").
		Select("sum(bmi)").Row().Scan(&sumBMI)

	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func NormalBMIDataIbu() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").
		Where("bmi >= ? AND bmi < ? AND family_role = ?", 18.5, 23, "Ibu").
		Count(&count)

	err := database.DBCon.Table("data_diris").
		Where("bmi >= ? AND bmi < ? AND family_role = ?", 18.5, 23, "Ibu").
		Select("sum(bmi)").Row().Scan(&sumBMI)

	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func OverweightBMIDataAyah() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").
		Where("bmi >= ? AND bmi < ? AND family_role = ?", 23, 25, "Ayah").
		Count(&count)

	err := database.DBCon.Table("data_diris").
		Where("bmi >= ? AND bmi < ? AND family_role = ?", 23, 25, "Ayah").
		Select("sum(bmi)").Row().Scan(&sumBMI)

	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func OverweightBMIDataIbu() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").
		Where("bmi >= ? AND bmi < ? AND family_role = ?", 23, 25, "Ibu").
		Count(&count)

	err := database.DBCon.Table("data_diris").
		Where("bmi >= ? AND bmi < ? AND family_role = ?", 23, 25, "Ibu").
		Select("sum(bmi)").Row().Scan(&sumBMI)

	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func ObeseBMIDataAyah() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").Where("bmi >= ? AND family_role = ?", 25, "Ayah").
		Count(&count)

	err := database.DBCon.Table("data_diris").Where("bmi >= ? AND family_role = ?", 25, "Ayah").
		Select("sum(bmi)").Row().Scan(&sumBMI)
	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}

func ObeseBMIDataIbu() (bmi float64, count int){
	var sumBMI float64

	database.DBCon.Table("data_diris").Where("bmi >= ? AND family_role = ?", 25, "Ibu").
		Count(&count)

	err := database.DBCon.Table("data_diris").Where("bmi >= ? AND family_role = ?", 25, "Ibu").
		Select("sum(bmi)").Row().Scan(&sumBMI)
	if err != nil {
		bmi = 0
	} else {
		bmi = sumBMI / float64(count)
	}

	return
}