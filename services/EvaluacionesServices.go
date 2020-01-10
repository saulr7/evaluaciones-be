package services

import (
	"errors"
	"fmt"
	"time"

	"../config"
	"../models"
)

func GetEvaluacionPorColaborador(idColaborador string, idEvaluacionAnual string) ([]models.Evaluacion, error) {

	var EvaluacionEnvia []models.Evaluacion

	type Result struct {
		IdEvaluacion      int       `gorm:"column:idEvaluacion"`
		IdEvaluacionAnual int       `gorm:"column:idEvaluacionAnual"`
		IdColaborador     int       `gorm:"column:idColaborador"`
		Anio              int       `gorm:"column:Anio"`
		IdGrado           int       `gorm:"column:idGrado"`
		Completo          bool      `gorm:"column:Completo"`
		Titulo            string    `gorm:"column:Titulo"`
		Descripcion       string    `gorm:"column:Descripcion"`
		AceptoEvaluacion  bool      `gorm:"column:AceptoEvaluacion"`
		FechaAcepto       time.Time `gorm:"column:FechaAcepto"`
		PermiteGuardar    bool      `gorm:"column:PermiteGuardar"`
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
		EvaluacionTemp.AceptoEvaluacion = dato.AceptoEvaluacion
		EvaluacionTemp.FechaAcepto = dato.FechaAcepto
		EvaluacionTemp.PermiteGuardar = dato.PermiteGuardar

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

func GetEvaluacionesCompletas(idColaborador string) ([]models.EvaluacionAnualCompleta, error) {

	var result []models.EvaluacionAnualCompleta

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEvaluacionesAplicadasPorColaborador ?", idColaborador).Scan(&result)

	return result, nil
}

func GuardarEvaluacionCompletada(evaluacionCompletada models.EvaluacionCompletada) (bool, error) {

	var result []models.EvaluacionAnual
	//var comentario models.Comentario

	db := config.ConnectDB()
	defer db.Close()

	LogToFile("Guardar evaluacion completa", evaluacionCompletada)

	for _, respuesta := range evaluacionCompletada.Respuestas {

		// respuestaPP := models.RespuestasPorPreguntas{IdRespuestaPorPregunta: 0, IdEvaluacionAnual: evaluacionCompletada.IdEvaluacionAnual, IdEvaluacion: evaluacionCompletada.EvaluacionId, IdPregunta: respuesta.IdPregunta, IdDetallePregunta: respuesta.IdRespuesta, IdColaborador: evaluacionCompletada.IdColaborador, Valor: respuesta.Valor}

		// db.Create(respuestaPP)

		// if respuesta.TxtComentario != "" {
		// 	comentario.IdComentario = 0
		// 	comentario.Comentario = respuesta.TxtComentario
		// 	comentario.IdRespuestaPorPregunta = respuestaPP.IdRespuestaPorPregunta

		// 	db.Create(&comentario)
		// }

		db.Raw(" exec usp_GuardarRespuesta ?, ?, ?, ?, ?, ?, ?", evaluacionCompletada.IdEvaluacionAnual, evaluacionCompletada.EvaluacionId, respuesta.IdPregunta, respuesta.IdRespuesta, evaluacionCompletada.IdColaborador, respuesta.Valor, respuesta.TxtComentario).Scan(&result)

		// respuestaPP.IdRespuestaPorPregunta = 0
	}

	db.Raw("UPDATE Evaluaciones SET Completo = 1, evaluadoPor =?, FechaCompletado = GETDATE() WHERE idEvaluacion = ?", evaluacionCompletada.EvaluadoPor, evaluacionCompletada.EvaluacionId).Scan(&result)

	return true, nil
}

func NuevaEvaluacionPorMeta(evaluacionPorMeta models.NuevaEvaluacionPorMeta) (models.EvaluacionAnual, error) {

	var result models.EvaluacionAnual
	var resultEvaluacionIndividuales []models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	LogToFile("Nueva evaluacion por meta", evaluacionPorMeta)

	db.Raw(" SELECT * FROM EvaluacionesAnuales WHERE idPadre = ? AND idCargoPadre  = ?", evaluacionPorMeta.IdPadre, evaluacionPorMeta.IdCargoPadre).Scan(&result)

	if evaluacionPorMeta.TodoElEquipo {

		if result.IdEvaluacionAnual != 0 {
			return result, errors.New("Ya se ha registrado una evaluación para todo el equipo de esta área ")
		}

	} else {

		var colaboradoresList []string

		for _, colaborad := range evaluacionPorMeta.Colaboradores {
			colaboradoresList = append(colaboradoresList, colaborad.IdColaborador)
		}

		db.Raw("select * from Evaluaciones where  idEvaluacionAnual = ? and idColaborador in (?)", result.IdEvaluacionAnual, colaboradoresList).Scan(&resultEvaluacionIndividuales)

		if len(resultEvaluacionIndividuales) > 0 {

			return result, errors.New("Hay colaboradores que ya tienen una subevaluación para esta evaluación ")
		}
	}

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

	var PreguntasPorEvaluacionPorMeta models.PreguntasPorEvaluacionPorMeta

	for _, pregunta := range evaluacionPorMeta.Preguntas {
		var PreguntaModel models.PreguntasPorMetaModel
		PreguntaModel.IdPregunta = 0
		PreguntaModel.Peso = pregunta.Meta
		PreguntaModel.Pregunta = pregunta.Pregunta
		PreguntaModel.IdTipoRespuesta = pregunta.IdTipoRespuesta

		// db.NewRecord(&PreguntaModel)
		db.Create(&PreguntaModel)

		PreguntasPorEvaluacionPorMeta.IdPreguntasPorEvaluacionPorMeta = 0
		PreguntasPorEvaluacionPorMeta.IdPregunta = PreguntaModel.IdPregunta
		PreguntasPorEvaluacionPorMeta.IdEvaluacionAnual = evaluacionPorMeta.IdEvaluacionAnual
		PreguntasPorEvaluacionPorMeta.Meta = pregunta.Meta

		db.Create(&PreguntasPorEvaluacionPorMeta)
		PreguntaModel.IdPregunta = 0
	}

	var PreguntasGuardadas []models.PreguntasPorEvaluacionPorMeta

	db.Where("idEvaluacionAnual = ?", evaluacionPorMeta.IdEvaluacionAnual).Find(&PreguntasGuardadas)

	var colaboradores []models.Usuario

	if evaluacionPorMeta.TodoElEquipo {
		db.Raw(" EXEC  usp_GetEquipoPorLider ?", evaluacionPorMeta.CreadaPor).Scan(&colaboradores)
	} else {
		colaboradores = evaluacionPorMeta.Colaboradores
	}

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

			if evaluacionPorMeta.TodoElEquipo == false {
				fmt.Println("Entro")
				var evaluacionMetaColaborador models.EvaluacionPorMetaPorColaborador
				evaluacionMetaColaborador.IdEvaluacionPorMetaPorColaborador = 0
				evaluacionMetaColaborador.IdColaborador = colaborador.IdColaborador
				evaluacionMetaColaborador.IdEvaluacionPorMeta = evaluacionPorMeta.IdEvaluacionAnual
				db.Create(&evaluacionMetaColaborador)
			}
		}
	}

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

func AceptarEvaluacion(IdEvaluacion string) (models.Evaluacion, error) {

	var result models.Evaluacion

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("UPDATE Evaluaciones SET AceptoEvaluacion = 1,FechaAcepto = GETDATE() WHERE idEvaluacion = ?", IdEvaluacion).Scan(&result)

	return result, nil
}

func NuevaEvaluacionAnual(modelo models.EvaluacionAnual) (models.EvaluacionAnual, error) {

	var result models.EvaluacionAnual
	modelo.FechaCreacion = time.Now()
	modelo.FechaModificacion = time.Now()

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&modelo).Scan(&result)

	if modelo.IdEvaluacionAnual > 0 {
		db.Raw(" exec usp_CrearEvaluacionesAnuales ?", time.Now().Year()).Scan(&result)
	}

	fmt.Println(modelo.IdEvaluacionAnual)

	return result, nil
}

func EliminarEvaluacionPorMetaService(IdEvaluacionMeta string) (models.EvaluacionAnual, error) {

	var result models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	db.Raw(" exec usp_UpdateHumansMegaMistake ?", IdEvaluacionMeta).Scan(&result)

	return result, nil
}

func ResetearNotaEvaluacionPorMetaService(modelo models.ResetearNotaEvaluacion) (models.EvaluacionAnual, error) {

	var result models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	db.Raw(" EXEC usp_ReiniciaEvaluacionMeta ?, ?, ?", modelo.ColaboradorId, modelo.EvluacionId, modelo.EliminadaPor).Scan(&result)
	modelo.TipoNota = 1
	RegistrarReseteoDeNota(modelo)

	return result, nil
}

func ResetearNotaEvaluacionGeneralService(modelo models.ResetearNotaEvaluacion) (models.EvaluacionAnual, error) {

	var result models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	db.Raw(" EXEC usp_ReiniciaEvaluacionMetrica ?, ?, ?", modelo.ColaboradorId, modelo.EvluacionId, modelo.EliminadaPor).Scan(&result)

	modelo.TipoNota = 2
	RegistrarReseteoDeNota(modelo)

	return result, nil
}

func RegistrarReseteoDeNota(modelo models.ResetearNotaEvaluacion) (models.EvaluacionAnual, error) {

	var result models.EvaluacionAnual

	db := config.ConnectDB()
	defer db.Close()

	db.Raw(" EXEC usp_RegistrarReseteoDeNota  ?, ?, ?, ?", modelo.TipoNota, modelo.ColaboradorId, modelo.EliminadaPor, modelo.EvluacionId).Scan(&result)

	return result, nil
}
