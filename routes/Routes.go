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

	return myRouter
}
