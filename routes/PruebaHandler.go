package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
)

func PruebaHandler(w http.ResponseWriter, r *http.Request) {

	var usuarioModel, erro = services.Say_hello()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	fmt.Fprintf(w, string(response))
}
