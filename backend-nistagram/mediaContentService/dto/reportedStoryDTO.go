package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
)

type ReportedStoryDTO struct {

	Id string `json:"_id"`
	Text string `json:"text"`
	Story model.Story `json:"content"`
}
