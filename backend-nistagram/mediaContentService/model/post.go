package model

type Post struct {
	MediaContent
	Likes int `json:"likes"`
	Dislikes int `json:"dislikes"`
	Comment []Comment `json:"comments"`
}