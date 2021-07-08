package service

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/dto"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryService struct {
	StoryRepository *repository.StoryRepository
	NotificationRepository *repository.NotificationRepository
}

func (service *StoryService) GetAllRegularUserStories(username string) []model.Story {
	regularUserStoryDocuments := service.StoryRepository.GetAllStoriesByUsername(username)

	regularUserStories := CreateStoriesFromDocuments(regularUserStoryDocuments)
	return regularUserStories
}

func (service *StoryService) CreateNewStory(storyUploadDTO dto.StoryUploadDTO) error {
	fmt.Println("Creating new story")

	story, err := createStoryFromStoryUploadDTO(&storyUploadDTO)
	if err != nil {
		return err
	}
	storyId, err1 := service.StoryRepository.CreateStory(story)
	if err1 != nil {
		return err
	}

	followersIds, err2 := getUserFollowersWithNotificationsTurnedOn(story.RegularUser.Id)
	if err2 != nil {
		return err2
	}

	notification := CreateNotificationFromEvent(storyId, story.RegularUser.Id, followersIds, model.NotificationType(1))
	err3 := service.NotificationRepository.CreateNotification(notification)
	if err3 != nil {
		return err3
	}

	return nil
}

func (service *StoryService) UpdateStoriesPrivacy(privacyUpdateDTO *dto.PrivacyUpdateDTO) error {
	fmt.Println("updating stories privacy...")
	userStoriesDocuments := service.StoryRepository.FindAllStoriesByUserId(privacyUpdateDTO.Id)
	userStories := CreatePostsFromDocuments(userStoriesDocuments)

	var privacyType model.PrivacyType
	if privacyUpdateDTO.PrivacyType == "0" {
		privacyType = model.PrivacyType(0)
	} else {
		privacyType = model.PrivacyType(1)
	}
	for i:=0; i < len(userStories); i++ {
		err := service.StoryRepository.UpdateStoryPrivacy(privacyType, userStories[i].Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateStoriesFromDocuments(StoriesDocuments []bson.D) []model.Story {
	var publicStories []model.Story
	for i := 0; i < len(StoriesDocuments); i++ {
		var story model.Story
		bsonBytes, _ := bson.Marshal(StoriesDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &story)
		publicStories = append(publicStories, story)
	}
	return publicStories
}

func (service *StoryService) ReportStory(storyReportDTO dto.StoryReportDTO) error {
	fmt.Println("Disliking post...")

	storyId, _ := primitive.ObjectIDFromHex(storyReportDTO.StoryId)
	story, err := service.StoryRepository.FindStoryById(storyId)
	if err != nil {
		return err
	}

	regularUser, err := getRegularUserFromUsername(storyReportDTO.Username)
	if err != nil {
		return err
	}
	var inappropriateContentStory model.InappropriateContentStory
	inappropriateContentStory.Story = *story
	inappropriateContentStory.RegularUser = *regularUser
	inappropriateContentStory.Text = storyReportDTO.Text

	err = service.StoryRepository.CreateStoryReport(&inappropriateContentStory)
	if err != nil {
		return err
	}
	return nil
}

func createStoryFromStoryUploadDTO(storyUploadDTO *dto.StoryUploadDTO) (*model.Story, error){
	regularUser, err := getRegularUserFromUsername(storyUploadDTO.Username)
	if err != nil {
		return nil, err
	}
	var story model.Story
	story.Hashtags = storyUploadDTO.Hashtags
	var tags []model.RegularUser
	for i := 0; i < len(storyUploadDTO.Tags); i++{
		var regUser *model.RegularUser
		regUser, err = getRegularUserFromUsername(storyUploadDTO.Tags[i])
		regUser.Username = storyUploadDTO.Tags[i]
		tags = append(tags,*regUser)

	}
	story.Tags = tags
	story.Description = storyUploadDTO.Description
	story.MediaPaths = storyUploadDTO.MediaPaths
	story.UploadDate = storyUploadDTO.UploadDate
	story.Location = storyUploadDTO.Location
	story.RegularUser = *regularUser
	story.RegularUser.Username = storyUploadDTO.Username
	story.ForCloseFriends = storyUploadDTO.ForCloseFriends
	story.IsHighlighted = false
	story.MediaContentType = model.MediaContentType(1)
	return &story, nil
}