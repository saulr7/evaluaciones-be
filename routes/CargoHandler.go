package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
)

func GetCargosHandler(w http.ResponseWriter, r *http.Request) {

	var subAreas, err = services.GetCargos()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&subAreas)

	fmt.Fprint(w, string(response))
}

func GetCargoPadreYEmpresHandler(w http.ResponseWriter, r *http.Request) {

	var cargos, err = services.GetCargoPadreYEmpresaService()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&cargos)

	fmt.Fprint(w, string(response))
}

func NewCargoHandler(w http.ResponseWriter, r *http.Request) {

	token, _ := services.GetToken(r)
	var modelo models.NuevoCargoModel
	err := json.NewDecoder(r.Body).Decode(&modelo)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Datos incorrectas")
		return
	}

	modelo.AgregadoPor = token.Usuario.IdColaborador
	var cargos, err2 = services.NewCargoService(modelo)

	if err2 != nil {
		fmt.Println(err2)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&cargos)

	fmt.Fprint(w, string(response))
}
