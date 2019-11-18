package services

import (
	"../config"
	"../models"
)

func GetResultadosEvaluacionesHistoricas(idColaborador string) (models.ResultadosHistorico, error) {

	var result models.ResultadosHistorico
	var colaborador models.ColaboradorInformacionCompleta
	var resultadoEvaluacion40 []models.ResultadoEvaluacion
	var resultadoEvaluacion60 []models.ResultadoEvaluacion

	db4DX := config.ConnectDB4DX()
	defer db4DX.Close()

	db4DX.Raw("exec usp_ColaboradorInfoCompleta  ?", idColaborador).Scan(&colaborador)

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_GetResultadosHistoricosEvaluacionPorMeta  ?", idColaborador).Scan(&resultadoEvaluacion40)
	db.Raw("exec usp_GetResultadosHistoricosEvaluacionesGenerales  ?", idColaborador).Scan(&resultadoEvaluacion60)

	result.Colaborador = colaborador
	result.ResultadoEvaluacionPorMeta40 = resultadoEvaluacion40
	result.ResultadoEvaluacionGeneral60 = resultadoEvaluacion60

	return result, nil
}
