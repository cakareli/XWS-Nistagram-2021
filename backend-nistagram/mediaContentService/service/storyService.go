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

func (service *StoryService) FindStoryById(id primitive.ObjectID) (dto.StoryDTO,error) {
	var story dto.StoryDTO
	storyModel, err := service.StoryRepository.FindStoryById(id)
	if err != nil{
		return story,err
	}
	postDto := createStoryDTOFromStoryModel(*storyModel)

	return *postDto, nil
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

func (service *StoryService) DeleteReportedStory(id primitive.ObjectID) error{
	err := service.StoryRepository.DeleteReportedStory(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *StoryService) DeleteStory(id primitive.ObjectID) error{
	err := service.StoryRepository.DeleteStory(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *StoryService) DeleteUserStories(id primitive.ObjectID) error{
	storyDocuments := service.StoryRepository.FindAllStoriesByUserId(id.Hex())
	storyModel := CreateStoriesFromDocuments(storyDocuments)
	err := service.StoryRepository.DeleteUserStories(storyModel)
	if err != nil {
		return err
	}
	return nil
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


func (service *StoryService) GetAllStoryReports() []dto.ReportedStoryDTO {
	storyReports := service.StoryRepository.GetAllStoryReports()

	storyReportsModel := createStoryReportsFromDocuments(storyReports)

	storyReportsDto := createStoryReportedPostDtoFromModel(storyReportsModel)

	return storyReportsDto
}

func (service *StoryService) GetStoryFeed(usersIds []string) (*[]dto.StoryDTO, error) {
	var stories []model.Story
	for i:=0; i < len(usersIds); i++ {
		story := service.StoryRepository.FindAllStoriesByUserId(usersIds[i])
		storyModel := CreateStoriesFromDocuments(story)
		stories = appendStory(stories, storyModel)
	}

	storyDTOs := createStoryDTOsFromStoryModel(stories)
	return storyDTOs, nil
}

func createStoryReportsFromDocuments(ReportDocuments []bson.D) []model.InappropriateContentStory {
	var reportPosts []model.InappropriateContentStory
	for i := 0; i < len(ReportDocuments); i++ {
		var report model.InappropriateContentStory
		bsonBytes, _ := bson.Marshal(ReportDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &report)
		reportPosts = append(reportPosts, report)
	}
	return reportPosts
}

func createStoryDTOsFromStoryModel(storyModels []model.Story) *[]dto.StoryDTO{
	var storyDTOs []dto.StoryDTO
	for i := 0; i < len(storyModels); i++{
		var storyDTO dto.StoryDTO
		storyDTO.Id = storyModels[i].Id.Hex()
		storyDTO.Hashtags = storyModels[i].Hashtags
		storyDTO.Tags = storyModels[i].Tags
		storyDTO.Description = storyModels[i].Description
		storyDTO.MediaPaths = storyModels[i].MediaPaths
		storyDTO.UploadDate = storyModels[i].UploadDate
		storyDTO.MediaContentType = storyModels[i].MediaContentType
		storyDTO.RegularUser = storyModels[i].RegularUser
		storyDTO.Location = storyModels[i].Location
		storyDTOs = append(storyDTOs, storyDTO)
	}
	return &storyDTOs
}

func createStoryReportedPostDtoFromModel(ReportedDocumentsModel []model.InappropriateContentStory) []dto.ReportedStoryDTO{
	var reporedPostsDTO []dto.ReportedStoryDTO
	for i := 0 ;i < len(ReportedDocumentsModel); i++{
		var reportDTO dto.ReportedStoryDTO
		reportDTO.Id = ReportedDocumentsModel[i].Id.Hex()
		reportDTO.Text = ReportedDocumentsModel[i].Text
		reportDTO.Story = ReportedDocumentsModel[i].Story
		reporedPostsDTO = append(reporedPostsDTO, reportDTO)
	}
	return reporedPostsDTO
}

func createStoryDTOFromStoryModel(storyModel model.Story) *dto.StoryDTO{
	var storyDTO dto.StoryDTO
	storyDTO.Id = storyModel.Id.Hex()
	storyDTO.Hashtags = storyModel.Hashtags
	storyDTO.Tags = storyModel.Tags
	storyDTO.Description = storyModel.Description
	storyDTO.MediaPaths = storyModel.MediaPaths
	storyDTO.UploadDate = storyModel.UploadDate
	storyDTO.MediaContentType = storyModel.MediaContentType
	storyDTO.RegularUser = storyModel.RegularUser
	storyDTO.Location = storyModel.Location
	return &storyDTO
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

func appendStory(allPosts []model.Story, newPosts []model.Story) []model.Story{
	for i := 0; i<len(newPosts); i++{
		allPosts = append(allPosts, newPosts[i])
	}
	return allPosts
}