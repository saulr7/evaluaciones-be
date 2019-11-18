package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gorilla/mux"
)

func GetResultadosEvaluacionesHistoricasHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetResultadosEvaluacionesHistoricas(idColaborador)

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

func ColaboradoresPendientesCompletarEvaluacionHanlder(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idEvaluacionAnual := vars["idEvaluacionAnual"]

	var Resultados, erro = services.ColaboradoresPendientesCompletarEvaluacion(idEvaluacionAnual)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}
