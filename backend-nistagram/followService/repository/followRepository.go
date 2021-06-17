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

func (repository *FollowRepository) FollowUser(newFollow dto.NewFollowDTO) bool{
	session := *repository.DatabaseSession
	result, err := session.Run("merge (u1:User{Id:$follower})" +
		"-[:follow{close:false,muted:false,request:$isPrivate}]-> (u2:User{Id:$followed}) return u1, u2;",
		map[string]interface{}{"follower":newFollow.FollowerId, "followed":newFollow.FollowedId, "isPrivate":newFollow.IsPrivate })
	if err != nil {
		return false
	}
	if result.Next() {
		println(result)
		return true
	}
	return false
}