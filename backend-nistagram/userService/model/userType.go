package model

type UserType int

const(
	NoVerification UserType = iota
	Influencer
	Sports
	Business
	Brand
	Organization
	NewMedia
)