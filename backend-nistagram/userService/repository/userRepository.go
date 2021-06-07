package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepository struct {
	Database *mongo.Database
}

func (repository *UserRepository) Hello (){
	fmt.Printf("Hello from Repository")
}

func (repository *UserRepository) CreateRegularUser(user *model.RegularUser) error{
	regularUserCollection := repository.Database.Collection("regularUsers")
		_, err := regularUserCollection.InsertOne(context.TODO(), &user)
	if err != nil {
		return fmt.Errorf("regular user is NOT created")
	}
	return nil
}

func (repository *UserRepository) ExistByUsername(username string) bool{
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

func (repository *UserRepository) UpdateRegularUser(user *model.RegularUser) error{
	regularUserCollection := repository.Database.Collection("regularUsers")

	_, err := regularUserCollection.UpdateOne(context.TODO(), bson.M{"_id": user.Id},
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
		})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (repository *UserRepository) UsernameChanged(username string, id primitive.ObjectID) bool{
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