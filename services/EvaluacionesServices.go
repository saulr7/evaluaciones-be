package services

import (
	"fmt"

	"../config"
	"../models"
)

func GetEvaluacionPorColaborador(idColaborador string) ([]models.Evaluacion, error) {

	var EvaluacionEnvia []models.Evaluacion

	type Result struct {
		IdEvaluacion  int `gorm:"column:idEvaluacion"`
		IdColaborador int `gorm:"column:idColaborador"`
		Anio          int `gorm:"column:Anio"`
		IdGrado       int `gorm:"column:idGrado"`
	}

	var result []Result
	db := config.ConnectDB()
	defer db.Close()

	db.Raw("EXEC usp_GetEvaluacionPorColaborador ?", idColaborador).Scan(&result)
	for _, dato := range result {

		var Encabezado []models.EncabezadoPreguntas
		var EvaluacionTemp models.Evaluacion

		EvaluacionTemp.IdEvaluacion = dato.IdEvaluacion
		EvaluacionTemp.IdColaborador = dato.IdColaborador
		EvaluacionTemp.Anio = dato.Anio
		EvaluacionTemp.IdGrado = dato.IdGrado

		db.Raw("EXEC usp_GetEncabezadosPorColaborador ?", idColaborador).Scan(&Encabezado)

		for _, Cabeza := range Encabezado {
			fmt.Println(Cabeza)
			var Preguntas []models.Preguntas

			db.Raw("EXEC usp_GetPreguntasPorColaboradorYGrado ?,?", idColaborador, Cabeza.IdGradoPorCompetencia).Scan(&Preguntas)

			for _, Pregunt := range Preguntas {
				fmt.Println(Pregunt)
				var Respuestas []models.Respuestas

				db.Raw("EXEC usp_GetRespuestasPorPregunta ?,?,?", Pregunt.IdTipoRespuesta, Pregunt.IdPregunta, idColaborador).Scan(&Respuestas)
				for _, Respu := range Respuestas {
					fmt.Println(Respu)
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
