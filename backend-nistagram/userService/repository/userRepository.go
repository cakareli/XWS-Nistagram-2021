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
	fmt.Printf("Hello from Repository");
}

func (repository *UserRepository) CreateRegularUser(user *model.RegularUser) error{
	regularUserCollection := repository.Database.Collection("regularUsers")
		_, err := regularUserCollection.InsertOne(context.TODO(), bson.D{
		{Key: "regularUser", Value: user},
	})
	if err != nil {
		return fmt.Errorf("Regular user is NOT created")
	}
	return nil
}

func (repository *UserRepository) ExistByUsername(username string) bool{
	regularUserCollection := repository.Database.Collection("regularUsers")
	filterCursor, err := regularUserCollection.Find(context.TODO(), bson.M{"regularUser.username": username})
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
			{"$set", bson.D{{"regularUser.name", user.Name}}},
			{"$set", bson.D{{"regularUser.surname", user.Surname}}},
			{"$set", bson.D{{"regularUser.username", user.Username}}},
			{"$set", bson.D{{"regularUser.email", user.Email}}},
			{"$set", bson.D{{"regularUser.phoneNumber", user.PhoneNumber}}},
			{"$set", bson.D{{"regularUser.gender", user.Gender}}},
			{"$set", bson.D{{"regularUser.birthDate", user.BirthDate}}},
			{"$set", bson.D{{"regularUser.biography", user.Biography}}},
			{"$set", bson.D{{"regularUser.webSite", user.WebSite}}},
		})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (repository *UserRepository) UsernameChanged(username string, id primitive.ObjectID) bool{
	regularUserCollection := repository.Database.Collection("regularUsers")
	filterCursor, err := regularUserCollection.Find(context.TODO(), bson.M{"_id": id, "regularUser.username": username})
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