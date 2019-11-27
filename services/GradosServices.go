package services

import (
	"../config"
	"../models"
)

func GetGetGrados() ([]models.Grados, error) {

	var result []models.Grados

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetGrados").Scan(&result)

	return result, nil
}
