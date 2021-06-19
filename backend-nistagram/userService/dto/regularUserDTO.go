package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type RegularUserDTO struct {
	Id primitive.ObjectID `json:"_id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
}
