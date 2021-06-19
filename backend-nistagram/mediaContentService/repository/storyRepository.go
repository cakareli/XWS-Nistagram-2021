package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
