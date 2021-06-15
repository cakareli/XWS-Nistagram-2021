package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
)

type RegularUserPostDTO struct {
	Id string `bson:"_id"`
	PrivacyType *model.PrivacyType `bson:"privacyType"`
}