package services

import (
	"../config"
	"../models"
)

func GetCargos() ([]models.CargoModel, error) {

	var cargos []models.CargoModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw(" exec usp_GetCargos").Scan(&cargos)

	return cargos, nil
}

func GetCargoPadreYEmpresaService() ([]models.CargoPadreYEmpresaModel, error) {

	var cargos []models.CargoPadreYEmpresaModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw(" exec usp_GetCargosConPadreYEmpresa").Scan(&cargos)

	return cargos, nil
}

func NewCargoService(modelo models.NuevoCargoModel) (models.NuevoCargoModel, error) {

	var result models.NuevoCargoModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw(" exec usp_CreateCargo ?, ?, ?, ?", modelo.Cargo, modelo.CargoPadreId, modelo.AgregadoPor, modelo.AreaId).Scan(&result)

	return result, nil
}
