package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type InappropriateContentPost struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Text string `bson:"text"`
	Post Post `bson:"post"`
	RegularUser RegularUser `bson:"regularUser"`
}

