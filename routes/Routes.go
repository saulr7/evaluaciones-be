package routes

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/prueba", PruebaHandler).Methods("GET")
	myRouter.HandleFunc("/login", LoginHandler)
	myRouter.HandleFunc("/GetColaboradoresEquipo/{idColaborador}", GetColaboradoresEquipo).Methods("GET")
	myRouter.HandleFunc("/reasignaciones", GetReasignaciones).Methods("GET")
	myRouter.HandleFunc("/deleteReasignacion/{idAsignacion}", DeleteReasignacion).Methods("GET")
	myRouter.HandleFunc("/SubAreas", SubAreas).Methods("GET")
	myRouter.HandleFunc("/GetColaboradoresSubArea/{idSubArea}", GetColaboradoresSubArea).Methods("GET")
	myRouter.HandleFunc("/NewReasignacion", NewReasignacion).Methods("POST")
	myRouter.HandleFunc("/GetCompetenciasPorColaborador/{idColaborador}", GetCompetenciasPorColaborador).Methods("GET")
	myRouter.HandleFunc("/GetEvaluacionPorColaborador/{idColaborador}/{idEvaluacionAnual}", GetEvaluacionPorColaborador).Methods("GET")
	myRouter.HandleFunc("/GetEvaluacionAnual/{idColaborador}", GetEvaluacionAnual).Methods("GET")

	return myRouter
}
