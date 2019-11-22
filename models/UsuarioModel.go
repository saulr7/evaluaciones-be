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

type UsuarioPorcentaje struct {
	Usuario
	Completo         bool `gorm:"column:Completo"`
	AceptoEvaluacion bool `gorm:"column:AceptoEvaluacion"`
}

type UsuarioCajeros struct {
	Usuario
	Cargo string `gorm:"column:Cargo"`
}
