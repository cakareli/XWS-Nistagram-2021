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

func (repository *StoryRepository) CreateStory(story *model.Story) (string, error) {
	storiesCollection := repository.Database.Collection("stories")
	res, err := storiesCollection.InsertOne(context.TODO(), &story)
	if err != nil {
		return "", fmt.Errorf("Story is NOT created!")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
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
