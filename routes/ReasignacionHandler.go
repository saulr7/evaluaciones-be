package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func GetReasignaciones(w http.ResponseWriter, r *http.Request) {

	var Resultados, erro = services.GetReasignaciones()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)
	fmt.Fprint(w, string(response))
}

func NewReasignacion(w http.ResponseWriter, r *http.Request) {

	var reasignacion models.Reasignacion
	err := json.NewDecoder(r.Body).Decode(&reasignacion)

	fmt.Println(reasignacion)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Datos incorrectas")
		return
	}

	var Resultados, erro = services.NewReasignacion(reasignacion)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)
	fmt.Fprint(w, string(response))
}

func DeleteReasignacion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	IdAsignacion := vars["idAsignacion"]

	var Resultados, erro = services.DeleteReasignacion(IdAsignacion)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)
	fmt.Fprint(w, string(response))
}
