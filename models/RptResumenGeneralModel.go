package models

import (
	_ "github.com/jinzhu/gorm"
)

type RptResumenGeneralModel struct {
	IdColaborador       int     `gorm:"column:idColaborador"`
	NombreColaborador   string  `gorm:"column:NombreColaborador"`
	IdCargo             int     `gorm:"column:idCargo"`
	Cargo               string  `gorm:"column:Cargo"`
	IdArea              int     `gorm:"column:idArea"`
	Area                string  `gorm:"column:Area"`
	NombreJefeImmediato string  `gorm:"column:NombreJefeImmediato"`
	Valor60             float64 `gorm:"column:Valor60"`
	Valor40             float64 `gorm:"column:Valor40"`
}
