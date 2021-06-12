package main

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/handler"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/repository"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func initPostRepository(database *mongo.Database) *repository.PostRepository {
	return &repository.PostRepository{Database: database}
}

func initPostService(repository *repository.PostRepository) *service.PostService {
	return &service.PostService{PostRepository: repository}
}

func initPostHandler(service *service.PostService) *handler.PostHandler {
	return &handler.PostHandler{PostService: service}
}

func handleFunc(handler *handler.PostHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", handler.Hello).Methods("GET")
	router.HandleFunc("/regular-user-posts/{username}", handler.GetAllRegularUserPosts).Methods("GET")
	router.HandleFunc("/public-posts", handler.GetAllPublicPosts).Methods("GET")

	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8081"), router))
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	mediaContentDatabase := client.Database("media-content")

	postRepository := initPostRepository(mediaContentDatabase)
	postService := initPostService(postRepository)
	postHandler := initPostHandler(postService)
	handleFunc(postHandler)
}