package models

import (
	_ "github.com/jinzhu/gorm"
)

type Usuario struct {
	IdSubArea     int    `gorm:"column:idSubArea"`
	IdColaborador int    `gorm:"column:idColaborador"`
	Nombre        string `gorm:"column:Nombre"`
}
