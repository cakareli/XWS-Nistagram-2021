package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
)

type ReportedPostDTO struct {

	Id string `json:"_id"`
	Text string `json:"text"`
	Post model.Post `json:"content"`
}
