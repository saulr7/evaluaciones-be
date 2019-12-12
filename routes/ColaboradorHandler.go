package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func GetColaboradoresPorAreaHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	AreaId := vars["AreaId"]

	var subAreas, err = services.GetColaboradoresPorArea(AreaId)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&subAreas)

	fmt.Fprint(w, string(response))
}

func GetColaboradoresPorCargoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	CargoId := vars["CargoId"]

	var usuarioModel, erro = services.GetColaboradoresPorCargo(CargoId)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	fmt.Fprint(w, string(response))
}

func GetCargosGradosHandler(w http.ResponseWriter, r *http.Request) {

	var cargosGrado, erro = services.GetCargosGrados()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&cargosGrado)

	fmt.Fprint(w, string(response))
}

func UpdateCargosGradosHandler(w http.ResponseWriter, r *http.Request) {

	var cargoGradoActualizarModel models.CargoGradoActualizarModel
	err := json.NewDecoder(r.Body).Decode(&cargoGradoActualizarModel)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	var cargosGrado, erro = services.UpdateCargosGradosService(cargoGradoActualizarModel)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&cargosGrado)

	fmt.Fprint(w, string(response))
}

func GetColaboradoresInfoHandler(w http.ResponseWriter, r *http.Request) {

	var cargosGrado, erro = services.GetColaboradoresInfo()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&cargosGrado)

	fmt.Fprint(w, string(response))
}

func GetColaboradoresSinUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	var cargosGrado, erro = services.GetColaboradoresSinUsuarioService()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&cargosGrado)

	fmt.Fprint(w, string(response))
}

func GetColaboradoresCargoHandler(w http.ResponseWriter, r *http.Request) {

	var cargosGrado, erro = services.GetColaboradoresCargoService()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&cargosGrado)

	fmt.Fprint(w, string(response))
}

func UpdateColaboradorActivarHandler(w http.ResponseWriter, r *http.Request) {

	token, _ := services.GetToken(r)

	type ActualizarModel struct {
		ColaboradorId string
		Activar       bool
	}

	var actualizar ActualizarModel
	err := json.NewDecoder(r.Body).Decode(&actualizar)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	var resul, erro = services.UpdateColaboradorActivar(actualizar.ColaboradorId, actualizar.Activar, token.Usuario.IdColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&resul)

	fmt.Fprint(w, string(response))
}

func CreateColaboradorHandler(w http.ResponseWriter, r *http.Request) {

	token, _ := services.GetToken(r)

	type ActualizarModel struct {
		ColaboradorId int
		Activar       bool
	}

	var actualizar models.NuevoColaboradorModel
	err := json.NewDecoder(r.Body).Decode(&actualizar)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Los datos proporcionados no son correctos")
		return
	}

	actualizar.AgregadoPor = token.Usuario.IdColaborador

	var cargosGrado, erro = services.CreateColaboradorService(actualizar)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&cargosGrado)

	fmt.Fprint(w, string(response))
}
