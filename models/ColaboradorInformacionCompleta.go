package models

import (
	_ "github.com/jinzhu/gorm"
)

type ColaboradorInformacionCompleta struct {
	IdColaborador        int    `gorm:"column:ColaboradorId"`
	IdPerfil             int    `gorm:"column:idPerfil"`
	NombreColaborador    string `gorm:"column:Nombre"`
	FechaIngreso         string `gorm:"column:FechaIngresp"`
	IdZona               int    `gorm:"column:idZona"`
	IdComentario         int    `gorm:"column:idComentario"`
	Zona                 string `gorm:"column:Zona"`
	Correo               string `gorm:"column:Correo"`
	Extension            string `gorm:"column:Extension"`
	IdCargo              int    `gorm:"column:idCargo"`
	Cargo                string `gorm:"column:Cargo"`
	IdSubArea            int    `gorm:"column:idSubArea"`
	SubArea              string `gorm:"column:Cargo"`
	IdArea               int    `gorm:"column:idArea"`
	Area                 string `gorm:"column:Area"`
	IdEmpresa            int    `gorm:"column:idEmpresa"`
	Empresa              string `gorm:"column:Empresa"`
	IdCargoJefeImmediato int    `gorm:"column:idCargoJefe"`
	CargoJefe            string `gorm:"column:CargoJefe"`
	NombreJefeImmediato  string `gorm:"column:Jefe"`
}
