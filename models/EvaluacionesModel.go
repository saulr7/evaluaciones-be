package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Evaluacion struct {
	IdEvaluacion        int       `gorm:"column:idEvaluacion"`
	IdEvaluacionAnual   int       `gorm:"column:idEvaluacionAnual"`
	IdColaborador       int       `gorm:"column:idColaborador"`
	Anio                int       `gorm:"column:Anio"`
	IdGrado             int       `gorm:"column:idGrado"`
	Completo            bool      `gorm:"column:Completo"`
	Titulo              string    `gorm:"column:Titulo"`
	Descripcion         string    `gorm:"column:Descripcion"`
	AceptoEvaluacion    bool      `gorm:"column:AceptoEvaluacion"`
	PermiteGuardar      bool      `gorm:"column:PermiteGuardar"`
	FechaAcepto         time.Time `gorm:"column:FechaAcepto"`
	EncabezadoPreguntas []EncabezadoPreguntas
}

type EncabezadoPreguntas struct {
	IdGradoPorCompetencia int    `gorm:"column:idGradoPorCompetencia"`
	Descripcion           string `gorm:"column:Descripcion"`
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
	IdRespuestaPorPregunta int    `gorm:"column:idRespuestasPorPregunta"`
	IdDetallePregunta      int    `gorm:"column:idDetallePregunta"`
	Descripcion            string `gorm:"column:Descripcion"`
	Etiqueta               string `gorm:"column:Etiqueta"`
	Valor                  int    `gorm:"column:Valor"`
	ValorSeteado           int    `gorm:"column:ValorSeteado"`
	IdComentario           int    `gorm:"column:idComentario"`
	Comentario             string `gorm:"column:Comentario"`
}

type EvaluacionAnual struct {
	IdEvaluacionAnual int       `gorm:"column:idEvaluacionAnual;AUTO_INCREMENT"`
	Titulo            string    `gorm:"column:Titulo"`
	Descripcion       string    `gorm:"column:Descripcion"`
	Desde             time.Time `gorm:"column:Desde"`
	Hasta             time.Time `gorm:"column:Hasta"`
	CreadaPor         int       `gorm:"column:CreadaPor"`
	FechaCreacion     time.Time `gorm:"column:FechaCreacion"`
	FechaModificacion time.Time `gorm:"column:FechaModificacion"`
	ModificadaPor     int       `gorm:"column:ModificadaPor"`
	IdCargoPadre      int       `gorm:"column:idCargoPadre"`
	TodasLasAreas     bool      `gorm:"column:TodasLasAreas"`
	IdPadre           int       `gorm:"column:idPadre"`
	TodoElEquipo      bool      `gorm:"column:TodoElEquipo"`
	Estado            bool      `gorm:"column:Estado"`
}

func (EvaluacionAnual) TableName() string {
	return "EvaluacionesAnuales"
}

type EvaluacionMeta struct {
	IdEvaluacionAnual int       `gorm:"column:idEvaluacionAnual"`
	IdEvaluacion      int       `gorm:"column:idEvaluacion"`
	IdPadre           int       `gorm:"column:idPadre"`
	Descripcion       string    `gorm:"column:Descripcion"`
	Titulo            string    `gorm:"column:Titulo"`
	FechaCreacion     time.Time `gorm:"column:FechaCreacion"`
	CreadaPor         int       `gorm:"column:CreadaPor"`
	NombreCreador     string    `gorm:"column:NombreCreador"`
	TodoElEquipo      bool      `gorm:"column:TodoElEquipo"`
	IdCargoPadre      int       `gorm:"column:idCargoPadre"`
	TotalPreguntas    int       `gorm:"column:TotalPreguntas"`
	Completo          bool      `gorm:"column:Completo"`
	PreguntasMeta     []PreguntasPorMetaContestadas
}

type PreguntasMeta struct {
	IdPreguntasPorEvaluacionPorMeta int    `gorm:"column:idPreguntasPorEvaluacionPorMeta"`
	IdRespuestasPorPregunta         int    `gorm:"column:idRespuestasPorPregunta"`
	IdPregunta                      int    `gorm:"column:idPregunta"`
	Pregunta                        string `gorm:"column:Pregunta"`
	Meta                            int    `gorm:"column:Meta"`
	Valor                           int    `gorm:"column:Valor"`
}

type EvaluacionAnualCompleta struct {
	EvaluacionAnual
	NombreEvaluo    string    `gorm:"column:NombreEvaluo"`
	FechaCompletado time.Time `gorm:"column:FechaCompletado"`
}
