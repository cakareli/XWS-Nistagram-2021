package model

import "time"

type MediaContent struct {
	Id int `json:"id"`
	Tags []string `json:"tags"`
	Description string `json:"description"`
	MediaPaths []string `json:"mediaPaths"`
	UploadDate time.Time `json:"uploadDate"`
	MediaContentType MediaContentType `json:"mediaContentType"`
	Location Location `json:"location"`
}