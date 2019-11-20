package services

import (
	"../config"
	"../models"
)

func GetResultadosEvaluacionesHistoricas(idColaborador string) (models.ResultadosHistorico, error) {

	var result models.ResultadosHistorico
	var colaborador models.ColaboradorInformacionCompleta
	var resultados []models.ResultadoEvaluacion

	db4DX := config.ConnectDB4DX()
	defer db4DX.Close()

	db4DX.Raw("exec usp_ColaboradorInfoCompleta  ?", idColaborador).Scan(&colaborador)

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_GetREsultadosEvaluacionesPorColaborador  ?", idColaborador).Scan(&resultados)

	result.Colaborador = colaborador
	result.Resultados = resultados

	return result, nil
}
