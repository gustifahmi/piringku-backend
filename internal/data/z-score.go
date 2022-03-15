package data

import (
	"github.com/gin-gonic/gin"
	"github.com/piringku-web/piringku-backend/database"
)

func GetBMIForAge(c *gin.Context){
	var sumZ			float64
	var sumPercentile	float64
	var totalCount		int

	//Seluruh Z-Score
	database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Count(&totalCount)

	err := database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Select("sum(bmi_for_age_z_score)").Row().Scan(&sumZ)

	var totalZ float64
	if err != nil {
		totalZ = 0
	} else {
		totalZ = sumZ / float64(totalCount)
	}

	err = database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Select("sum(bmi_for_age_percentile)").Row().Scan(&sumPercentile)

	var totalPercentile float64
	if err != nil {
		totalPercentile = 0
	} else {
		totalPercentile = sumPercentile / float64(totalCount)
	}

	//Z-Score -3
	var minusThreeCount int
	var minusThreeZ float64
	var minusThreePercentile float64

	database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Count(&minusThreeCount)

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Select("sum(bmi_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusThreeZ = 0
	} else {
		minusThreeZ = sumZ / float64(minusThreeCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Select("sum(bmi_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusThreePercentile = 0
	} else {
		minusThreePercentile = sumPercentile / float64(minusThreeCount)
	}

	//Z-Score -2
	var minusTwoCount int
	var minusTwoZ float64
	var minusTwoPercentile float64

	database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Count(&minusTwoCount)

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Select("sum(bmi_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusTwoZ = 0
	} else {
		minusTwoZ = sumZ / float64(minusTwoCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Select("sum(bmi_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusTwoPercentile = 0
	} else {
		minusTwoPercentile = sumPercentile / float64(minusTwoCount)
	}

	//Z-Score -1
	var minusOneCount int
	var minusOneZ float64
	var minusOnePercentile float64

	database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Count(&minusOneCount)

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Select("sum(bmi_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusOneZ = 0
	} else {
		minusOneZ = sumZ / float64(minusOneCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Select("sum(bmi_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusOnePercentile = 0
	} else {
		minusOnePercentile = sumPercentile / float64(minusOneCount)
	}

	//Z-Score -1 sampai 1
	var normalCount int
	var normalZ float64
	var normalPercentile float64

	database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Count(&normalCount)

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Select("sum(bmi_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		normalZ = 0
	} else {
		normalZ = sumZ / float64(normalCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score > ? AND bmi_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Select("sum(bmi_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		normalPercentile = 0
	} else {
		normalPercentile = sumPercentile / float64(normalCount)
	}

	//Z-Score 1
	var plusOneCount int
	var plusOneZ float64
	var plusOnePercentile float64

	database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND bmi_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Count(&plusOneCount)

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND bmi_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Select("sum(bmi_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusOneZ = 0
	} else {
		plusOneZ = sumZ / float64(plusOneCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND bmi_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Select("sum(bmi_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusOnePercentile = 0
	} else {
		plusOnePercentile = sumPercentile / float64(plusOneCount)
	}

	//Z-Score 2
	var plusTwoCount int
	var plusTwoZ float64
	var plusTwoPercentile float64

	database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND bmi_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Count(&plusTwoCount)

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND bmi_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Select("sum(bmi_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusTwoZ = 0
	} else {
		plusTwoZ = sumZ / float64(plusTwoCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND bmi_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Select("sum(bmi_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusTwoPercentile = 0
	} else {
		plusTwoPercentile = sumPercentile / float64(plusTwoCount)
	}

	//Z-Score 3
	var plusThreeCount int
	var plusThreeZ float64
	var plusThreePercentile float64

	database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Count(&plusThreeCount)

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Select("sum(bmi_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusThreeZ = 0
	} else {
		plusThreeZ = sumZ / float64(plusThreeCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("bmi_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Select("sum(bmi_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusThreePercentile = 0
	} else {
		plusThreePercentile = sumPercentile / float64(plusThreeCount)
	}

	c.JSON(200, gin.H{
		"total_count":				totalCount,
		"total_z":					totalZ,
		"total_percentile":			totalPercentile,
		"minus_three_count":		minusThreeCount,
		"minus_three_z":			minusThreeZ,
		"minus_three_percentile":	minusThreePercentile,
		"minus_two_count":			minusTwoCount,
		"minus_two_z":				minusTwoZ,
		"minus_two_percentile":		minusTwoPercentile,
		"minus_one_count":			minusOneCount,
		"minus_one_z":				minusOneZ,
		"minus_one_percentile":		minusOnePercentile,
		"normal_count":				normalCount,
		"normal_z":					normalZ,
		"normal_percentile":		normalPercentile,
		"plus_one_count":			plusOneCount,
		"plus_one_z":				plusOneZ,
		"plus_one_percentile":		plusOnePercentile,
		"plus_two_count":			plusTwoCount,
		"plus_two_z":				plusTwoZ,
		"plus_two_percentile":		plusTwoPercentile,
		"plus_three_count":			plusThreeCount,
		"plus_three_z":				plusThreeZ,
		"plus_three_percentile":	plusThreePercentile,
	})
}

func GetWeightForAge(c *gin.Context){
	var sumZ			float64
	var sumPercentile	float64
	var totalCount		int

	//Seluruh Z-Score
	database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Count(&totalCount)

	err := database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Select("sum(weight_for_age_z_score)").Row().Scan(&sumZ)

	var totalZ float64
	if err != nil {
		totalZ = 0
	} else {
		totalZ = sumZ / float64(totalCount)
	}

	err = database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Select("sum(weight_for_age_percentile)").Row().Scan(&sumPercentile)

	var totalPercentile float64
	if err != nil {
		totalPercentile = 0
	} else {
		totalPercentile = sumPercentile / float64(totalCount)
	}

	//Z-Score -3
	var minusThreeCount int
	var minusThreeZ float64
	var minusThreePercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Count(&minusThreeCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Select("sum(weight_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusThreeZ = 0
	} else {
		minusThreeZ = sumZ / float64(minusThreeCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Select("sum(weight_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusThreePercentile = 0
	} else {
		minusThreePercentile = sumPercentile / float64(minusThreeCount)
	}

	//Z-Score -2
	var minusTwoCount int
	var minusTwoZ float64
	var minusTwoPercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Count(&minusTwoCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Select("sum(weight_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusTwoZ = 0
	} else {
		minusTwoZ = sumZ / float64(minusTwoCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Select("sum(weight_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusTwoPercentile = 0
	} else {
		minusTwoPercentile = sumPercentile / float64(minusTwoCount)
	}

	//Z-Score -1
	var minusOneCount int
	var minusOneZ float64
	var minusOnePercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Count(&minusOneCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Select("sum(weight_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusOneZ = 0
	} else {
		minusOneZ = sumZ / float64(minusOneCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Select("sum(weight_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusOnePercentile = 0
	} else {
		minusOnePercentile = sumPercentile / float64(minusOneCount)
	}

	//Z-Score -1 sampai 1
	var normalCount int
	var normalZ float64
	var normalPercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Count(&normalCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Select("sum(weight_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		normalZ = 0
	} else {
		normalZ = sumZ / float64(normalCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score > ? AND weight_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Select("sum(weight_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		normalPercentile = 0
	} else {
		normalPercentile = sumPercentile / float64(normalCount)
	}

	//Z-Score 1
	var plusOneCount int
	var plusOneZ float64
	var plusOnePercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND weight_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Count(&plusOneCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND weight_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Select("sum(weight_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusOneZ = 0
	} else {
		plusOneZ = sumZ / float64(plusOneCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND weight_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Select("sum(weight_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusOnePercentile = 0
	} else {
		plusOnePercentile = sumPercentile / float64(plusOneCount)
	}

	//Z-Score 2
	var plusTwoCount int
	var plusTwoZ float64
	var plusTwoPercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND weight_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Count(&plusTwoCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND weight_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Select("sum(weight_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusTwoZ = 0
	} else {
		plusTwoZ = sumZ / float64(plusTwoCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND weight_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Select("sum(weight_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusTwoPercentile = 0
	} else {
		plusTwoPercentile = sumPercentile / float64(plusTwoCount)
	}

	//Z-Score 3
	var plusThreeCount int
	var plusThreeZ float64
	var plusThreePercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Count(&plusThreeCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Select("sum(weight_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusThreeZ = 0
	} else {
		plusThreeZ = sumZ / float64(plusThreeCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Select("sum(weight_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusThreePercentile = 0
	} else {
		plusThreePercentile = sumPercentile / float64(plusThreeCount)
	}

	c.JSON(200, gin.H{
		"total_count":				totalCount,
		"total_z":					totalZ,
		"total_percentile":			totalPercentile,
		"minus_three_count":		minusThreeCount,
		"minus_three_z":			minusThreeZ,
		"minus_three_percentile":	minusThreePercentile,
		"minus_two_count":			minusTwoCount,
		"minus_two_z":				minusTwoZ,
		"minus_two_percentile":		minusTwoPercentile,
		"minus_one_count":			minusOneCount,
		"minus_one_z":				minusOneZ,
		"minus_one_percentile":		minusOnePercentile,
		"normal_count":				normalCount,
		"normal_z":					normalZ,
		"normal_percentile":		normalPercentile,
		"plus_one_count":			plusOneCount,
		"plus_one_z":				plusOneZ,
		"plus_one_percentile":		plusOnePercentile,
		"plus_two_count":			plusTwoCount,
		"plus_two_z":				plusTwoZ,
		"plus_two_percentile":		plusTwoPercentile,
		"plus_three_count":			plusThreeCount,
		"plus_three_z":				plusThreeZ,
		"plus_three_percentile":	plusThreePercentile,
	})
}

func GetHeightForAge(c *gin.Context){
	var sumZ			float64
	var sumPercentile	float64
	var totalCount		int

	//Seluruh Z-Score
	database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Count(&totalCount)

	err := database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Select("sum(height_for_age_z_score)").Row().Scan(&sumZ)

	var totalZ float64
	if err != nil {
		totalZ = 0
	} else {
		totalZ = sumZ / float64(totalCount)
	}

	err = database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Select("sum(height_for_age_percentile)").Row().Scan(&sumPercentile)

	var totalPercentile float64
	if err != nil {
		totalPercentile = 0
	} else {
		totalPercentile = sumPercentile / float64(totalCount)
	}

	//Z-Score -3
	var minusThreeCount int
	var minusThreeZ float64
	var minusThreePercentile float64

	database.DBCon.Table("data_diris").
		Where("height_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Count(&minusThreeCount)

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Select("sum(height_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusThreeZ = 0
	} else {
		minusThreeZ = sumZ / float64(minusThreeCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score <= ? AND family_role = ?", -3, "Anak").
		Select("sum(height_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusThreePercentile = 0
	} else {
		minusThreePercentile = sumPercentile / float64(minusThreeCount)
	}

	//Z-Score -2
	var minusTwoCount int
	var minusTwoZ float64
	var minusTwoPercentile float64

	database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Count(&minusTwoCount)

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Select("sum(height_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusTwoZ = 0
	} else {
		minusTwoZ = sumZ / float64(minusTwoCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Select("sum(height_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusTwoPercentile = 0
	} else {
		minusTwoPercentile = sumPercentile / float64(minusTwoCount)
	}

	//Z-Score -1
	var minusOneCount int
	var minusOneZ float64
	var minusOnePercentile float64

	database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Count(&minusOneCount)

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Select("sum(height_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusOneZ = 0
	} else {
		minusOneZ = sumZ / float64(minusOneCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Select("sum(height_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusOnePercentile = 0
	} else {
		minusOnePercentile = sumPercentile / float64(minusOneCount)
	}

	//Z-Score -1 sampai 1
	var normalCount int
	var normalZ float64
	var normalPercentile float64

	database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Count(&normalCount)

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Select("sum(height_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		normalZ = 0
	} else {
		normalZ = sumZ / float64(normalCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score > ? AND height_for_age_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Select("sum(height_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		normalPercentile = 0
	} else {
		normalPercentile = sumPercentile / float64(normalCount)
	}

	//Z-Score 1
	var plusOneCount int
	var plusOneZ float64
	var plusOnePercentile float64

	database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND height_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Count(&plusOneCount)

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND height_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Select("sum(height_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusOneZ = 0
	} else {
		plusOneZ = sumZ / float64(plusOneCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND height_for_age_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Select("sum(height_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusOnePercentile = 0
	} else {
		plusOnePercentile = sumPercentile / float64(plusOneCount)
	}

	//Z-Score 2
	var plusTwoCount int
	var plusTwoZ float64
	var plusTwoPercentile float64

	database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND height_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Count(&plusTwoCount)

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND height_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Select("sum(height_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusTwoZ = 0
	} else {
		plusTwoZ = sumZ / float64(plusTwoCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND height_for_age_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Select("sum(height_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusTwoPercentile = 0
	} else {
		plusTwoPercentile = sumPercentile / float64(plusTwoCount)
	}

	//Z-Score 3
	var plusThreeCount int
	var plusThreeZ float64
	var plusThreePercentile float64

	database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Count(&plusThreeCount)

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Select("sum(height_for_age_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusThreeZ = 0
	} else {
		plusThreeZ = sumZ / float64(plusThreeCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("height_for_age_z_score >= ? AND family_role = ?", 3, "Anak").
		Select("sum(height_for_age_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusThreePercentile = 0
	} else {
		plusThreePercentile = sumPercentile / float64(plusThreeCount)
	}

	c.JSON(200, gin.H{
		"total_count":				totalCount,
		"total_z":					totalZ,
		"total_percentile":			totalPercentile,
		"minus_three_count":		minusThreeCount,
		"minus_three_z":			minusThreeZ,
		"minus_three_percentile":	minusThreePercentile,
		"minus_two_count":			minusTwoCount,
		"minus_two_z":				minusTwoZ,
		"minus_two_percentile":		minusTwoPercentile,
		"minus_one_count":			minusOneCount,
		"minus_one_z":				minusOneZ,
		"minus_one_percentile":		minusOnePercentile,
		"normal_count":				normalCount,
		"normal_z":					normalZ,
		"normal_percentile":		normalPercentile,
		"plus_one_count":			plusOneCount,
		"plus_one_z":				plusOneZ,
		"plus_one_percentile":		plusOnePercentile,
		"plus_two_count":			plusTwoCount,
		"plus_two_z":				plusTwoZ,
		"plus_two_percentile":		plusTwoPercentile,
		"plus_three_count":			plusThreeCount,
		"plus_three_z":				plusThreeZ,
		"plus_three_percentile":	plusThreePercentile,
	})
}

func GetWeightForHeight(c *gin.Context){
	var sumZ			float64
	var sumPercentile	float64
	var totalCount		int

	//Seluruh Z-Score
	database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Count(&totalCount)

	err := database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Select("sum(weight_for_height_z_score)").Row().Scan(&sumZ)

	var totalZ float64
	if err != nil {
		totalZ = 0
	} else {
		totalZ = sumZ / float64(totalCount)
	}

	err = database.DBCon.Table("data_diris").Where("family_role = ?", "Anak").
		Select("sum(weight_for_height_percentile)").Row().Scan(&sumPercentile)

	var totalPercentile float64
	if err != nil {
		totalPercentile = 0
	} else {
		totalPercentile = sumPercentile / float64(totalCount)
	}

	//Z-Score -3
	var minusThreeCount int
	var minusThreeZ float64
	var minusThreePercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score <= ? AND family_role = ?", -3, "Anak").
		Count(&minusThreeCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score <= ? AND family_role = ?", -3, "Anak").
		Select("sum(weight_for_height_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusThreeZ = 0
	} else {
		minusThreeZ = sumZ / float64(minusThreeCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score <= ? AND family_role = ?", -3, "Anak").
		Select("sum(weight_for_height_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusThreePercentile = 0
	} else {
		minusThreePercentile = sumPercentile / float64(minusThreeCount)
	}

	//Z-Score -2
	var minusTwoCount int
	var minusTwoZ float64
	var minusTwoPercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Count(&minusTwoCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Select("sum(weight_for_height_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusTwoZ = 0
	} else {
		minusTwoZ = sumZ / float64(minusTwoCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score <= ? AND family_role = ?",
			-3, -2, "Anak").
		Select("sum(weight_for_height_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusTwoPercentile = 0
	} else {
		minusTwoPercentile = sumPercentile / float64(minusTwoCount)
	}

	//Z-Score -1
	var minusOneCount int
	var minusOneZ float64
	var minusOnePercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Count(&minusOneCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Select("sum(weight_for_height_z_score)").Row().Scan(&sumZ)

	if err != nil {
		minusOneZ = 0
	} else {
		minusOneZ = sumZ / float64(minusOneCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score <= ? AND family_role = ?",
			-2, -1, "Anak").
		Select("sum(weight_for_height_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		minusOnePercentile = 0
	} else {
		minusOnePercentile = sumPercentile / float64(minusOneCount)
	}

	//Z-Score -1 sampai 1
	var normalCount int
	var normalZ float64
	var normalPercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Count(&normalCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Select("sum(weight_for_height_z_score)").Row().Scan(&sumZ)

	if err != nil {
		normalZ = 0
	} else {
		normalZ = sumZ / float64(normalCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score > ? AND weight_for_height_z_score < ? AND family_role = ?",
			-1, 1, "Anak").
		Select("sum(weight_for_height_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		normalPercentile = 0
	} else {
		normalPercentile = sumPercentile / float64(normalCount)
	}

	//Z-Score 1
	var plusOneCount int
	var plusOneZ float64
	var plusOnePercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND weight_for_height_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Count(&plusOneCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND weight_for_height_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Select("sum(weight_for_height_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusOneZ = 0
	} else {
		plusOneZ = sumZ / float64(plusOneCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND weight_for_height_z_score < ? AND family_role = ?",
			1, 2, "Anak").
		Select("sum(weight_for_height_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusOnePercentile = 0
	} else {
		plusOnePercentile = sumPercentile / float64(plusOneCount)
	}

	//Z-Score 2
	var plusTwoCount int
	var plusTwoZ float64
	var plusTwoPercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND weight_for_height_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Count(&plusTwoCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND weight_for_height_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Select("sum(weight_for_height_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusTwoZ = 0
	} else {
		plusTwoZ = sumZ / float64(plusTwoCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND weight_for_height_z_score < ? AND family_role = ?",
			2, 3, "Anak").
		Select("sum(weight_for_height_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusTwoPercentile = 0
	} else {
		plusTwoPercentile = sumPercentile / float64(plusTwoCount)
	}

	//Z-Score 3
	var plusThreeCount int
	var plusThreeZ float64
	var plusThreePercentile float64

	database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND family_role = ?", 3, "Anak").
		Count(&plusThreeCount)

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND family_role = ?", 3, "Anak").
		Select("sum(weight_for_height_z_score)").Row().Scan(&sumZ)

	if err != nil {
		plusThreeZ = 0
	} else {
		plusThreeZ = sumZ / float64(plusThreeCount)
	}

	err = database.DBCon.Table("data_diris").
		Where("weight_for_height_z_score >= ? AND family_role = ?", 3, "Anak").
		Select("sum(weight_for_height_percentile)").Row().Scan(&sumPercentile)

	if err != nil {
		plusThreePercentile = 0
	} else {
		plusThreePercentile = sumPercentile / float64(plusThreeCount)
	}

	c.JSON(200, gin.H{
		"total_count":				totalCount,
		"total_z":					totalZ,
		"total_percentile":			totalPercentile,
		"minus_three_count":		minusThreeCount,
		"minus_three_z":			minusThreeZ,
		"minus_three_percentile":	minusThreePercentile,
		"minus_two_count":			minusTwoCount,
		"minus_two_z":				minusTwoZ,
		"minus_two_percentile":		minusTwoPercentile,
		"minus_one_count":			minusOneCount,
		"minus_one_z":				minusOneZ,
		"minus_one_percentile":		minusOnePercentile,
		"normal_count":				normalCount,
		"normal_z":					normalZ,
		"normal_percentile":		normalPercentile,
		"plus_one_count":			plusOneCount,
		"plus_one_z":				plusOneZ,
		"plus_one_percentile":		plusOnePercentile,
		"plus_two_count":			plusTwoCount,
		"plus_two_z":				plusTwoZ,
		"plus_two_percentile":		plusTwoPercentile,
		"plus_three_count":			plusThreeCount,
		"plus_three_z":				plusThreeZ,
		"plus_three_percentile":	plusThreePercentile,
	})
}