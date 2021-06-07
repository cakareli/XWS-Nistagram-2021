package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Surname string `bson:"surname,omitempty"`
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
	Email string `bson:"email,omitempty"`
	PhoneNumber string `bson:"phoneNumber,omitempty"`
	Gender *Gender `bson:"gender,omitempty"`
	BirthDate *primitive.DateTime `bson:"birthDate,omitempty"`
	UserRole UserRole `bson:"userRole,omitempty"`
	Biography string `bson:"biography,omitempty"`
	WebSite string `bson:"webSite,omitempty"`
}

