package model

type RegularUser struct {
	UserId string `bson:"userId,omitempty"`
	Username string `bson:"username,omitempty"`
	PrivacyType PrivacyType `bson:"privacyType,omitempty"`
}