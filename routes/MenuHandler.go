package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../services"
	"github.com/gorilla/mux"
)

func OpcionesDeMenuByUserHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	colaborador, _ := strconv.Atoi(idColaborador)

	var subAreas, err = services.GetOpcionesDeMenuByUserService(colaborador)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&subAreas)

	fmt.Fprint(w, string(response))
}

func OpcionesDeMenuHandler(w http.ResponseWriter, r *http.Request) {

	token, _ := services.GetToken(r)

	var subAreas, err = services.GetOpcionesDeMenu(token.Usuario.IdColaborador)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&subAreas)

	fmt.Fprint(w, string(response))
}
