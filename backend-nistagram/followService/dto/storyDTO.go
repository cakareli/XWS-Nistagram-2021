package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryDTO struct {
	Id string `json:"Id"`
	Hashtags []string `json:"Hashtags"`
	Tags []model.RegularUser `json:"Tags"`
	Description string `json:"Description"`
	MediaPaths []string `json:"MediaPaths"`
	UploadDate *primitive.DateTime `json:"UploadDate"`
	MediaContentType model.MediaContentType `json:"MediaContentType"`
	Location string `json:"Location"`
	RegularUser model.RegularUser `json:"RegularUser"`
}
