package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularUserDTO struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:username`
	Password string `json:"password"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Gender *model.Gender `json:"gender"`
	BirthDate *primitive.DateTime `json:"birthDate"`
	UserRole *model.UserRole `json:"userRole"`
	UserType *model.UserType `json:"userType"`
	Biography string `json:"biography"`
	WebSite string `json:"webSite"`
	IsDisabled bool `json:"isDisabled"`
	PrivacyType *model.PrivacyType `json:"privacyType"`
	AllMessageRequests bool `json:"allMessageRequests"`
	TagsAllowed bool `json:"tagsAllowed"`
}