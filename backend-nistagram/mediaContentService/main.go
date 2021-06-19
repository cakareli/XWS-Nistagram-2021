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

func handleFunc(handler *handler.PostHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/regular-user-posts/{username}", handler.GetAllRegularUserPosts).Methods("GET")
	router.HandleFunc("/public-posts", handler.GetAllPublicPosts).Methods("GET")
	router.HandleFunc("/new-post", handler.CreateNewPost).Methods("POST")
	router.HandleFunc("/comment-post", handler.CommentPost).Methods("PUT")
	router.HandleFunc("/search-location/{searchInput}",handler.GetLocationSearchResults).Methods("GET")
	router.HandleFunc("/search-user/{searchInput}",handler.GetUserSearchResults).Methods("GET")
	router.HandleFunc("/search-tag/{searchInput}",handler.GetTagSearchResults).Methods("GET")

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

	handleFunc(postHandler)
}