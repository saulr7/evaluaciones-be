package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gorilla/mux"
)

func GetColaboradoresEquipo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetColaboradoresEquipo(idColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetEquipoEvaluacion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]
	idEvaluacionAnual := vars["idEvaluacionAnual"]

	var Resultados, erro = services.GetEquipoEvaluacion(idColaborador, idEvaluacionAnual)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetEquipoCajeros(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetEquipoCajeros(idColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
