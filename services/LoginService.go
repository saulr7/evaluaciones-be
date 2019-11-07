package services

import (
	"encoding/json"
	"errors"

	"../config"
	"../models"
)

func Login(credenciales models.UsuarioCredenciales) (string, error) {

	var result models.Usuario

	db := config.ConnectDB4DX()
	defer db.Close()

	db.Raw("EXEC Usp_dbAuthUser ?, ?", credenciales.CodigoEmpleado, credenciales.Password).Scan(&result)

	if result.IdColaborador == 0 {
		return "", errors.New("Credenciales incorrectas")
	}

	//var eventosistema models.EventoDelSistema
	// eventosistema.Evento = "Inicio de sesión"
	// eventosistema.IdColaborador, _ = strconv.Atoi(credenciales.CodigoEmpleado)
	// RegistarEventoDelSistema(eventosistema)

	json.Marshal(&result)

	token, _ := Create_JWT(result)

	return token, nil

}

func LoginWithToken(credenciales models.UsuarioCredenciales) (string, error) {

	var result models.Usuario

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbAuthUserWithToken ?, ?", credenciales.CodigoEmpleado, credenciales.Token).Scan(&result)

	if result.IdColaborador == 0 {
		return "", errors.New("Credenciales incorrectas")
	}

	// var eventosistema models.EventoDelSistema
	// eventosistema.Evento = "Inicio de sesión"
	// eventosistema.IdColaborador, _ = strconv.Atoi(credenciales.CodigoEmpleado)
	// RegistarEventoDelSistema(eventosistema)

	json.Marshal(&result)

	token, _ := Create_JWT(result)

	return token, nil

}
