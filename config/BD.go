package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	//db, err := gorm.Open("mssql", "sqlserver://app_BPEvaluaciones:Bp3v@Lu8c1oN3$@PRD-APLICABE04:1433?database=BPEvaluaciones")
	// db, err := gorm.Open("mssql", "sqlserver://app_BPEvaluaciones:Bp3v@Lu8c1oN3$@pu-aplicabe04:1433?database=BPEvaluaciones")
	db, err := gorm.Open("mssql", "sqlserver://app_4dxtablero:app_4dxtablero@des-cobbe01:1433?database=BPEvaluaciones")

	if err != nil {
		fmt.Println("Algo salió mal")
		panic(err)
	}

	db.LogMode(true)
	db.SingularTable(true)

	return db
}

func ConnectDBEO() *gorm.DB {
	//db, err := gorm.Open("mssql", "sqlserver://app_BPEvaluaciones:Bp3v@Lu8c1oN3$@PRD-APLICABE04:1433?database=EstructuraOrganizacional")
	// db, err := gorm.Open("mssql", "sqlserver://app_BPEvaluaciones:Bp3v@Lu8c1oN3$@pu-aplicabe04:1433?database=EstructuraOrganizacional")
	db, err := gorm.Open("mssql", "sqlserver://app_4dxtablero:app_4dxtablero@des-cobbe01:1433?database=EstructuraOrganizacional")

	if err != nil {
		fmt.Println("Algo salió mal")
		panic(err)
	}

	db.LogMode(true)
	db.SingularTable(true)

	return db
}
