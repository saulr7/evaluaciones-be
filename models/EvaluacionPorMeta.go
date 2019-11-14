package models

import (
	_ "github.com/jinzhu/gorm"
)

type NuevaEvaluacionPorMeta struct {
	EvaluacionAnual
	Preguntas []PreguntasPorMeta
}

type PreguntasPorMeta struct {
	Meta int
	PreguntasPorMetaModel
}

type PreguntasPorMetaModel struct {
	IdPregunta      int    `gorm:"column:idPregunta;AUTO_INCREMENT"`
	Pregunta        string `gorm:"column:Pregunta"`
	IdTipoRespuesta int    `gorm:"column:idTipoRespuesta"`
}

type PreguntasPorEvaluacionPorMeta struct {
	IdPreguntasPorEvaluacionPorMeta int `gorm:"column:idPreguntasPorEvaluacionPorMeta;AUTO_INCREMENT"`
	IdPregunta                      int `gorm:"column:idPregunta;"`
	IdEvaluacionAnual               int `gorm:"column:idEvaluacionAnual"`
	Meta                            int `gorm:"column:Meta"`
}

func (PreguntasPorEvaluacionPorMeta) TableName() string {
	return "PreguntasPorEvaluacionPorMeta "
}

func (PreguntasPorMetaModel) TableName() string {
	return "preguntas "
}

type Evaluaciones struct {
	IdEvaluacion      int    `gorm:"column:idEvaluacion;AUTO_INCREMENT"`
	IdEvaluacionAnual int    `gorm:"column:idEvaluacionAnual"`
	IdColaborador     string `gorm:"column:idColaborador"`
	Anio              int    `gorm:"column:Anio"`
	IdGrado           int    `gorm:"column:idGrado"`
	Completo          bool   `gorm:"column:Completo"`
	Estado            int    `gorm:"column:Estado"`
}

func (Evaluaciones) TableName() string {
	return "Evaluaciones"
}
