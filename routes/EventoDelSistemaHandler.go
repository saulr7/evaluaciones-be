package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
)

func RegistrarEventoDelSistema(w http.ResponseWriter, r *http.Request) {

	var newEvento models.EventoDelSistema
	err := json.NewDecoder(r.Body).Decode(&newEvento)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrecta")
		return
	}
	token, _ := services.GetToken(r)
	newEvento.IdColaborador = token.Usuario.IdColaborador

	var evento, err2 = services.RegistarEventoDelSistema(newEvento)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(evento)

	fmt.Fprint(w, string(response))
}
