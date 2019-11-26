package services

import (
	"../config"
	"../models"
)

func GetRptCompetencias(idColaborador string) ([]models.ReporteCompetencias, error) {

	var result []models.ReporteCompetencias

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetRptCompetencias ?", idColaborador).Scan(&result)

	return result, nil
}
