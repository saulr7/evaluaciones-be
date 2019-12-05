package services

import (
	"../config"
	"../models"
)

func GetGetAreasConSuGrado() ([]models.Areas, error) {

	var result []models.Areas

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetAreasConSuGrado").Scan(&result)

	return result, nil
}

func UpdateAreaGrado(areaGrado models.AreaGrado) ([]models.AreaGrado, error) {

	var result []models.AreaGrado

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_InsertGradosPorArea ?, ?", areaGrado.IdArea, areaGrado.IdGrado).Scan(&result)

	return result, nil
}

func GetAreas() ([]models.AreaModel, error) {

	var result []models.AreaModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_GetAreas").Scan(&result)

	return result, nil
}
