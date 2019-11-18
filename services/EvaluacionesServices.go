package services

import (
	"fmt"
	"time"

	"../config"
	"../models"
)

func GetEvaluacionPorColaborador(idColaborador string, idEvaluacionAnual string) ([]models.Evaluacion, error) {

	var EvaluacionEnvia []models.Evaluacion

	type Result struct {
		IdEvaluacion      int    `gorm:"column:idEvaluacion"`
		IdEvaluacionAnual int    `gorm:"column:idEvaluacionAnual"`
		IdColaborador     int    `gorm:"column:idColaborador"`
		Anio              int    `gorm:"column:Anio"`
		IdGrado           int    `gorm:"column:idGrado"`
		Completo          bool   `gorm:"column:Completo"`
		Titulo            string `gorm:"column:Titulo"`
		Descripcion       string `gorm:"column:Descripcion"`
	}

	var result []Result
	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEvaluacionPorColaborador ?,?", idColaborador, idEvaluacionAnual).Scan(&result)
	for _, dato := range result {

		var Encabezado []models.EncabezadoPreguntas
		var EvaluacionTemp models.Evaluacion

		EvaluacionTemp.IdEvaluacion = dato.IdEvaluacion
		EvaluacionTemp.IdEvaluacionAnual = dato.IdEvaluacionAnual
		EvaluacionTemp.IdColaborador = dato.IdColaborador
		EvaluacionTemp.Anio = dato.Anio
		EvaluacionTemp.IdGrado = dato.IdGrado
		EvaluacionTemp.Completo = dato.Completo
		EvaluacionTemp.Titulo = dato.Titulo
		EvaluacionTemp.Descripcion = dato.Descripcion

		db.Raw("EXEC usp_GetEncabezadosPorColaborador ?,?", idColaborador, idEvaluacionAnual).Scan(&Encabezado)

		for _, Cabeza := range Encabezado {
			var Preguntas []models.Preguntas

			db.Raw("EXEC usp_GetPreguntasPorColaboradorYGrado ?,?", idColaborador, Cabeza.IdGradoPorCompetencia).Scan(&Preguntas)

			for _, Pregunt := range Preguntas {
				var Respuestas []models.Respuestas

				db.Raw("EXEC usp_GetRespuestasPorPregunta ?,?,?,?", Pregunt.IdTipoRespuesta, Pregunt.IdPregunta, idColaborador, idEvaluacionAnual).Scan(&Respuestas)
				for _, Respu := range Respuestas {
					Pregunt.Respuestas = append(Pregunt.Respuestas, Respu)
				}

				Cabeza.Preguntas = append(Cabeza.Preguntas, Pregunt)
			}
			EvaluacionTemp.EncabezadoPreguntas = append(EvaluacionTemp.EncabezadoPreguntas, Cabeza)
		}
		EvaluacionEnvia = append(EvaluacionEnvia, EvaluacionTemp)
	}
	return EvaluacionEnvia, nil
}

func GetEvaluacionAnual(idColaborador string) ([]models.EvaluacionAnual, error) {

	var result []models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEvaluacionesAnuales ?", idColaborador).Scan(&result)

	return result, nil
}

func GuardarEvaluacionCompletada(evaluacionCompletada models.EvaluacionCompletada) (bool, error) {

	var result []models.EvaluacionAnual
	var comentario models.Comentario

	db := config.ConnectDB()
	defer db.Close()

	for _, respuesta := range evaluacionCompletada.Respuestas {

		respuestaPP := models.RespuestasPorPreguntas{IdEvaluacionAnual: evaluacionCompletada.IdEvaluacionAnual, IdEvaluacion: evaluacionCompletada.EvaluacionId, IdPregunta: respuesta.IdPregunta, IdDetallePregunta: respuesta.IdRespuesta, IdColaborador: evaluacionCompletada.IdColaborador, Valor: respuesta.Valor}

		db.Create(&respuestaPP)

		if respuesta.TxtComentario != "" {
			comentario.IdComentario = 0
			comentario.Comentario = respuesta.TxtComentario
			comentario.IdRespuestaPorPregunta = respuestaPP.IdRespuestaPorPregunta

			db.Create(&comentario)
		}
		respuestaPP.IdRespuestaPorPregunta = 0
	}

	db.Raw("UPDATE Evaluaciones SET Completo = 1, evaluadoPor =?, FechaCompletado = GETDATE() WHERE idEvaluacion = ?", evaluacionCompletada.EvaluadoPor, evaluacionCompletada.EvaluacionId).Scan(&result)

	return true, nil
}

func NuevaEvaluacionPorMeta(evaluacionPorMeta models.NuevaEvaluacionPorMeta) (models.EvaluacionAnual, error) {

	var result models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	// tx := db.Begin()

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()

	// 	}
	// }()

	db.Raw(" SELECT * FROM EvaluacionesAnuales WHERE idEvaluacionAnual = ?", evaluacionPorMeta.IdEvaluacionAnual).Scan(&result)

	evaluacionPorMeta.Desde = result.Desde
	evaluacionPorMeta.Hasta = result.Hasta
	evaluacionPorMeta.Estado = true
	evaluacionPorMeta.IdPadre = result.IdEvaluacionAnual
	evaluacionPorMeta.FechaCreacion = time.Now()
	evaluacionPorMeta.FechaModificacion = time.Now()
	evaluacionPorMeta.IdEvaluacionAnual = 0
	evaluacionPorMeta.TodasLasAreas = false

	db.Create(&evaluacionPorMeta)

	var colaboradores []models.Usuario

	db.Raw(" EXEC  usp_GetEquipoPorLider ?", evaluacionPorMeta.CreadaPor).Scan(&colaboradores)

	// var PreguntaModel models.PreguntasPorMetaModel
	var PreguntasPorEvaluacionPorMeta models.PreguntasPorEvaluacionPorMeta

	for _, pregunta := range evaluacionPorMeta.Preguntas {
		var PreguntaModel models.PreguntasPorMetaModel
		PreguntaModel.IdPregunta = 0
		PreguntaModel.Peso = pregunta.Meta
		PreguntaModel.Pregunta = pregunta.Pregunta
		PreguntaModel.IdTipoRespuesta = pregunta.IdTipoRespuesta

		db.NewRecord(&PreguntaModel)
		db.Create(&PreguntaModel)

		fmt.Println(PreguntaModel.IdPregunta)

		PreguntasPorEvaluacionPorMeta.IdPreguntasPorEvaluacionPorMeta = 0
		PreguntasPorEvaluacionPorMeta.IdPregunta = PreguntaModel.IdPregunta
		PreguntasPorEvaluacionPorMeta.IdEvaluacionAnual = evaluacionPorMeta.IdEvaluacionAnual
		PreguntasPorEvaluacionPorMeta.Meta = pregunta.Meta

		db.Create(&PreguntasPorEvaluacionPorMeta)
		PreguntaModel.IdPregunta = 0
	}

	var PreguntasGuardadas []models.PreguntasPorEvaluacionPorMeta

	db.Where("idEvaluacionAnual = ?", evaluacionPorMeta.IdEvaluacionAnual).Find(&PreguntasGuardadas)

	for _, colaborador := range colaboradores {
		var nuevaEvaluacion models.Evaluaciones

		if colaborador.Accion != "R" {
			nuevaEvaluacion.IdEvaluacionAnual = evaluacionPorMeta.IdEvaluacionAnual
			nuevaEvaluacion.IdColaborador = colaborador.IdColaborador
			nuevaEvaluacion.Anio = time.Now().Year()
			nuevaEvaluacion.IdGrado = 0
			nuevaEvaluacion.Completo = false
			nuevaEvaluacion.Estado = true
			db.Create(&nuevaEvaluacion)
		}
	}

	// return result, tx.Commit().Error
	return result, nil
}

func GetEvaluacionMetaPorColaborador(idColaborador string, idPadre string) (models.EvaluacionMeta, error) {

	db := config.ConnectDB()
	defer db.Close()

	var EvaluacionTemp models.EvaluacionMeta
	var preguntas []models.PreguntasPorMetaContestadas

	db.Raw("EXEC usp_GetEvaluacionesPorMetaPorColaborador ?,?", idColaborador, idPadre).Scan(&EvaluacionTemp)

	db.Raw("EXEC usp_GetPreguntasEvaluacionMeta ?, ?, ?", EvaluacionTemp.IdEvaluacionAnual, idColaborador, EvaluacionTemp.IdEvaluacion).Scan(&preguntas)

	EvaluacionTemp.PreguntasMeta = preguntas

	return EvaluacionTemp, nil
}

func GetEvaluacionsTodas() ([]models.EvaluacionAnual, error) {

	var evaluaciones []models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	db.Order("FechaCreacion desc").Where("idPadre = 0").Find(&evaluaciones)

	return evaluaciones, nil
}
