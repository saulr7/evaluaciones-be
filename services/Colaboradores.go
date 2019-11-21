package services

import (
	"../config"
	"../models"
)

func GetColaboradoresEquipo(idColaborador string) ([]models.Usuario, error) {

	var result []models.Usuario

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEquipoPorLider ?", idColaborador).Scan(&result)

	return result, nil
}

func GetEquipoEvaluacion(idColaborador string, idEvaluacionAnual string) ([]models.UsuarioPorcentaje, error) {

	var result []models.UsuarioPorcentaje

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEquipoPorLider ?,?", idColaborador, idEvaluacionAnual).Scan(&result)

	return result, nil
}

func GetEquipoCajeros(idColaborador string) ([]models.UsuarioCajeros, error) {

	var result []models.UsuarioCajeros

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEquipoCajeros ?", idColaborador).Scan(&result)

	return result, nil
}
