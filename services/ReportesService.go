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

func RptResumenGeneralService(idEvaluacionAnual string) ([]models.RptResumenGeneralModel, error) {

	var result []models.RptResumenGeneralModel

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_RptResumenGeneral  ?", idEvaluacionAnual).Scan(&result)

	return result, nil
}

func RptResumenGeneralPorEquipoService(idEvaluacionAnual string, colaboradorId int) ([]models.RptResumenGeneralModel, error) {

	var result []models.RptResumenGeneralModel

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_RptResumenGeneralPorEquipo ?, ?", idEvaluacionAnual, colaboradorId).Scan(&result)

	return result, nil
}
