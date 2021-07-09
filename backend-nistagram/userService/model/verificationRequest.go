package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type VerificationRequest struct {

	Id primitive.ObjectID `bson:"_id,omitempty"`
	UserId string `bson:"userId,omitempty"`
	Name string `bson:"verificationName,omitempty"`
	Surname string `bson:"verificationSurname,omitempty"`
	ImageUrl string `bson:"verificationImageUrl,omitempty"`
	VerificationType UserType `bson:"verificationType,omitempty"`
}