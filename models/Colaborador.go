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
	AgregadoPor   string
	IdCargo       int
}
