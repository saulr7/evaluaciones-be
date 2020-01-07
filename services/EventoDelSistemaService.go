package services

import (
	"time"

	"../config"
	"../models"
)

func RegistarEventoDelSistema(Modelo models.EventoDelSistema) (models.EventoDelSistema, error) {

	Modelo.Id = 0
	Modelo.Fecha = time.Now()

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}
