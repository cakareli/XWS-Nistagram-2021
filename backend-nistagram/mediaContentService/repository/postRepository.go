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

func (repository *PostRepository) FindPostById(postId primitive.ObjectID) (*model.Post, error){
	var post *model.Post
	result:= repository.Database.Collection("posts").FindOne(context.Background(), bson.M{"_id": postId})
	result.Decode(&post)

	if post == nil {
		return nil, fmt.Errorf("Post is NOT found!")
	}
	return post, nil
}

func (repository *PostRepository) Update(post *model.Post) error{
	postCollection := repository.Database.Collection("posts")

	updatedPost, err := postCollection.UpdateOne(context.TODO(), bson.M{"_id": post.Id},
		bson.D{
			{"$set", bson.D{{"likes", post.Likes}}},
			{"$set", bson.D{{"dislikes", post.Dislikes}}},
			{"$set", bson.D{{"comments", post.Comment}}},
			{"$set", bson.D{{"tags", post.Tags}}},
			{"$set", bson.D{{"location", post.Location}}},
			{"$set", bson.D{{"description", post.Description}}},
		})
	if err != nil {
		log.Fatal(err)
		return err
	} else if updatedPost.MatchedCount == 0 {
		return fmt.Errorf("Post does not exist!")
	}
	return nil
}