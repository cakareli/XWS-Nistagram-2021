package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type StoryRepository struct {
	Database *mongo.Database
}

func (repository *StoryRepository) CreateStory(story *model.Story) error {
	storiesCollection := repository.Database.Collection("stories")
	_, err := storiesCollection.InsertOne(context.TODO(), &story)
	if err != nil {
		return fmt.Errorf("Story is NOT created!")
	}
	return nil
}

func (repository *StoryRepository) UpdateStoryPrivacy(privacyType model.PrivacyType, storyId primitive.ObjectID) error {
	postCollection := repository.Database.Collection("stories")

	updatedStory, err := postCollection.UpdateOne(context.TODO(), bson.M{"_id": storyId},
		bson.D{
			{"$set", bson.D{{"regularUser.privacyType", privacyType}}},
		})
	if err != nil {
		log.Fatal(err)
		return err
	} else if updatedStory.MatchedCount == 0 {
		return fmt.Errorf("Story does not exist!")
	}
	return nil
}

func (repository *StoryRepository) FindAllStoriesByUserId(id string) []bson.D{
	storiesCollection := repository.Database.Collection("stories")
	filterCursor, err := storiesCollection.Find(context.TODO(), bson.M{"regularUser._id": id})
	if err != nil {
		log.Fatal(err)
	}

	var storiesFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &storiesFiltered); err != nil {
		log.Fatal(err)
	}
	return storiesFiltered
}

func (repository *StoryRepository) GetAllStoriesByUsername(username string) []bson.D{
	storiesCollection := repository.Database.Collection("stories")
	filterCursor, err := storiesCollection.Find(context.TODO(), bson.M{"regularUser.username": username})
	if err != nil {
		log.Fatal(err)
	}

	var storiesFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &storiesFiltered); err != nil {
		log.Fatal(err)
	}
	return storiesFiltered
}

func (repository *StoryRepository) FindStoryById(storyId primitive.ObjectID) (*model.Story, error){
	var story *model.Story
	result:= repository.Database.Collection("stories").FindOne(context.Background(), bson.M{"_id": storyId})
	result.Decode(&story)

	if story == nil {
		return nil, fmt.Errorf("Story is NOT found!")
	}
	return story, nil
}

func (repository *StoryRepository) CreateStoryReport(inappropriateContentStory *model.InappropriateContentStory) error {
	storyReportsCollection := repository.Database.Collection("storyReports")
	_, err := storyReportsCollection.InsertOne(context.TODO(), &inappropriateContentStory)
	if err != nil {
		return fmt.Errorf("Story report is NOT created")
	}
	return nil
}
