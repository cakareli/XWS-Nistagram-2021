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

	router.HandleFunc("/follow", handler.FollowUser).Methods("POST")
	router.HandleFunc("/accept-follow/{loggedUserId}/{followerId}", handler.AcceptFollowRequest).Methods("PUT")
	router.HandleFunc("/mute-following/{loggedUserId}/{followingId}", handler.MuteFollowing).Methods("PUT")
	router.HandleFunc("/block-user/{loggedUserId}/{userId}", handler.BlockUser).Methods("POST")
	router.HandleFunc("/unblock-user/{loggedUserId}/{userId}", handler.UnblockUser).Methods("POST")
	router.HandleFunc("/remove-following/{loggedUserId}/{followingId}", handler.RemoveFollowing).Methods("POST")
	router.HandleFunc("/remove-follower/{loggedUserId}/{followerId}", handler.RemoveFollower).Methods("POST")
	router.HandleFunc("/followers/{loggedUserId}", handler.FindAllUserFollowers).Methods("GET")
	router.HandleFunc("/followings/{loggedUserId}", handler.FindAllUserFollowings).Methods("GET")
	router.HandleFunc("/blocked-users/{loggedUserId}", handler.FindAllUserBlockedUsers).Methods("GET")
	router.HandleFunc("/muted-users/{loggedUserId}", handler.FindAllUserMutedUsers).Methods("GET")
	router.HandleFunc("/follow-requests/{loggedUserId}", handler.FindAllUserFollowRequests).Methods("GET")

	c := SetupCors()

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
	if driver, err = neo4j.NewDriver("neo4j://neo4j:7687", neo4j.BasicAuth("neo4j", "12345", "")); err != nil {
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

	authenticationRepository := initFollowRepository(&session)
	authenticationService := initFollowService(authenticationRepository)
	authenticationHandler := initFollowHandler(authenticationService)

	handleFunc(authenticationHandler)
}
