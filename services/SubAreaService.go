package services

import (
	"../config"
	"../models"
)

func SubAreasAll() ([]models.SubArea, error) {

	var SubArea []models.SubArea

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw(" exec usp_GetAreas").Scan(&SubArea)

	return SubArea, nil
}
