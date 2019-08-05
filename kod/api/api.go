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
	v1 := r.Group("/api/v1/tracks")
	{
		v1.POST("/", createTrack)
		v1.GET("/", fetchAllTracks)
		v1.GET("/:id", fetchSingleTrack)
		// v1.PUT("/:id", updateTrack)
		v1.DELETE("/:id", deleteTrack)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}

func deleteTrack(c *gin.Context) {
	var single models.TrackInfoModel
	db.Find(&single, c.Param("id"))
	db.Delete(&single)
	c.JSON(http.StatusOK, gin.H{"status": "deleted", "message": "Deleted track " + single.Trackname})
}

func fetchAllTracks(c *gin.Context) {
	var all []models.TrackInfoModel
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
