package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
)

func GetGetGrados(w http.ResponseWriter, r *http.Request) {

	var subAreas, err = services.GetGetGrados()

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
