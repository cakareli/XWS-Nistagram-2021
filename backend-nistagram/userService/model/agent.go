package model

type Agent struct {
	User `bson:",inline"`
	IsDisabled bool `bson:"isDisabled,omitempty"`
	ProfilePrivacy ProfilePrivacy `bson:",inline,omitempty"`
	Verified bool `bson:"verified,omitempty"`
}