package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gorilla/mux"
)

func OpcionesDeMenuHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var subAreas, err = services.GetOpcionesDeMenu(idColaborador)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&subAreas)

	fmt.Fprint(w, string(response))
}
