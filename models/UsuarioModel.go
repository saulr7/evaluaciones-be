package models

import (
	_ "github.com/jinzhu/gorm"
)

type Usuario struct {
	IdSubArea     int    `gorm:"column:idSubArea"`
	IdColaborador string `gorm:"column:idColaborador"`
	Nombre        string `gorm:"column:Nombre"`
}
