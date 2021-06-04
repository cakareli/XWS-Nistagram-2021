package model

type ProfilePrivacy struct {
	PrivacyType PrivacyType `json:"privacyType"`
	AllMessageRequests bool `json:"allMessageRequests"`
	TagsAllowed bool `json:"tagsAllowed"`
}