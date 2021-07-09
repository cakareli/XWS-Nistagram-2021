package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

type PostRepository struct {
	Database *mongo.Database
}

func (repository *PostRepository) Create(post *model.Post) error {
	postsCollection := repository.Database.Collection("posts")
	_, err := postsCollection.InsertOne(context.TODO(), &post)
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

func (repository *PostRepository) FindAllPostsByUserId(id string) []bson.D{
	postsCollection := repository.Database.Collection("posts")
	filterCursor, err := postsCollection.Find(context.TODO(), bson.M{"regularUser._id": id})
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

func (repository *PostRepository) Update(post *model.Post) error {
	postCollection := repository.Database.Collection("posts")

	updatedPost, err := postCollection.UpdateOne(context.TODO(), bson.M{"_id": post.Id},
		bson.D{
			{"$set", bson.D{{"likes", post.Likes}}},
			{"$set", bson.D{{"dislikes", post.Dislikes}}},
			{"$set", bson.D{{"comments", post.Comment}}},
			{"$set", bson.D{{"tags", post.Tags}}},
			{"$set", bson.D{{"hashtags", post.Hashtags}}},
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

func (repository *PostRepository) UpdatePostPrivacy(privacyType model.PrivacyType, postId primitive.ObjectID) error {
	postCollection := repository.Database.Collection("posts")

	updatedPost, err := postCollection.UpdateOne(context.TODO(), bson.M{"_id": postId},
		bson.D{
			{"$set", bson.D{{"regularUser.privacyType", privacyType}}},
		})
	if err != nil {
		log.Fatal(err)
		return err
	} else if updatedPost.MatchedCount == 0 {
		return fmt.Errorf("Post does not exist!")
	}
	return nil
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
		for j:=0; j < len(allPublicPosts[i].Hashtags); j++ {
			if (strings.Contains(strings.ToLower(allPublicPosts[i].Hashtags[j]), strings.ToLower(searchInput))){
				searchResult = append(searchResult, allPublicPosts[i])
			}
		}
	}
	return searchResult

}

func (repository *PostRepository) GetAllPostsByIds(ids []string) []bson.D{
	oids := make([]primitive.ObjectID, len(ids))
	for i := range ids {
		objID, err := primitive.ObjectIDFromHex(ids[i])
		if err == nil {
			oids = append(oids, objID)
		}
	}
	postsCollection := repository.Database.Collection("posts")
	filterCursor, err := postsCollection.Find(context.TODO(), bson.M{"_id": bson.M{"$in": oids}})
	if err != nil {
		log.Fatal(err)
	}

	var postsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &postsFiltered); err != nil {
		log.Fatal(err)
	}
	return postsFiltered
}

func (repository *PostRepository) CreatePostReport(inappropriateContentPost *model.InappropriateContentPost) error {
	postReportsCollection := repository.Database.Collection("postReports")
	_, err := postReportsCollection.InsertOne(context.TODO(), &inappropriateContentPost)
	if err != nil {
		return fmt.Errorf("Post report is NOT created")
	}
	return nil
}

func (repository *PostRepository) GetAllPostReports() []bson.D{
	reportsCollection := repository.Database.Collection("postReports")
	filterCursor, err := reportsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var postsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &postsFiltered); err != nil {
		log.Fatal(err)
	}
	return postsFiltered
}

func (repository *PostRepository) DeleteReportedPost(id primitive.ObjectID) error{

	reportedPostsCollection := repository.Database.Collection("postReports")
	_, err := reportedPostsCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (repository *PostRepository) DeletePost(id primitive.ObjectID) error{

	postsCollection := repository.Database.Collection("posts")
	_, err := postsCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

