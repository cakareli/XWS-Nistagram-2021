package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

type RegularUserRepository struct {
	Database *mongo.Database
}

func (repository *RegularUserRepository) Register(user *model.RegularUser) (string, error){
	regularUserCollection := repository.Database.Collection("regularUsers")
	res, err := regularUserCollection.InsertOne(context.TODO(), &user)
	if err != nil {
		return "", fmt.Errorf("regular user is NOT created")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repository *RegularUserRepository) ExistByUsername(username string) bool{
	regularUserCollection := repository.Database.Collection("regularUsers")
	filterCursor, err := regularUserCollection.Find(context.TODO(), bson.M{"username": username})
	if err != nil {
		log.Fatal(err)
	}

	var userFiltered bson.D
	if err = filterCursor.All(context.TODO(), &userFiltered); err != nil {
		log.Fatal(err)
	}
	if len(userFiltered) != 0 {
		return true
	}
	return false
}

func (repository *RegularUserRepository) Update(user *model.RegularUser) error{
	regularUserCollection := repository.Database.Collection("regularUsers")

	updatedRegularUser, err := regularUserCollection.UpdateOne(context.TODO(), bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"name", user.Name}}},
			{"$set", bson.D{{"surname", user.Surname}}},
			{"$set", bson.D{{"username", user.Username}}},
			{"$set", bson.D{{"email", user.Email}}},
			{"$set", bson.D{{"phoneNumber", user.PhoneNumber}}},
			{"$set", bson.D{{"gender", user.Gender}}},
			{"$set", bson.D{{"birthDate", user.BirthDate}}},
			{"$set", bson.D{{"biography", user.Biography}}},
			{"$set", bson.D{{"webSite", user.WebSite}}},
			{"$set", bson.D{{"likedPosts", user.LikedPosts}}},
			{"$set", bson.D{{"dislikedPosts", user.DislikedPosts}}},
		})
	if err != nil {
		log.Fatal(err)
		return err
	} else if updatedRegularUser.MatchedCount == 0 {
		return fmt.Errorf("user does not exist")
	}
	return nil
}

func (repository *RegularUserRepository) UsernameChanged(username string, id primitive.ObjectID) bool{
	regularUserCollection := repository.Database.Collection("regularUsers")
	filterCursor, err := regularUserCollection.Find(context.TODO(), bson.M{"_id": id, "username": username})
	if err != nil {
		log.Fatal(err)
	}

	var userFiltered bson.D
	if err = filterCursor.All(context.TODO(), &userFiltered); err != nil {
		log.Fatal(err)
	}
	if len(userFiltered) == 0 {
		return true
	}
	return false
}

func (repository *RegularUserRepository) FindUserById(userId primitive.ObjectID) (*model.RegularUser, error){
	var regularUser *model.RegularUser
	result:= repository.Database.Collection("regularUsers").FindOne(context.Background(), bson.M{"_id": userId})
	result.Decode(&regularUser)

	if regularUser == nil {
		return nil, fmt.Errorf("regular user is NOT found")
	}
	return regularUser, nil
}

func (repository *RegularUserRepository) FindUserByUsername(username string) (*model.RegularUser, error){
	var regularUser *model.RegularUser
	result:= repository.Database.Collection("regularUsers").FindOne(context.Background(), bson.M{"username": username})
	result.Decode(&regularUser)

	if regularUser == nil {
		return nil, fmt.Errorf("regular user is NOT found")
	}
	return regularUser, nil
}

func (repository *RegularUserRepository) GetAllPublicRegularUsers() ([]bson.D, error){

	usersCollection := repository.Database.Collection("regularUsers")
	filterCursor, err := usersCollection.Find(context.TODO(), bson.M{"profilePrivacy.privacyType": 0})
	if err != nil {
		log.Fatal(err)
	}

	var postsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &postsFiltered); err != nil {
		log.Fatal(err)
	}
	return postsFiltered, nil
}

func (repository *RegularUserRepository) GetUserSearchResults(searchInput string, allPublicRegularUsers []model.RegularUser) []model.RegularUser{
	var searchResults []model.RegularUser
	for i := 0; i<len(allPublicRegularUsers); i++{
		if(strings.Contains(strings.ToLower(allPublicRegularUsers[i].Username), strings.ToLower(searchInput))){
			searchResults = append(searchResults, allPublicRegularUsers[i])
		}
	}
	return searchResults
}