package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:iceman@tcp(172.17.0.2:3306)/FormulaUno?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&trackInfoModel{})
}

type trackInfoModel struct {
	gorm.Model
	Name           string `json:"name"`
	Lenght         int    `json:"length"`
	OfficialRecord string `json:"official_record"`
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
	length, _ := strconv.Atoi(c.PostForm("length"))
	a := c.PostForm("name")
	println(a)

	track := trackInfoModel{Name: c.PostForm("name"), Lenght: length, OfficialRecord: c.PostForm("official_record")}
	db.Save(&track)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Track created successfully!", "resourceId": track.ID})
}
