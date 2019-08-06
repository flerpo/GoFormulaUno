package api

import (
	"GoFormulaUno/kod/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *gorm.DB

//Start the api
func Start(db2 *gorm.DB) {
	db = db2
	r := gin.Default()
	url := ginSwagger.URL("doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	v1Tracks := r.Group("/api/v1/tracks")
	{
		v1Tracks.POST("/", createTrack)
		v1Tracks.GET("/", fetchAllTracks)
		v1Tracks.GET("/:id", fetchSingleTrack)
		// v1.PUT("/:id", updateTrack)
		v1Tracks.DELETE("/:id", deleteTrack)
	}
	v1Players := r.Group("/api/v1/players")
	{
		v1Players.POST("/", createPlayer)
		v1Players.GET("/", fetchAllPlayers)
		v1Players.GET("/:id", fetchSinglePlayer)
		// v1.PUT("/:id", updateTrack)
		v1Players.DELETE("/:id", deletePlayer)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
func createPlayer(c *gin.Context) {
	var json models.PlayerModel
	search := new(models.PlayerModel)
	if c.BindJSON(&json) == nil {
		db.Where("username= ?", json.UserName).First(&search)
		if search.UserName != "" {
			c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Player alreade exists", "resourceId": search.ID})
		} else {
			db.Create(&json)
			c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Player created successfully!", "resourceId": json.ID})
		}
	}
}
func deleteTrack(c *gin.Context) {
	var single models.TrackInfoModel
	db.Find(&single, c.Param("id"))
	db.Delete(&single)
	c.JSON(http.StatusOK, gin.H{"status": "deleted", "message": "Deleted track " + single.Trackname})
}
func deletePlayer(c *gin.Context) {
	var single models.PlayerModel
	db.Find(&single, c.Param("id"))
	db.Delete(&single)
	c.JSON(http.StatusOK, gin.H{"status": "deleted", "message": "Deleted player " + single.UserName})
}
func fetchAllTracks(c *gin.Context) {
	var all []models.TrackInfoModel
	db.Find(&all)
	c.JSON(http.StatusOK, all)
}

func fetchAllPlayers(c *gin.Context) {
	var all []models.PlayerModel
	db.Find(&all)
	c.JSON(http.StatusOK, all)
}

// Get single track godoc
// @Summary Get track
// @Description get track
// @Produce  json
// @Path id query string false "track search by id"
// @Success 200 {array} main.TrackInfoModel
// @Failure 404 {array} main.responseMessage
// @Router /tracks/:id [get]
func fetchSingleTrack(c *gin.Context) {
	var single models.TrackInfoModel
	db.Find(&single, c.Param("id"))
	if single.ID == 0 {
		resp := new(models.ResponseMessage)
		resp.Message = "Track not found"
		resp.Status = http.StatusNotFound
		c.JSON(http.StatusNotFound, resp)
	} else {
		c.JSON(http.StatusOK, single)
	}
}

func fetchSinglePlayer(c *gin.Context) {
	var single models.PlayerModel
	db.Find(&single, c.Param("id"))
	if single.ID == 0 {
		resp := new(models.ResponseMessage)
		resp.Message = "Player not found"
		resp.Status = http.StatusNotFound
		c.JSON(http.StatusNotFound, resp)
	} else {
		c.JSON(http.StatusOK, single)
	}
}

func createTrack(c *gin.Context) {
	var json models.TrackInfoModel
	search := new(models.TrackInfoModel)
	if c.BindJSON(&json) == nil {
		db.Where("trackname= ?", json.Trackname).First(&search)
		if search.Trackname != "" {
			c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Track alreade exists", "resourceId": search.ID})
		} else {
			db.Create(&json)
			c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Track created successfully!", "resourceId": json.ID})
		}
	}
}
