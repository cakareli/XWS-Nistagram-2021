package repository

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowRepository struct {
	DatabaseDriver *neo4j.Driver
}

func (repository *FollowRepository) Hello (){
	fmt.Printf("Hello from Repository")
}
