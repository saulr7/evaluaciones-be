package models

import (
	_ "github.com/jinzhu/gorm"
)

type EvaluacionCompletada struct {
	IdColaborador     int
	IdEvaluacionAnual int
	EvaluacionId      int
	EvaluadoPor       int
	Respuestas        []RespuestasCompletadas
}

type RespuestasCompletadas struct {
	IdPregunta             int
	TxtComentario          string
	IdRespuestaPorPregunta int
	IdRespuesta            int
	Valor                  int
}
