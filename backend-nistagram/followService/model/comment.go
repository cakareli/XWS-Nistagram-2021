package model

type Comment struct {
	Id int `json:"id"`
	Text string `json:"text"`
	RegularUser RegularUser `json:"regularUser"`
}

