package model

type Post struct {
	MediaContent `bson:",inline,omitempty"`
	Likes int `bson:"likes"`
	Dislikes int `bson:"dislikes"`
	Comment []Comment `bson:"comments,omitempty"`
}