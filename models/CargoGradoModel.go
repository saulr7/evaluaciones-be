package models

import (
	_ "github.com/jinzhu/gorm"
)

type CargoGradoModel struct {
	IdCargo int    `gorm:"column:idCargo"`
	Cargo   string `gorm:"column:Cargo"`
	Grados
}

type CargoGradoActualizarModel struct {
	IdGrado int
	Cargos  []int
}
