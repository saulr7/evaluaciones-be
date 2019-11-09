package models

import (
	_ "github.com/jinzhu/gorm"
)

type Evaluacion struct {
	IdEvaluacion        int `gorm:"column:idEvaluacion"`
	IdColaborador       int `gorm:"column:idColaborador"`
	Anio                int `gorm:"column:Anio"`
	PorcentajeAvance    int `gorm:"column:PorcentajeAvance"`
	IdGrado             int `gorm:"column:idGrado"`
	EncabezadoPreguntas []EncabezadoPreguntas
}

type EncabezadoPreguntas struct {
	IdGradoPorCompetencia int    `gorm:"column:idGradoPorCompetencia"`
	Detalle               string `gorm:"column:Detalle"`
	GradoPuesto           string `gorm:"column:GradoPuesto"`
	Nivel                 string `gorm:"column:Nivel"`
	Preguntas             []Preguntas
}

type Preguntas struct {
	IdPregunta     int    `gorm:"column:idPregunta"`
	Pregunta       string `gorm:"column:Pregunta"`
	IdRespuesta    int    `gorm:"column:idRespuesta"`
	ValorRespuesta int    `gorm:"column:ValorRespuesta"`
}
