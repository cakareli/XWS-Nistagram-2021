package service

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/repository"
	"fmt"
)

type FollowService struct {
	FollowRepository *repository.FollowRepository
}

func (service *FollowService) Hello () {
	fmt.Printf("Hello from service!")
	service.FollowRepository.Hello()
}