package model

type InappropriateContent struct {
	Id int `json:"id"`
	Text string `json:"text"`
	MediaContent MediaContent `json:"mediaContent"`
	RegularUser RegularUser `json:"regularUser"`
}

