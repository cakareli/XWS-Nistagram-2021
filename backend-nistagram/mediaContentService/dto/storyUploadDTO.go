package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type StoryUploadDTO struct {
	Tags []string `json:"tags"`
	Description string `json:"description"`
	MediaPaths []string `json:"mediaPaths"`
	UploadDate *primitive.DateTime `json:"uploadDate"`
	Location string `json:"location"`
	Username string `json:"username"`
	ForCloseFriends bool `json:"forCloseFriends"`
}
