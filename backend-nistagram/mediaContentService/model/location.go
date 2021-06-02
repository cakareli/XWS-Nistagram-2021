package model

type Location struct {
	Country string `json:"country"`
	City string `json:"city"`
	StreetName string `json:"streetName"`
	StreetNumber int `json:"streetNumber"`
}
