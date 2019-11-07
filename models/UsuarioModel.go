package models

import (
	_ "github.com/jinzhu/gorm"
)

type Usuario struct {
	Empleado          int    `gorm:"column:Empleado"`
	EmpleadoId        string `gorm:"column:UsrCod"`
	EmpleadoNombre    string `gorm:"column:UsrNom"`
	Correo            string `gorm:"column:Correo"`
	PerfilId          int    `gorm:"column:PerfilCod"`
	PerfilDescripcion string `gorm:"column:PerfilDesc"`
	ZonaId            int    `gorm:"column:zona"`
	ZonaDescripcion   string `gorm:"column:Zona_Descripcion"`
	CambiarClave      bool   `gorm:"column:CambiarClave"`
	EsLider           bool   `gorm:"column:EsLider"`
	SubAreaId         int    `gorm:"column:idSubArea"`
}
