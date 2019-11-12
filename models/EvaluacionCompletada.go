package models

import (
	_ "github.com/jinzhu/gorm"
)

type EvaluacionCompletada struct {
	IdColaborador int
	IdEvaluacion  int
	EvaluadoPor   int
	Respuestas    []RespuestasCompletadas
}

type RespuestasCompletadas struct {
	IdPregunta             int
	TxtComentario          string
	IdRespuestaPorPregunta int
	IdRespuesta            int
}
