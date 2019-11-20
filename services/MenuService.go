package services

import (
	"../config"
	"../models"
)

func GetOpcionesDeMenu(IdColaborador string) ([]models.MenuModel, error) {

	var result []models.MenuModel

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_GetOpcionesDeMenu ?", IdColaborador).Scan(&result)

	return result, nil
}
