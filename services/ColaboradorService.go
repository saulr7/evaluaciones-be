package services

import (
	"../config"
	"../models"
)

func ColaboradoresPorArea(AreaId string) ([]models.Colaborador, error) {

	var colaborador []models.Colaborador

	db := config.ConnectDBEO()
	defer db.Close()

	var codigos []int

	db.Raw("SELECT idColaborador FROM EquiposSubAreas WHERE idSubArea = ?", AreaId).Pluck("idColaborador", &codigos)

	db.Where("idColaborador IN (?)", codigos).Find(&colaborador)

	return colaborador, nil
}

// func GetColaboradoresSubArea(IdSubArea string) ([]models.Colaborador, error) {

// 	var result []models.Colaborador

// 	db := config.ConnectDBEO()
// 	defer db.Close()

// 	db.Raw("EXEC dbo.usp_dbgetColaboradoresByEquipo ?", IdSubArea).Scan(&result)

// 	return result, nil
// }

func GetAdminEvaluaciones() ([]models.Colaborador, error) {

	var result []models.Colaborador

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetAdminsEvaluaciones").Scan(&result)

	return result, nil
}

func GetColaboradorInfoCompleta(IdColaborador string) ([]models.Colaborador, error) {

	var result []models.Colaborador

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_ColaboradorInfoCompleta ?", IdColaborador).Scan(&result)

	return result, nil
}

func GetCargosGrados() ([]models.CargoGradoModel, error) {

	var result []models.CargoGradoModel

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetCargoGrado").Scan(&result)

	return result, nil
}

func UpdateCargosGradosService(modelo models.CargoGradoActualizarModel) ([]models.CargoGradoActualizarModel, error) {

	var result []models.CargoGradoActualizarModel

	db := config.ConnectDB()
	defer db.Close()

	for _, cargo := range modelo.Cargos {

		db.Raw("EXEC usp_UpdateGradoPorCargo ? , ?", modelo.IdGrado, cargo).Scan(&result)
	}

	return result, nil
}

func GetColaboradoresPorCargo(CargoId string) ([]models.Colaborador, error) {

	var result []models.Colaborador

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_GetColaboradoresPorCargo ? ", CargoId).Scan(&result)

	return result, nil
}

func GetColaboradoresPorArea(AreaId string) ([]models.Colaborador, error) {

	var result []models.Colaborador

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_GetColaboradoresPorArea ? ", AreaId).Scan(&result)

	return result, nil
}

func GetColaboradoresInfo() ([]models.ColaboradorInfo, error) {

	var result []models.ColaboradorInfo

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_GetColaboradores ").Scan(&result)

	// for _, colaborador := range result {
	// 	var usuario models.NuevoUsuarioModel
	// 	usuario.Usuario = colaborador.IdColaborador
	// 	fmt.Println(strconv.Itoa(colaborador.IdColaborador))
	// 	clave, _ := HashPassword(strconv.Itoa(colaborador.IdColaborador))
	// 	usuario.Clave = clave
	// 	usuario.ColaboradorId = colaborador.IdColaborador
	// 	usuario.AgreadoPor = 51782
	// 	usuario.CambiarClave = true

	// 	CreateUsuarioService(usuario)
	// }

	return result, nil
}

func GetColaboradoresSinUsuarioService() ([]models.ColaboradorInfo, error) {
	var result []models.ColaboradorInfo

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_GetColaboradoresSinUsuario ").Scan(&result)

	return result, nil
}

func UpdateColaboradorActivar(colaboradorId string, Activar bool, ModificadoPor int) (models.ColaboradorInfo, error) {

	var result models.ColaboradorInfo

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_ToggleActivarColaborador ?, ?, ?", colaboradorId, Activar, ModificadoPor).Scan(&result)

	return result, nil
}

func CreateColaboradorService(modelo models.NuevoColaboradorModel) (models.NuevoColaboradorModel, error) {

	var result models.NuevoColaboradorModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_CreateColaborador ?, ?, ?, ?", modelo.IdColaborador, modelo.Nombre, modelo.FechaIngreso, modelo.AgregadoPor).Scan(&result)
	db.Raw("EXEC usp_UpdateColaboradorCargo ?, ?, ?", modelo.IdColaborador, modelo.IdCargo, modelo.AgregadoPor).Scan(&result)

	return result, nil
}

func GetColaboradoresCargoService() ([]models.ColaboradorCargo, error) {

	var result []models.ColaboradorCargo

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_GetColaboradoresConCargo ").Scan(&result)

	return result, nil
}
