package models

type UsuarioCredenciales struct {
	CodigoEmpleado string `json:"codigoEmpleado"`
	Token          string `json:"token"`
	Password       string `json:"password"`
}
