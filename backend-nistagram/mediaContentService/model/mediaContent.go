package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaContent struct {
	Id int `bson:"id,omitempty"`
	Tags []string `bson:"tags,omitempty"`
	Description string `bson:"description,omitempty"`
	MediaPaths []string `bson:"mediaPaths,omitempty"`
	UploadDate *primitive.DateTime `bson:"uploadDate,omitempty"`
	MediaContentType MediaContentType `bson:"mediaContentType,omitempty"`
	Location Location `bson:"location,omitempty"`
	RegularUser RegularUser `bson:"regularUser,omitempty"`
}