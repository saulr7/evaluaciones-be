package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var credenciales models.UsuarioCredenciales
	err := json.NewDecoder(r.Body).Decode(&credenciales)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Credenciales incorrectas")
		return
	}

	var usuarioModel, erro = services.Login(credenciales)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	fmt.Fprintf(w, string(response))
}

func LoginWithToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var credenciales models.UsuarioCredenciales
	err := json.NewDecoder(r.Body).Decode(&credenciales)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Credenciales incorrectas")
		return
	}

	var usuarioModel, erro = services.LoginWithToken(credenciales)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	fmt.Fprintf(w, string(response))
}
