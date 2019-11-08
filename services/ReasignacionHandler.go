package services

import (
	"../config"
	"../models"
)

func GetReasignaciones() ([]models.Reasignacion, error) {

	var result []models.Reasignacion

	db := config.ConnectDB4DX()
	defer db.Close()

	db.Raw("exec usp_GetReasignaciones").Scan(&result)

	return result, nil
}

func DeleteReasignacion(IdAsignacion string) (bool, error) {

	var result []models.Reasignacion

	db := config.ConnectDB4DX()
	defer db.Close()

	db.Raw("exec usp_RemoveAsignacion ?", IdAsignacion).Scan(&result)

	return true, nil
}
