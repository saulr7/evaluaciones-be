package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func DeleteReasignacion(w http.ResponseWriter, r *http.Request) {

	fmt.Println("eNt")

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
