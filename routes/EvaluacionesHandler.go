package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func GetEvaluacionPorColaborador(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]
	idEvaluacionAnual := vars["idEvaluacionAnual"]

	var Resultados, erro = services.GetEvaluacionPorColaborador(idColaborador, idEvaluacionAnual)

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

func GetEvaluacionAnual(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetEvaluacionAnual(idColaborador)

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

func GetEvaluacionesCompletas(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetEvaluacionesCompletas(idColaborador)

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

func EvaluacionCompletadaHandler(w http.ResponseWriter, r *http.Request) {

	var evaluacionCompletada models.EvaluacionCompletada
	err := json.NewDecoder(r.Body).Decode(&evaluacionCompletada)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Credenciales incorrectas")
		return
	}

	var respuesta, erro = services.GuardarEvaluacionCompletada(evaluacionCompletada)

	if erro != nil || respuesta == false {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&respuesta)

	fmt.Fprintf(w, string(response))
}

func NuevaEvaluacionPorMetaHandler(w http.ResponseWriter, r *http.Request) {

	var evaluacionPorMeta models.NuevaEvaluacionPorMeta
	err := json.NewDecoder(r.Body).Decode(&evaluacionPorMeta)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	var respuesta, erro = services.NuevaEvaluacionPorMeta(evaluacionPorMeta)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&respuesta)

	fmt.Fprintf(w, string(response))
}

func GetEvaluacionMetaPorColaborador(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]
	idPadre := vars["idPadre"]

	var Resultados, erro = services.GetEvaluacionMetaPorColaborador(idColaborador, idPadre)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)
	// responseString :=

	fmt.Fprint(w, string(response))
}

func GetEvaluacionesTodasHanlder(w http.ResponseWriter, r *http.Request) {

	var Resultados, erro = services.GetEvaluacionsTodas()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)
	// responseString :=

	fmt.Fprint(w, string(response))
}
