package models

import (
	_ "github.com/jinzhu/gorm"
)

type Grados struct {
	IdGrado     int    `gorm:"column:idGrado"`
	GradoPuesto string `gorm:"column:GradoPuesto"`
	Nivel       string `gorm:"column:Nivel"`
}
