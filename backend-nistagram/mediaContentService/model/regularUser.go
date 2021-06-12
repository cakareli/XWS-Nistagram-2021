package model

type RegularUser struct {
	Username string `bson:"username,omitempty"`
	PrivacyType PrivacyType `bson:"privacyType,omitempty"`
}