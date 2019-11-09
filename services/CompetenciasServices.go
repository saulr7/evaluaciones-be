package services

import (
	"../config"
	"../models"
)

func GetCompetenciasPorColaborador(idColaborador string) ([]models.Competencias, error) {

	var result []models.Competencias

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetCompetenciasPorColaborador ?", idColaborador).Scan(&result)

	return result, nil
}
