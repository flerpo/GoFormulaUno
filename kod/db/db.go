package db

import (
	"github.com/flerpo/GoFormulaUno/kod/models"
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	//for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Db variable
var db *gorm.DB

func init() {
	var b = GetConnection()
	println("skapar databastabeller")
	//Migrate the schema
	b.AutoMigrate(&models.TrackInfoModel{})
	b.AutoMigrate(&models.PlayerModel{})
	b.AutoMigrate(&models.TimeRegistrations{})
	b.Close()
}

//GetConnection a db connection
func GetConnection() *gorm.DB {
	var err error
	var sqlpass = os.Getenv("SQLPASS")
	var sqlserver = os.Getenv("SQLSERVER")
	fmt.Println(sqlserver)
	//db, err = gorm.Open("mssql", "sqlserver://sa:!_ucKy-!uKe@localhost:1433?database=labb")
	//db, err := gorm.Open("mysql", "formula:sofarfrompassword@tcp(172.17.0.2:3306)/formula?charset=utf8&pa	rseTime=True")
	db, err := gorm.Open("mysql", "formula:" + sqlpass + "@tcp(" + sqlserver + ")/formula?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}
	return db
}
