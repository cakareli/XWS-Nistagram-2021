package main

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/handler"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/repository"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
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

func initStoryRepository(database *mongo.Database) *repository.StoryRepository {
	return &repository.StoryRepository{Database: database}
}

func initStoryService(repository *repository.StoryRepository) *service.StoryService {
	return &service.StoryService{StoryRepository: repository}
}

func initStoryHandler(service *service.StoryService) *handler.StoryHandler {
	return &handler.StoryHandler{StoryService: service}
}

func handleFunc(handlerPost *handler.PostHandler, handlerStory *handler.StoryHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/regular-user-posts/{username}", handlerPost.GetAllRegularUserPosts).Methods("GET")
	router.HandleFunc("/public-posts", handlerPost.GetAllPublicPosts).Methods("GET")
	router.HandleFunc("/new-post", handlerPost.CreateNewPost).Methods("POST")
	router.HandleFunc("/new-story", handlerStory.CreateNewStory).Methods("POST")
	router.HandleFunc("/regular-user-stories/{username}", handlerStory.GetAllRegularUserStories).Methods("GET")
	router.HandleFunc("/comment-post", handlerPost.CommentPost).Methods("PUT")
	router.HandleFunc("/like-post", handlerPost.LikePost).Methods("PUT")
	router.HandleFunc("/dislike-post", handlerPost.DislikePost).Methods("PUT")
	router.HandleFunc("/search-location/{searchInput}",handlerPost.GetLocationSearchResults).Methods("GET")
	router.HandleFunc("/search-user/{searchInput}",handlerPost.GetUserSearchResults).Methods("GET")
	router.HandleFunc("/search-tag/{searchInput}",handlerPost.GetTagSearchResults).Methods("GET")

	c := SetupCors()

	http.Handle("/", c.Handler(router))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(router))
	if err != nil {
		log.Println(err)
	}
}

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
}

func initDatabase() *mongo.Database{
	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	mediaContentDatabase := client.Database("media-content")
	return mediaContentDatabase
}

func main() {
	mediaContentDatabase := initDatabase()

	postRepository := initPostRepository(mediaContentDatabase)
	postService := initPostService(postRepository)
	postHandler := initPostHandler(postService)

	storyRepository := initStoryRepository(mediaContentDatabase)
	storyService := initStoryService(storyRepository)
	storyHandler := initStoryHandler(storyService)

	handleFunc(postHandler, storyHandler)
}