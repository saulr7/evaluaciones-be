package models

import (
	_ "github.com/jinzhu/gorm"
)

type AreaGrado struct {
	IdArea  int `gorm:"column:idArea"`
	IdGrado int `gorm:"column:idGrado"`
}
