package models

import "github.com/jinzhu/gorm"

//ResponseMessage for responsemessages
type ResponseMessage struct {
	Status     int    `json:"status" example:"200"`
	Message    string `json:"message" example:"Meaningful message"`
	ResourceID int    `json:"resourceId" example:"1"`
}

//TrackInfoModel for track info
type TrackInfoModel struct {
	gorm.Model
	Trackname      string `json:"trackname"`
	Lenght         int    `json:"length"`
	OfficialRecord string `json:"official_record"`
	Country        string `json:"country" example:"Sweden"`
	LayoutImage    string `gorm:"type:longtext;" json:"layout_image"`
}

//PlayerModel for userinfo and stuff
type PlayerModel struct {
	gorm.Model
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Image    string `gorm:"type:blob;" json:"image"`
	Platform string `json:"platform"`
}

//TimeRegistrations for time registration
type TimeRegistrations struct {
	gorm.Model
	UserID         int    `json:"userID"`
	TrackID        int    `json:"trackID"`
	CarName        string `json:"carname"`
	RegisteredTime string `json:"registered_time"`
}
