package models

import "time"

type EventoDelSistema struct {
	Id            int       `gorm:"AUTO_INCREMENT;column:Id"`
	IdColaborador int       `gorm:"column:IdColaborador"`
	Evento        string    `gorm:"column:Evento"`
	Fecha         time.Time `gorm:"column:Fecha"`
}

func (EventoDelSistema) TableName() string {
	return "EventosDelSistema"
}
