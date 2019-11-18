package models

import (
	_ "github.com/jinzhu/gorm"
)

type ColaboradesPendientesEvaluacion struct {
	NombreColaborador string `gorm:"column:NombreColaborador"`
	SubArea           string `gorm:"column:SubArea"`
	IdSubArea         int    `gorm:"column:idSubArea"`
	Evaluaciones
}
