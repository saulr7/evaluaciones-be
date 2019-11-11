package models

import (
	_ "github.com/jinzhu/gorm"
)

type Evaluacion struct {
	IdEvaluacion        int `gorm:"column:idEvaluacion"`
	IdColaborador       int `gorm:"column:idColaborador"`
	Anio                int `gorm:"column:Anio"`
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
	IdPregunta      int    `gorm:"column:idPregunta"`
	Pregunta        string `gorm:"column:Pregunta"`
	IdTipoRespuesta int    `gorm:"column:idTipoRespuesta"`
	Respuestas      []Respuestas
}

type Respuestas struct {
	IdRespuestaPorPregunta int    `gorm:"column:idRespuestaPorPregunta"`
	IdDetallePregunta      int    `gorm:"column:idDetallePregunta"`
	Descripcion            string `gorm:"column:Descripcion"`
	Etiqueta               string `gorm:"column:Etiqueta"`
	Valor                  int    `gorm:"column:Valor"`
	ValorSeteado           int    `gorm:"column:ValorSeteado"`
}
