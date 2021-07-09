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

func initPostService(repository1 *repository.PostRepository, repository2 *repository.NotificationRepository) *service.PostService {
	return &service.PostService{PostRepository: repository1, NotificationRepository: repository2}
}

func initPostHandler(service *service.PostService) *handler.PostHandler {
	return &handler.PostHandler{PostService: service}
}

func initStoryRepository(database *mongo.Database) *repository.StoryRepository {
	return &repository.StoryRepository{Database: database}
}

func initStoryService(repository1 *repository.StoryRepository, repository2 *repository.NotificationRepository) *service.StoryService {
	return &service.StoryService{StoryRepository: repository1, NotificationRepository: repository2}
}

func initStoryHandler(service *service.StoryService) *handler.StoryHandler {
	return &handler.StoryHandler{StoryService: service}
}

func initNotificationRepository(database *mongo.Database) *repository.NotificationRepository {
	return &repository.NotificationRepository{Database: database}
}

func initNotificationService(repository *repository.NotificationRepository) *service.NotificationService {
	return &service.NotificationService{NotificationRepository: repository}
}

func initNotificationHandler(service *service.NotificationService) *handler.NotificationHandler {
	return &handler.NotificationHandler{NotificationService: service}
}

func handleFunc(handlerPost *handler.PostHandler, handlerStory *handler.StoryHandler, handlerNotification *handler.NotificationHandler) {
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
	router.HandleFunc("/update-posts-privacy", handlerPost.UpdatePostsPrivacy).Methods("POST")
	router.HandleFunc("/update-stories-privacy", handlerStory.UpdateStoriesPrivacy).Methods("POST")
	router.HandleFunc("/liked-posts/{username}", handlerPost.GetAllLikedPostsByUsername).Methods("GET")
	router.HandleFunc("/disliked-posts/{username}", handlerPost.GetAllDislikedPostsByUsername).Methods("GET")
	router.HandleFunc("/saved-posts/{username}", handlerPost.GetAllSavedPostsByUsername).Methods("GET")
	router.HandleFunc("/new-notification", handlerNotification.CreateNewNotification).Methods("POST")
	router.HandleFunc("/notifications/{userId}", handlerNotification.GetAllNotificationsByUserId).Methods("GET")
	router.HandleFunc("/feed-posts", handlerPost.GetUsersFeed).Methods("POST")
	router.HandleFunc("/report-post", handlerPost.ReportPost).Methods("POST")
	router.HandleFunc("/report-story", handlerStory.ReportStory).Methods("POST")

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
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
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

	notificationRepository := initNotificationRepository(mediaContentDatabase)
	notificationService := initNotificationService(notificationRepository)
	notificationHandler := initNotificationHandler(notificationService)

	postRepository := initPostRepository(mediaContentDatabase)
	postService := initPostService(postRepository, notificationRepository)
	postHandler := initPostHandler(postService)

	storyRepository := initStoryRepository(mediaContentDatabase)
	storyService := initStoryService(storyRepository, notificationRepository)
	storyHandler := initStoryHandler(storyService)

	handleFunc(postHandler, storyHandler, notificationHandler)
}