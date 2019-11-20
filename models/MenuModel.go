package models

import (
	_ "github.com/jinzhu/gorm"
)

type MenuModel struct {
	Menu  string `gorm:"column:Menu"`
	Icono string `gorm:"column:Icono"`
	Ruta  string `gorm:"column:Ruta"`
}
