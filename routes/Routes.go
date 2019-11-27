package routes

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/prueba", PruebaHandler).Methods("GET")
	myRouter.HandleFunc("/login", LoginHandler)
	myRouter.HandleFunc("/GetColaboradoresEquipo/{idColaborador}", GetColaboradoresEquipo).Methods("GET")
	myRouter.HandleFunc("/GetEquipoEvaluacion/{idColaborador}/{idEvaluacionAnual}", GetEquipoEvaluacion).Methods("GET")
	myRouter.HandleFunc("/GetEquipoCajeros/{idColaborador}", GetEquipoCajeros).Methods("GET")
	myRouter.HandleFunc("/reasignaciones", GetReasignaciones).Methods("GET")
	myRouter.HandleFunc("/deleteReasignacion/{idAsignacion}", DeleteReasignacion).Methods("GET")
	myRouter.HandleFunc("/SubAreas", SubAreas).Methods("GET")
	myRouter.HandleFunc("/GetColaboradoresSubArea/{idSubArea}", GetColaboradoresSubArea).Methods("GET")
	myRouter.HandleFunc("/NewReasignacion", NewReasignacion).Methods("POST")
	myRouter.HandleFunc("/GetCompetenciasPorColaborador/{idColaborador}", GetCompetenciasPorColaborador).Methods("GET")
	myRouter.HandleFunc("/GetEvaluacionPorColaborador/{idColaborador}/{idEvaluacionAnual}", GetEvaluacionPorColaborador).Methods("GET")
	myRouter.HandleFunc("/GetEvaluacionAnual/{idColaborador}", GetEvaluacionAnual).Methods("GET")
	myRouter.HandleFunc("/GetEvaluacionesCompletas/{idColaborador}", GetEvaluacionesCompletas).Methods("GET")
	myRouter.HandleFunc("/EvaluacionCompletada", EvaluacionCompletadaHandler).Methods("POST")
	myRouter.HandleFunc("/NuevaEvaluacionAnual", NuevaEvaluacionAnualHandler).Methods("POST")
	myRouter.HandleFunc("/NuevaEvaluacionPorMeta", NuevaEvaluacionPorMetaHandler).Methods("POST")
	myRouter.HandleFunc("/GetEvaluacionMetaPorColaborador/{idColaborador}/{idPadre}", GetEvaluacionMetaPorColaborador).Methods("GET")
	myRouter.HandleFunc("/GetResultadosEvaluacionesHistoricas/{idColaborador}", GetResultadosEvaluacionesHistoricasHandler).Methods("GET")
	myRouter.HandleFunc("/GetColaboradoresPendientesCompletarEvaluacion/{idEvaluacionAnual}", ColaboradoresPendientesCompletarEvaluacionHanlder).Methods("GET")
	myRouter.HandleFunc("/GetEvaluacionesTodas", GetEvaluacionesTodasHanlder).Methods("GET")
	myRouter.HandleFunc("/GetOpcionesDeMenu/{idColaborador}", OpcionesDeMenuHandler).Methods("GET")
	myRouter.HandleFunc("/AceptarEvaluacion/{idEvaluacion}", AceptarEvaluacionHanlder).Methods("GET")
	myRouter.HandleFunc("/GetRptCompetencias/{idColaborador}", GetRptCompetencias).Methods("GET")
	myRouter.HandleFunc("/GetGetAreasConSuGrado", GetGetAreasConSuGrado).Methods("GET")
	myRouter.HandleFunc("/GetGetGrados", GetGetGrados).Methods("GET")
	return myRouter
}
