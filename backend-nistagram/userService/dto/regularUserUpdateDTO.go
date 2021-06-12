package dto

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularUserUpdateDTO struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Gender *model.Gender `json:"gender"`
	BirthDate *primitive.DateTime `json:"birthDate"`
	Biography string `json:"biography"`
	WebSite string `json:"webSite"`
}