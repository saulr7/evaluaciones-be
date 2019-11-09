package models

import (
	_ "github.com/jinzhu/gorm"
)

type Competencias struct {
	IdGradoPorCompetencia int    `gorm:"column:IdGradoPorCompetencia"`
	Detalle               string `gorm:"column:Detalle"`
	GradoPuesto           string `gorm:"column:GradoPuesto"`
	Nivel                 string `gorm:"column:Nivel"`
	Mensaje               string `gorm:"column:Mensaje"`
}
