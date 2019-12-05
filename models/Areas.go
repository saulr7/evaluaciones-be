package models

import (
	_ "github.com/jinzhu/gorm"
)

type Areas struct {
	IdArea          int    `gorm:"column:idArea"`
	Area            string `gorm:"column:Area"`
	IdEmpresa       int    `gorm:"column:idEmpresa"`
	Empresa         string `gorm:"column:Empresa"`
	IdGradoSugerido int    `gorm:"column:idGradoSugerido"`
	IdGrado         int    `gorm:"column:idGrado"`
	GradoPuesto     string `gorm:"column:GradoPuesto"`
	Nivel           string `gorm:"column:Nivel"`
}

type AreaModel struct {
	IdArea int    `gorm:"column:Id"`
	Area   string `gorm:"column:Area"`
}
