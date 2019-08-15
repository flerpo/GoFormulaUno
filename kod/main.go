package main

import (
	"GoFormulaUno/kod/api"

	//"github.com/flerpo/GoFormulaUno/kod/api"
	"github.com/flerpo/GoFormulaUno/kod/db"
	_ "github.com/flerpo/GoFormulaUno/kod/docs" // docs is generated by Swag CLI, you have to import it.

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db.GetConnection()
}

// @title GoFormulaUno
// @version 1.0
// @description Used for GoFormululaUno.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	api.Start()
}
