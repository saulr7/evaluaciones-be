package services

import (
	"../config"
	"../models"
)

func ColaboradoresPendientesCompletarEvaluacion(idEvaluacion string) ([]models.ColaboradesPendientesEvaluacion, error) {

	var result []models.ColaboradesPendientesEvaluacion

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_GetColaboradoresPendientesCompletarEvaluacion ?", idEvaluacion).Scan(&result)

	return result, nil
}
