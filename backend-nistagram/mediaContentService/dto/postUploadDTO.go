package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
)

type PostUploadDTO struct {
	Tags []string `bson:"tags,omitempty"`
	Description string `bson:"description,omitempty"`
	MediaPaths []string `bson:"mediaPaths,omitempty"`
	UploadDate *primitive.DateTime `bson:"uploadDate,omitempty"`
	MediaContentType *model.MediaContentType `bson:"mediaContentType,omitempty"`
	Country string `json:"country"`
	City string `json:"city"`
	StreetName string `json:"streetName"`
	StreetNumber int `json:"streetNumber"`
	Username string `bson:"username,omitempty"`
}