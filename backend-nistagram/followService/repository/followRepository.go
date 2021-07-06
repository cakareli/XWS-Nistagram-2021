package repository

import (

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowRepository struct {
	DatabaseSession *neo4j.Session
}
