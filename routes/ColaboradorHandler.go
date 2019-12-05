package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func GetColaboradoresSubArea(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idSubArea := vars["idSubArea"]

	var usuarioModel, erro = services.GetColaboradoresSubArea(idSubArea)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

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

func UpdateColaboradorActivarHandler(w http.ResponseWriter, r *http.Request) {

	reqToken := r.Header.Get("Authorization")

	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "No se ha proporcionado el token")
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])

	token, _ := services.ExtractClaims(reqToken)

	data, _ := json.Marshal(token)

	var result services.Claims
	json.Unmarshal(data, &result)

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

	var cargosGrado, erro = services.UpdateColaboradorActivar(actualizar.ColaboradorId, actualizar.Activar, result.Usuario.IdColaborador)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&cargosGrado)

	fmt.Fprint(w, string(response))
}

func CreateColaboradorHandler(w http.ResponseWriter, r *http.Request) {

	token, _ := services.GetToken(r)

	type ActualizarModel struct {
		ColaboradorId string
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

	fmt.Println(actualizar)

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
