package model

type MediaContentType int

const(
	PostType MediaContentType = iota
	StoryType
	AlbumType
)