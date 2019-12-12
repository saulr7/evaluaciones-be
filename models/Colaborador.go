package models

import (
	_ "github.com/jinzhu/gorm"
)

type Colaborador struct {
	IdColaborador int    `gorm:"column:idColaborador"`
	Nombre        string `gorm:"column:NombreColaborador"`
}

type ColaboradorInfo struct {
	IdColaborador int    `gorm:"column:ColaboradorId"`
	Nombre        string `gorm:"column:Nombre"`
	FechaIngreso  string `gorm:"column:FechaIngreso"`
	Cargo         string `gorm:"column:Cargo"`
	Area          string `gorm:"column:Area"`
	Activo        bool   `gorm:"column:Activo"`
}

type NuevoColaboradorModel struct {
	IdColaborador int
	Nombre        string
	FechaIngreso  string
	AgregadoPor   int
	IdCargo       int
}
type ColaboradorInfoToken struct {
	IdColaborador int    `gorm:"column:ColaboradorId"`
	Nombre        string `gorm:"column:Nombre"`
	CargoId       int    `gorm:"column:CargoId"`
	Cargo         string `gorm:"column:Cargo"`
	AreaId        int    `gorm:"column:AreaId"`
	Area          string `gorm:"column:Area"`
	EmpresaId     int    `gorm:"column:EmpresaId"`
	Empresa       string `gorm:"column:Empresa"`
	PerfilCod     int    `gorm:"column:PerfilCod"`
	Usuario       string `gorm:"column:Usuario"`
	CambiarClave  bool   `gorm:"column:CambiarClave"`
}

type ColaboradorCargo struct {
	IdColaborador int    `gorm:"column:ColaboradorId"`
	Nombre        string `gorm:"column:Nombre"`
	CargoId       int    `gorm:"column:CargoId"`
}
