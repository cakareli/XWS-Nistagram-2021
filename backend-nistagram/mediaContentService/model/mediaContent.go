package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaContent struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Hashtags []string `bson:"hashtags,omitempty"`
	Tags []RegularUser `bson:"tags,omitempty"`
	Description string `bson:"description,omitempty"`
	MediaPaths []string `bson:"mediaPaths,omitempty"`
	UploadDate *primitive.DateTime `bson:"uploadDate,omitempty"`
	MediaContentType MediaContentType `bson:"mediaContentType,omitempty"`
	Location string `bson:"location,omitempty"`
	RegularUser RegularUser `bson:"regularUser,omitempty"`
}