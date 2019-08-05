package models

import "github.com/jinzhu/gorm"
//ResponseMessage for responsemessages
type ResponseMessage struct {
	Status     int    `json:"status" example:200`
	Message    string `json:"message" example:"Meaningful message"`
	ResourceID int    `json:"resourceId" example:1`
}
//TrackInfoModel for track info
type TrackInfoModel struct {
	gorm.Model
	Trackname      string `json:"trackname"`
	Lenght         int    `json:"length"`
	OfficialRecord string `json:"official_record"`
	Country        string `json:"country" example:"Sweden"`
	LayoutImage    string `gorm:"size:10000" json:"layout_image"`
}
