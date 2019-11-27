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
