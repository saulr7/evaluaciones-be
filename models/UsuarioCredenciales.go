package models

type UsuarioCredenciales struct {
	Usuario  string `json:"Usuario"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
