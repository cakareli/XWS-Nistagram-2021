package main

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/handler"
	"XWS-Nistagram-2021/backend-nistagram/followService/repository"
	"XWS-Nistagram-2021/backend-nistagram/followService/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func initFollowRepository(databaseSession *neo4j.Session) *repository.FollowRepository {
	return &repository.FollowRepository{DatabaseSession: databaseSession}
}

func initFollowService(repository *repository.FollowRepository) *service.FollowService {
	return &service.FollowService{FollowRepository: repository}
}

func initFollowHandler(service *service.FollowService) *handler.FollowHandler {
	return &handler.FollowHandler{FollowService: service}
}

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
}

func handleFunc(handler *handler.FollowHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", handler.Hello).Methods("GET")

	c := SetupCors()

	fmt.Println("Server is running...")

	http.Handle("/", c.Handler(router))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(router))
	if err != nil {
		log.Println(err)
	}
}

func initDatabase() (neo4j.Session, error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	if driver, err = neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345", "")); err != nil {
		return nil, err
	}

	if session = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}); err != nil {
		return nil, err
	}

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("match (u) return u;", map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], err
		}
		return nil, result.Err()
	})

	if err != nil {
		return nil, err
	}

	return session, nil
}

func main() {
	session, err := initDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	println("kurac")
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		userId := "1234567"
		result, err := session.Run("merge (u:User{Id:$userId}) return u;", map[string]interface{}{"userId":userId,})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], err
		}
		println(result)
		return nil, result.Err()
	})

	authenticationRepository := initFollowRepository(&session)
	authenticationService := initFollowService(authenticationRepository)
	authenticationHandler := initFollowHandler(authenticationService)

	handleFunc(authenticationHandler)
}
