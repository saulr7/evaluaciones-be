package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gorilla/mux"
)

func ColaboradoresPorArea(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	AreaId := vars["AreaId"]

	var subAreas, err = services.ColaboradoresPorArea(AreaId)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&subAreas)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetColaboradoresSubArea(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idSubArea := vars["idSubArea"]

	var usuarioModel, erro = services.GetColaboradoresSubArea(idSubArea)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
