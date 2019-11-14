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

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()

		}
	}()

	for _, respuesta := range evaluacionCompletada.Respuestas {

		query := db.Raw("UPDATE RespuestasPorPregunta SET valor = ? WHERE idRespuestasPorPregunta =  ?", respuesta.IdRespuesta, respuesta.IdRespuestaPorPregunta).Scan(&result)

		fmt.Println(query.RowsAffected, query.Error)

		if query.Error != nil {
			tx.Rollback()
			panic(query.Error)

		}

		if respuesta.TxtComentario != "" {
			comentario.IdComentario = 0
			comentario.Comentario = respuesta.TxtComentario
			comentario.IdRespuestaPorPregunta = respuesta.IdRespuestaPorPregunta

			db.Create(&comentario)
		}
	}

	db.Raw("UPDATE Evaluaciones SET Completo = 1, evaluadoPor =?, FechaCompletado = GETDATE() WHERE idEvaluacion = ?", evaluacionCompletada.EvaluadoPor, evaluacionCompletada.IdEvaluacion).Scan(&result)

	return true, tx.Commit().Error
}

func NuevaEvaluacionPorMeta(evaluacionPorMeta models.NuevaEvaluacionPorMeta) (models.EvaluacionAnual, error) {

	var result models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()

		}
	}()

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

	fmt.Println(evaluacionPorMeta.IdEvaluacionAnual)

	var colaboradores []models.Usuario

	db.Raw(" EXEC  usp_GetEquipoPorLider ?", evaluacionPorMeta.CreadaPor).Scan(&colaboradores)

	var PreguntaModel models.PreguntasPorMetaModel
	var PreguntasPorEvaluacionPorMeta models.PreguntasPorEvaluacionPorMeta

	for _, pregunta := range evaluacionPorMeta.Preguntas {
		PreguntaModel.IdPregunta = 0
		PreguntaModel.Pregunta = pregunta.Pregunta
		PreguntaModel.IdTipoRespuesta = pregunta.IdTipoRespuesta

		db.Create(&PreguntaModel)

		PreguntasPorEvaluacionPorMeta.IdPreguntasPorEvaluacionPorMeta = 0
		PreguntasPorEvaluacionPorMeta.IdPregunta = PreguntaModel.IdPregunta
		PreguntasPorEvaluacionPorMeta.IdEvaluacionAnual = evaluacionPorMeta.IdEvaluacionAnual
		PreguntasPorEvaluacionPorMeta.Meta = pregunta.Meta

		db.Create(&PreguntasPorEvaluacionPorMeta)
	}

	var PreguntasGuardadas []models.PreguntasPorEvaluacionPorMeta

	db.Where("idEvaluacionAnual = ?", evaluacionPorMeta.IdEvaluacionAnual).Find(&PreguntasGuardadas)

	for _, preguntaG := range PreguntasGuardadas {
		for _, colaborador := range colaboradores {
			if colaborador.Accion != "R" {

				query := db.Raw("INSERT INTO RespuestasPorPregunta (idEvaluacionAnual, idEvaluacion, idPregunta, idDetallePregunta,idColaborador,Valor ) values (?,?,?,?,?,?)", evaluacionPorMeta.IdEvaluacionAnual, evaluacionPorMeta.IdEvaluacionAnual, preguntaG.IdPregunta, 0, colaborador.IdColaborador, 0).Scan(&result)
				// db.Raw("INSERT INTO RespuestasPorPregunta (idEvaluacionAnual, idEvaluacion, idPregunta, idDetallePregunta,idColaborador,Meta,Valor ) values (?,?, ?,?,?,?,?)", evaluacionPorMeta.IdEvaluacionAnual, evaluacionPorMeta.IdEvaluacionAnual, preguntaG.IdPregunta, 0, colaborador.IdColaborador, 16, 0).Scan(&result)
				fmt.Println(query.Error)

				// if query.Error != nil {
				// 	//tx.Rollback()
				// 	panic(query.Error)

				// }
			}

		}
	}

	return result, tx.Commit().Error
	// return result, nil
}

func GetEvaluacionMetaPorColaborador(idColaborador string, idPadre string) ([]models.EvaluacionMeta, error) {

	var EvaluacionEnvia []models.EvaluacionMeta

	type Result struct {
		IdEvaluacionAnual int       `gorm:"column:idEvaluacionAnual"`
		IdPadre           int       `gorm:"column:idPadre"`
		Descripcion       string    `gorm:"column:Descripcion"`
		FechaCreacion     time.Time `gorm:"column:FechaCreacion"`
		CreadaPor         int       `gorm:"column:CreadaPor"`
		NombreCreador     string    `gorm:"column:NombreCreador"`
		TodoElEquipo      bool      `gorm:"column:TodoElEquipo"`
		IdSubArea         int       `gorm:"column:idSubArea"`
		TotalPreguntas    int       `gorm:"column:TotalPreguntas"`
	}

	var result []Result
	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEvaluacionesPorMetaPorColaborador ?,?", idColaborador, idPadre).Scan(&result)
	for _, dato := range result {

		var Encabezado []models.PreguntasMeta
		var EvaluacionTemp models.EvaluacionMeta

		EvaluacionTemp.IdEvaluacionAnual = dato.IdEvaluacionAnual
		EvaluacionTemp.IdPadre = dato.IdPadre
		EvaluacionTemp.Descripcion = dato.Descripcion
		EvaluacionTemp.FechaCreacion = dato.FechaCreacion
		EvaluacionTemp.CreadaPor = dato.CreadaPor
		EvaluacionTemp.NombreCreador = dato.NombreCreador
		EvaluacionTemp.TodoElEquipo = dato.TodoElEquipo
		EvaluacionTemp.IdSubArea = dato.IdSubArea
		EvaluacionTemp.TotalPreguntas = dato.TotalPreguntas

		db.Raw("EXEC usp_GetPreguntasPorEvaluacionPorMeta ?,?", dato.IdEvaluacionAnual, idColaborador).Scan(&Encabezado)

		for _, Cabeza := range Encabezado {
			//var Preguntas []models.Preguntas

			// db.Raw("EXEC usp_GetPreguntasPorColaboradorYGrado ?,?", idColaborador, Cabeza.IdGradoPorCompetencia).Scan(&Preguntas)

			// for _, Pregunt := range Preguntas {
			// 	var Respuestas []models.Respuestas

			// 	db.Raw("EXEC usp_GetRespuestasPorPregunta ?,?,?,?", Pregunt.IdTipoRespuesta, Pregunt.IdPregunta, idColaborador, idEvaluacionAnual).Scan(&Respuestas)
			// 	for _, Respu := range Respuestas {
			// 		Pregunt.Respuestas = append(Pregunt.Respuestas, Respu)
			// 	}

			// 	Cabeza.Preguntas = append(Cabeza.Preguntas, Pregunt)
			// }
			EvaluacionTemp.PreguntasMeta = append(EvaluacionTemp.PreguntasMeta, Cabeza)
		}
		EvaluacionEnvia = append(EvaluacionEnvia, EvaluacionTemp)
	}
	return EvaluacionEnvia, nil
}
