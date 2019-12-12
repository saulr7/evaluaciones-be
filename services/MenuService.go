package services

import (
	"../config"
	"../models"
)

func GetOpcionesDeMenuByUserService(IdColaborador int) ([]models.MenuModel, error) {

	var result []models.MenuModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("exec usp_GetOpcionesDeMenu ?", IdColaborador).Scan(&result)

	return result, nil
}

func GetOpcionesDeMenu(IdColaborador int) ([]models.MenuModel, error) {

	var result []models.MenuModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("exec usp_GetOpcionesDeMenu ?", IdColaborador).Scan(&result)

	return result, nil
}
