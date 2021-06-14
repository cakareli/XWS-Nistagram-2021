package model

type Post struct {
	MediaContent `bson:",inline,omitempty"`
	Likes int `bson:"likes,omitempty"`
	Dislikes int `bson:"dislikes,omitempty"`
	Comment []Comment `bson:"comments,omitempty"`
}