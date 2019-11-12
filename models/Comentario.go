package models

import (
	_ "github.com/jinzhu/gorm"
)

type Comentario struct {
	IdComentario           int    `gorm:"column:idComentario;AUTO_INCREMENT"`
	IdRespuestaPorPregunta int    `gorm:"column:idRespuestaPorPregunta"`
	Comentario             string `gorm:"column:Comentario"`
}

func (Comentario) TableName() string {
	return "ComentariosPorPregunta"
}
