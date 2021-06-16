package model

type RegularUser struct {
	Id string `bson:"_id,omitempty"`
	PrivacyType *PrivacyType `bson:"privacyType"`
	Username string `bson:"username"`
}