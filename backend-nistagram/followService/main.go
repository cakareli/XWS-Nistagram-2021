package main

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/handler"
	"XWS-Nistagram-2021/backend-nistagram/followService/repository"
	"XWS-Nistagram-2021/backend-nistagram/followService/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/cors"
	"gopkg.in/jmcvetta/neoism.v1"
	"log"
	"net/http"
	"os"
)

func initFollowRepository(databaseDriver *neoism.Database) *repository.FollowRepository {
	return &repository.FollowRepository{DatabaseDriver: databaseDriver}
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

func initDatabase() *neo4j.Driver {
	var (
		driver neo4j.Driver
		err    error
	)
	for {
		driver, err = neo4j.NewDriver("bolt://"+os.Getenv("NEO4J_DBNAME")+":"+os.Getenv("NEO4J_PORT")+"/neo4j", neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASS"), "Neo4j"))

		if err != nil {
			fmt.Println("Cannot connect to database!")
		} else {
			fmt.Println(" Successfully connected to the database!")
			break
		}
	}

	return &driver
}

func main() {
	//_ = initDatabase()
	db, err := neoism.Connect("http://neo4j:7474/db/data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("kurcinela")
	authenticationRepository := initFollowRepository(db)
	authenticationService := initFollowService(authenticationRepository)
	authenticationHandler := initFollowHandler(authenticationService)

	handleFunc(authenticationHandler)
}
