package models

import (
	_ "github.com/jinzhu/gorm"
)

type ColaboradorAdmin struct {
	IdColaboradorPerfil string `gorm:"column:IdColaboradorPerfil"`
	IdPerfil            string `gorm:"column:IdPerfil"`
	IdColaborador       string `gorm:"column:IdColaborador"`
	ModificadoFecha     string `gorm:"column:ModificadoFecha"`
	ModificadoPor       string `gorm:"column:ModificadoPor"`
	Cargo               string `gorm:"column:Cargo"`
	NombreColaborador   string `gorm:"column:NombreColaborador"`
	Perfil              string `gorm:"column:Perfil"`
	ModificadoPorNombre string `gorm:"column:ModificadoPorNombre"`
}
