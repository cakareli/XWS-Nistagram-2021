package repository

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/model"
	"fmt"
	"gopkg.in/jmcvetta/neoism.v1"
)

type FollowRepository struct {
	DatabaseDriver *neoism.Database
}

func (repository *FollowRepository) Hello (){
	fmt.Printf("Hello from Repository")
}

func (repository *FollowRepository) Register(user *model.RegularUser){
	n, _ := repository.DatabaseDriver.CreateNode(neoism.Props{"regularUser": model.RegularUser{}})
	n.Relate("proba", n.Id(), neoism.Props{})

}