package models

import (
	_ "github.com/jinzhu/gorm"
)

type CargoModel struct {
	IdCargo int    `gorm:"column:Id"`
	Cargo   string `gorm:"column:Cargo"`
}

type NuevoCargoModel struct {
	Cargo        string
	CargoPadreId int
	AgregadoPor  int
	AreaId       int
}

type CargoPadreYEmpresaModel struct {
	IdCargo      int    `gorm:"column:Id"`
	Cargo        string `gorm:"column:Cargo"`
	CargoPadreId int    `gorm:"column:CargoPadreId"`
	EmpresaId    int    `gorm:"column:EmpresaId"`
	IdCargoPadre int    `gorm:"column:IdCargoPadre"`
	Area         string `gorm:"column:Area"`
	CargoPadre   string `gorm:"column:CargoPadre"`
	Empresa      string `gorm:"column:Empresa"`
	Activo       bool   `gorm:"column:Activo"`
}
