package models

type LoginRespondModel struct {
	CodigoError   int    `gorm:"column:CodigoError"`
	Clave         string `gorm:"column:Clave"`
	Mensaje       string `gorm:"column:Mensaje"`
	ColaboradorId int    `gorm:"column:ColaboradorId"`
}
