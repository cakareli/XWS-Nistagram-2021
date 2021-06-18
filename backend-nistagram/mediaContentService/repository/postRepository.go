package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

type PostRepository struct {
	Database *mongo.Database
}

func (repository *PostRepository) Create(post *model.Post) error {
	regularUserCollection := repository.Database.Collection("posts")
	_, err := regularUserCollection.InsertOne(context.TODO(), &post)
	if err != nil {
		return fmt.Errorf("post is NOT created")
	}
	return nil
}

func (repository *PostRepository) GetAllByUsername(username string) []bson.D{
	postsCollection := repository.Database.Collection("posts")
	filterCursor, err := postsCollection.Find(context.TODO(), bson.M{"regularUser.username": username})
	if err != nil {
		log.Fatal(err)
	}
	
	var postsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &postsFiltered); err != nil {
		log.Fatal(err)
	}
	return postsFiltered
}

func (repository *PostRepository) GetAllPublic() []bson.D{
	postsCollection := repository.Database.Collection("posts")
	filterCursor, err := postsCollection.Find(context.TODO(), bson.M{"regularUser.privacyType": 0})
	if err != nil {
		log.Fatal(err)
	}

	var postsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &postsFiltered); err != nil {
		log.Fatal(err)
	}
	return postsFiltered
}

func (repository *PostRepository) GetLocationSearchResults(searchInput string, allPublicPosts []model.Post) []model.Post{
	var searchResult []model.Post
	for i:=0; i < len(allPublicPosts); i++ {
		if (strings.Contains(strings.ToLower(allPublicPosts[i].Location), strings.ToLower(searchInput))){
			searchResult = append(searchResult, allPublicPosts[i])
		}
	}
	return searchResult
}

func (repository *PostRepository) GetUserSearchResults(searchInput string, allPublicPosts []model.Post) []model.Post{
	var searchResult []model.Post
	for i:=0; i < len(allPublicPosts); i++ {
		if (strings.Contains(strings.ToLower(allPublicPosts[i].RegularUser.Username), strings.ToLower(searchInput))){
			searchResult = append(searchResult, allPublicPosts[i])
		}
	}
	return searchResult
}

func (repository *PostRepository) GetTagSearchResults(searchInput string, allPublicPosts []model.Post) []model.Post{
	var searchResult []model.Post
	for i:=0; i < len(allPublicPosts); i++ {
		for j:=0; j < len(allPublicPosts[i].Tags); j++ {
			if (strings.Contains(strings.ToLower(allPublicPosts[i].Tags[j]), strings.ToLower(searchInput))){
				searchResult = append(searchResult, allPublicPosts[i])
			}
		}
	}
	return searchResult
}