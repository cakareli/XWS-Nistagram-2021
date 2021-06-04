package model

type Collection struct {
	Id int `json:"id"`
	Name string `json:"name"`
	SavedPosts []Post `json:"savedPosts"`
}
