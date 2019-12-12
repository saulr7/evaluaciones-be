package services

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"../config"
	"../models"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func CreateUsuarioService(model models.NuevoUsuarioModel) (string, error) {

	var result []models.Usuario

	db := config.ConnectDBEO()
	defer db.Close()

	var Clave = String(12)

	model.Clave, _ = HashPassword(Clave)

	db.Raw("EXEC usp_CreateUsuario ?,?,?,?,?", model.Usuario, model.Clave, model.ColaboradorId, model.AgreadoPor, model.CambiarClave).Scan(&result)

	return Clave, nil

}

func GetUsuariosService() ([]models.UsuarioInfoModel, error) {

	var result []models.UsuarioInfoModel

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_GetUsuarios").Scan(&result)

	return result, nil
}

func UpdateUsuarioActivar(usuario string, Activar bool, ModificadoPor int) (models.ColaboradorInfo, error) {

	var result models.ColaboradorInfo

	db := config.ConnectDBEO()
	defer db.Close()

	db.Raw("EXEC usp_ToggleActivarUsuario ?, ?, ?", usuario, Activar, ModificadoPor).Scan(&result)

	return result, nil
}

func CambiarContrasenaService(model models.CambiarContrasenaModel) ([]models.Usuario, error) {

	var result []models.Usuario

	db := config.ConnectDBEO()
	defer db.Close()

	var credential models.UsuarioCredenciales

	credential.Usuario = model.Usuario
	credential.Password = model.ClaveActual

	validated, err := LoginService(credential)

	if !validated || err != nil {
		return nil, errors.New("Credenciales no válidas")
	}

	if model.ClaveNueva == "" || model.ClaveNueva != model.ClaveConfirmacion {
		return nil, errors.New("La nueva contraseña no es válida")
	}

	model.ClaveNueva, _ = HashPassword(model.ClaveNueva)

	db.Raw("EXEC usp_UsuarioCambiarContrasena ?,?", model.ColaboradorId, model.ClaveNueva).Scan(&result)

	return result, nil
}

func ResetearContrasenaService(modelo models.ResetearContrasenaModel) (string, error) {

	var Clave = String(12)
	var result []models.Usuario

	db := config.ConnectDBEO()
	defer db.Close()

	modelo.ClaveNueva, _ = HashPassword(Clave)

	db.Raw("EXEC usp_ResetearContrasenaUsuario ?,?, ?", modelo.Usuario, modelo.ClaveNueva, modelo.ModificadoPor).Scan(&result)
	return Clave, nil
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		valor := charset[seededRand.Intn(len(charset))]
		if i == 0 {
			if _, err := strconv.Atoi(string(valor)); err == nil {
				valor = byte('X')
			}

		}
		b[i] = valor
	}
	return string(b)
}
