package model

type Story struct {
	MediaContent
	IsHighlighted bool `json:"isHighlighted"`
	ForCloseFriends bool `json:"forCloseFriends"`
}