package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../models"

	"../services"
	"github.com/gorilla/mux"
)

func GetColaboradoresEquipo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetColaboradoresEquipo(idColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetEquipoEvaluacion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]
	idEvaluacionAnual := vars["idEvaluacionAnual"]

	var Resultados, erro = services.GetEquipoEvaluacion(idColaborador, idEvaluacionAnual)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetEquipoCajeros(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]

	var Resultados, erro = services.GetEquipoCajeros(idColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetUsuariosHandler(w http.ResponseWriter, r *http.Request) {

	var Resultados, erro = services.GetUsuariosService()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}

func CreateUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	var model models.NuevoUsuarioModel
	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	token, _ := services.GetToken(r)

	if !services.EsAdminBPEvaluaciones(token.Usuario) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, errors.New("No eres administrador"))
	}

	model.AgreadoPor = token.Usuario.IdColaborador
	model.CambiarClave = true

	var Resultados, erro = services.CreateUsuarioService(model)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&Resultados)

	fmt.Fprint(w, string(response))
}

func UpdateUsuarioActivarHandler(w http.ResponseWriter, r *http.Request) {

	token, _ := services.GetToken(r)
	type ActualizarModel struct {
		Usuario string
		Activar bool
	}

	var actualizar ActualizarModel
	err := json.NewDecoder(r.Body).Decode(&actualizar)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	var resul, erro = services.UpdateUsuarioActivar(actualizar.Usuario, actualizar.Activar, token.Usuario.IdColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&resul)

	fmt.Fprint(w, string(response))
}

func CambiarContrasenaHandler(w http.ResponseWriter, r *http.Request) {

	var model models.CambiarContrasenaModel

	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	token, _ := services.GetToken(r)

	model.ColaboradorId = token.Usuario.IdColaborador
	model.Usuario = token.Usuario.Usuario

	var resp, erro = services.CambiarContrasenaService(model)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&resp)

	fmt.Fprint(w, string(response))
}

func ResetearContrasenaHandler(w http.ResponseWriter, r *http.Request) {

	var model models.ResetearContrasenaModel
	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	token, _ := services.GetToken(r)

	if !services.EsAdminBPEvaluaciones(token.Usuario) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, errors.New("No eres administrador"))
	}

	model.ModificadoPor = token.Usuario.Usuario

	var resp, erro = services.ResetearContrasenaService(model)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&resp)

	fmt.Fprint(w, string(response))
}
