package services

import (
	"encoding/json"
	"errors"

	"../config"
	"../models"
)

func Login(credenciales models.UsuarioCredenciales) (string, error) {

	var result models.ColaboradorInfoToken

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC Usp_dbAuthUser ?, ?", credenciales.Usuario, credenciales.Password).Scan(&result)

	if result.IdColaborador == 0 {

		return "", errors.New("Credenciales incorrectas")
	}
	json.Marshal(&result)

	token, _ := Create_JWT(result)

	return token, nil

}

func LoginAndGenerateTokenService(credenciales models.UsuarioCredenciales) (string, error) {

	var result models.LoginRespondModel
	var infoToken models.ColaboradorInfoToken

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_AuthWithUserAndPassword ?", credenciales.Usuario).Scan(&result)

	if result.CodigoError != 0 || result.ColaboradorId == 0 {

		return "", errors.New("Error al iniciar sesión. " + result.Mensaje)
	}

	match := CheckPasswordHash(credenciales.Password, result.Clave)

	if !match {
		return "", errors.New("Error al iniciar sesión. Usuario o contraseña incorrecta" + ".")
	}

	db.Raw("EXEC usp_ColaboradorInfoToken  ?", result.ColaboradorId).Scan(&infoToken)

	json.Marshal(&infoToken)

	token, _ := Create_JWT(infoToken)

	return token, nil

}

func LoginService(credenciales models.UsuarioCredenciales) (bool, error) {

	var result models.LoginRespondModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_AuthWithUserAndPassword ?", credenciales.Usuario).Scan(&result)

	match := CheckPasswordHash(credenciales.Password, result.Clave)

	if !match {
		return false, errors.New("Error al iniciar sesión. Usuario o contraseña incorrecta" + ".")
	}

	return true, nil

}

func LoginWithToken(credenciales models.UsuarioCredenciales) (string, error) {

	var result models.ColaboradorInfoToken

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbAuthUserWithToken ?, ?", credenciales.Usuario, credenciales.Token).Scan(&result)

	if result.IdColaborador == 0 {
		return "", errors.New("Credenciales incorrectas")
	}

	json.Marshal(&result)

	token, _ := Create_JWT(result)

	return token, nil

}

func EsAdminBPEvaluaciones(modelo models.ColaboradorInfoToken) bool {

	if modelo.PerfilCod == 1 {
		return true
	}

	return false

}
