package models

type SubArea struct {
	IdSubArea      int    `gorm:"column:idSubArea"`
	IdArea         int    `gorm:"column:idArea"`
	SubArea        string `gorm:"column:SubArea"`
	Activo         bool   `gorm:"column:activo"`
	IdSubAreaPadre int    `gorm:"column:idSubAreaPadre"`
	Nivel          int    `gorm:"column:Nivel"`
	IdZona         int    `gorm:"column:idZona"`
}

func (SubArea) TableName() string {
	return "SubAreas"
}
