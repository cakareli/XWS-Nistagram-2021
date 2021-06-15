package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
)

type RegularUserPostDTO struct {
	Id string `json:"_id"`
	PrivacyType *model.PrivacyType `json:"privacyType"`
}