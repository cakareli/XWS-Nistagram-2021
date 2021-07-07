package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type InappropriateContentStory struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Text string `bson:"text"`
	Story Story `bson:"story"`
	RegularUser RegularUser `bson:"regularUser"`
}