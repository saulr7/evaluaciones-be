package models

import (
	"time"

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

type ReporteReseteoDeNota struct {
	RowNum        int       `gorm:"column:RowNum"`
	ColaboradorId int       `gorm:"column:ColaboradorId"`
	Nombre        string    `gorm:"column:Nombre"`
	Nota          string    `gorm:"column:Nota"`
	ReseteadoPor  string    `gorm:"column:ReseteadoPor"`
	FechaAgregada time.Time `gorm:"column:FechaAgregada"`
}
