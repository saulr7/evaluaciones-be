package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Reasignacion struct {
	IdAsignacion      int       `gorm:"column:idAsignacion"`
	IdColaborador     string    `gorm:"column:idColaborador"`
	IdSubAreaResta    int       `gorm:"column:idSubAreaResta"`
	SubAreaOrigen     string    `gorm:"column:SubAreaOrigen"`
	IdSubAreaSuma     int       `gorm:"column:idSubAreaSuma"`
	SubAreaDestino    string    `gorm:"column:SubAreaDestino"`
	IdUsuarioModifico string    `gorm:"column:idUsuarioModifico"`
	FechaModificacion time.Time `gorm:"column:FechaModificacion"`
	Activo            bool      `gorm:"column:Activo"`
	NombreColaborador string    `gorm:"column:NombreColaborador"`
	Modifica          string    `gorm:"column:Modifica"`
}
