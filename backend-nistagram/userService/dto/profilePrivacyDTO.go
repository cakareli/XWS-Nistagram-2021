package dto

import "XWS-Nistagram-2021/backend-nistagram/userService/model"

type ProfilePrivacyDTO struct {
	Id string `json:"_id"`
	PrivacyType model.PrivacyType `json:"privacyType"`
	AllMessagesRequests bool `json:"allMessagesRequests"`
	TagsAllowed bool `json:"tagsAllowed"`
}
