package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/model"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowRepository struct {
	DatabaseSession *neo4j.Session
}

func (repository *FollowRepository) Hello (){
	fmt.Printf("Hello from Repository")
}

func (repository *FollowRepository) Register(user *model.RegularUser){










/*	n, _ := repository.DatabaseDriver.CreateNode(neoism.Props{"regularUser": model.RegularUser{}})
	n.Relate("proba", n.Id(), neoism.Props{})*/

}