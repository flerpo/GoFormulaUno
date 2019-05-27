package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	//db, err = gorm.Open("mssql", "formula:un0Unoun0@tcp(localhost:1433)/FormulaUno?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mssql", "sqlserver://formula:un0Unoun0@localhost:1433?database=labb")

	if err != nil {
		panic(err)
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&trackInfoModel{})
}

type trackInfoModel struct {
	gorm.Model
	Trackname string `json:"trackname"`
	//	Lenght         int    `json:"length"`
	//	OfficialRecord string `json:"official_record"`
}

func main() {
	start()
}

func start() {
	r := gin.Default()

	v1 := r.Group("/api/v1/tracks")
	{
		v1.POST("/", createTrack)
		// v1.GET("/", fetchAllTracks)
		// v1.GET("/:id", fetchSingleTrack)
		// v1.PUT("/:id", updateTrack)
		// v1.DELETE("/:id", deleteTrack)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

// createTodo add a new todo
func createTrack(c *gin.Context) {

	track := trackInfoModel{Trackname: c.PostForm("trackname")}
	println(&track.Trackname)
	db.Create(&track)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Track created successfully!", "resourceId": track.ID})
}
