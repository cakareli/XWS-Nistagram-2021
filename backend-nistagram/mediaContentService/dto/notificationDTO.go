package dto

import "XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"

type NotificationDTO struct {
	MediaContentId string `json:"mediaContentId"`
	Text string `json:"text"`
	UserId string `json:"userId"`
	NotificationType model.NotificationType `json:"notificationType"`
}
