package db

import (
	"github.com/flerpo/GoFormulaUno/kod/models"

	"github.com/jinzhu/gorm"
	//for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//Db variable
var Db *gorm.DB 

//GetConnection a db connection
func GetConnection(){
var err error
	//db, err = gorm.Open("mssql", "sqlserver://sa:!_ucKy-!uKe@localhost:1433?database=labb")
	Db, err := gorm.Open("mysql", "root:sofarfromhome@tcp(172.17.0.3:3306)/formula?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	println("skapar databastabeller")
	//Migrate the schema
	Db.AutoMigrate(&models.TrackInfoModel{})
	Db.AutoMigrate(&models.PlayerModel{})
	Db.AutoMigrate(&models.TimeRegistrations{})
}