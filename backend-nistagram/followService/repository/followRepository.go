package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/dto"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowRepository struct {
	DatabaseSession *neo4j.Session
}

func (repository *FollowRepository) Hello (){
	fmt.Printf("Hello from Repository")
}

func (repository *FollowRepository) CreateFollowing(newFollow dto.NewFollowDTO) bool{
	session := *repository.DatabaseSession
	result, err := session.Run("merge (u1:User{Id:$followerId})" +
		"-[:follow{close: FALSE,muted: FALSE,request: $isPrivate}]-> (u2:User{Id:$followedId}) return u1, u2;",
		map[string]interface{}{"followerId":newFollow.FollowerId, "followedId":newFollow.FollowedId, "isPrivate":newFollow.IsPrivate })
	if err != nil {
		return false
	}
	if result.Next() {
		println(result)
		return true
	}
	return false
}

func (repository *FollowRepository) SetFollowRequestFalse(loggedUserId string, followerId string) bool{
	session := *repository.DatabaseSession
	result, err := session.Run(" match (u1:User{Id:$loggedUserId})" +
		"-[f:follow {request: TRUE, close: FALSE, muted: FALSE}]->(u2:User{Id: $followerId}) set f.request = false return f;",
		map[string]interface{}{"loggedUserId":loggedUserId, "followerId":followerId,})
	if err != nil {
		return false
	}
	if result.Next() {
		println(result)
		return true
	}
	return false
}

