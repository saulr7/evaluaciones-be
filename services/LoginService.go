package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"../config"
	"../models"
)

func Login(credenciales models.UsuarioCredenciales) (string, error) {

	var result models.Usuario

	db := config.ConnectDB4DX()
	defer db.Close()

	db.Raw("EXEC Usp_dbAuthUser ?, ?", credenciales.CodigoEmpleado, credenciales.Password).Scan(&result)

	fmt.Println(result)

	if result.IdColaborador == "" {

		return "", errors.New("Credenciales incorrectas")
	}
	json.Marshal(&result)

	token, _ := Create_JWT(result)

	return token, nil

}

func LoginWithToken(credenciales models.UsuarioCredenciales) (string, error) {

	var result models.Usuario

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_dbAuthUserWithToken ?, ?", credenciales.CodigoEmpleado, credenciales.Token).Scan(&result)

	if result.IdColaborador == "" {
		return "", errors.New("Credenciales incorrectas")
	}

	json.Marshal(&result)

	token, _ := Create_JWT(result)

	return token, nil

}
