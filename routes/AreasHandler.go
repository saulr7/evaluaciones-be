package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
)

func GetGetAreasConSuGrado(w http.ResponseWriter, r *http.Request) {

	var subAreas, err = services.GetGetAreasConSuGrado()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&subAreas)
	fmt.Fprint(w, string(response))
}

func UpdateGradoAreaHandler(w http.ResponseWriter, r *http.Request) {

	var areaGrado models.AreaGrado
	err := json.NewDecoder(r.Body).Decode(&areaGrado)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	var res, er = services.UpdateAreaGrado(areaGrado)

	if er != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&res)
	fmt.Fprint(w, string(response))
}

func GetAreasHandler(w http.ResponseWriter, r *http.Request) {

	var resultado, err = services.GetAreas()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&resultado)
	fmt.Fprint(w, string(response))
}
