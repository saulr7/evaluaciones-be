package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
)

func SubAreas(w http.ResponseWriter, r *http.Request) {

	var subAreas, err = services.SubAreasAll()

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
