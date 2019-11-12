package models

import (
	_ "github.com/jinzhu/gorm"
)

type Colaborador struct {
	IdColaborador int    `gorm:"column:idColaborador"`
	Nombre        string `gorm:"column:NombreColaborador"`
}
