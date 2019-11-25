package models

import (
	_ "github.com/jinzhu/gorm"
)

type ReporteCompetencias struct {
	Competencia              string  `gorm:"column:Competencia"`
	ComportamientosEsperados string  `gorm:"column:ComportamientosEsperados"`
	CalificacionObtenida     float32 `gorm:"column:CalificacionObtenida"`
	Brecha                   float32 `gorm:"column:Brecha"`
	Promedio_Competencia     float32 `gorm:"column:Promedio_Competencia"`
	VariacionAnioAnterior    float32 `gorm:"column:VariacionAnioAnterior"`
}
