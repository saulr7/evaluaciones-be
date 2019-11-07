package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mssql", "sqlserver://app_4dxtablero:app_4dxtablero@des-cobbe01:1433?database=BPEvaluaciones")

	if err != nil {
		fmt.Println("Algo salió mal")
		panic(err)
	}

	// defer db.Close()
	db.LogMode(true)
	db.SingularTable(true)

	return db
}

func ConnectDB4DX() *gorm.DB {
	db, err := gorm.Open("mssql", "sqlserver://app_4dxtablero:app_4dxtablero@des-cobbe01:1433?database=4DX")

	if err != nil {
		fmt.Println("Algo salió mal")
		panic(err)
	}

	// defer db.Close()
	db.LogMode(true)
	db.SingularTable(true)

	return db
}
