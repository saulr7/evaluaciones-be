package models

import (
	_ "github.com/jinzhu/gorm"
)

type Usuario struct {
	IdSubArea     int    `gorm:"column:idSubArea"`
	IdColaborador string `gorm:"column:Empleado"`
	Nombre        string `gorm:"column:UsrNom"`
	Accion        string `gorm:"column:Accion"`
	PerfilCod     int    `gorm:"column:PerfilCod"`
}

type NuevoUsuarioModel struct {
	Usuario       int
	Clave         string
	ColaboradorId int
	AgreadoPor    int
	CambiarClave  bool
}

type UsuarioPorcentaje struct {
	Usuario
	Completo         bool `gorm:"column:Completo"`
	AceptoEvaluacion bool `gorm:"column:AceptoEvaluacion"`
}

type UsuarioCajeros struct {
	Usuario
	Cargo string `gorm:"column:Cargo"`
}

type CambiarContrasenaModel struct {
	ColaboradorId     int
	ClaveActual       string
	ClaveNueva        string
	ClaveConfirmacion string
	Usuario           string
}

type ResetearContrasenaModel struct {
	ColaboradorId int
	ClaveNueva    string
	Usuario       string
	ModificadoPor string
}

type UsuarioInfoModel struct {
	Usuario           string `gorm:"column:Usuario"`
	ColaboradorId     int    `gorm:"column:ColaboradorId"`
	UsuarioActivo     bool   `gorm:"column:UsuarioActivo"`
	CambiarClave      bool   `gorm:"column:CambiarClave"`
	Nombre            string `gorm:"column:Nombre"`
	ColaboradorActivo bool   `gorm:"column:ColaboradorActivo"`
	PerfilId          int    `gorm:"column:PerfilId"`
	Perfil            string `gorm:"column:Perfil"`
}
