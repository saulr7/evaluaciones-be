package models

import (
	"time"

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

type EvaluacionAnual struct {
	IdEvaluacionAnual int       `gorm:"column:idEvaluacionAnual"`
	Titulo            string    `gorm:"column:Titulo"`
	Descripcion       string    `gorm:"column:Descripcion"`
	Desde             time.Time `gorm:"column:desde"`
	Hasta             time.Time `gorm:"column:Hasta"`
	CreadaPor         int       `gorm:"column:CreadaPor"`
	NombreCreador     string    `gorm:"column:NombreCreador"`
	FechaCreacion     time.Time `gorm:"column:FechaCreacion"`
	FechaModificacion time.Time `gorm:"column:FechaModificacion"`
	ModificadaPor     int       `gorm:"column:ModificadaPor"`
	NombreModifico    string    `gorm:"column:NombreModifico"`
	IdArea            int       `gorm:"column:idArea"`
	IdGrado           int       `gorm:"column:idGrado"`
}
