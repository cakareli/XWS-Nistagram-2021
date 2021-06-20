package model

type Story struct {
	MediaContent `bson:",inline,omitempty"`
	IsHighlighted bool `bson:"isHighlighted"`
	ForCloseFriends bool `bson:"forCloseFriends"`
}