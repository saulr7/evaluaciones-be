package services

import (
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

		db.Raw("UPDATE RespuestasPorPregunta SET valor = ? WHERE idRespuestasPorPregunta =  ?", respuesta.IdRespuesta, respuesta.IdRespuestaPorPregunta).Scan(&result)

		if respuesta.TxtComentario != "" {
			comentario.IdComentario = 0
			comentario.Comentario = respuesta.TxtComentario
			comentario.IdRespuestaPorPregunta = respuesta.IdRespuestaPorPregunta

			db.Create(&comentario)
		}
	}

	db.Raw("UPDATE Evaluaciones SET Completo = 1, evaluadoPor =?, FechaCompletado = GETDATE() WHERE idEvaluacion = ?", evaluacionCompletada.EvaluadoPor, evaluacionCompletada.IdEvaluacion).Scan(&result)

	return true, nil
}
