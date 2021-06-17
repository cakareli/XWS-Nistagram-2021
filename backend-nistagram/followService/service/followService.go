package service

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/dto"
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

func (service *FollowService) FollowUser(newFollow dto.NewFollowDTO) bool {
	fmt.Println("following user...")

	userIsFollowed := service.FollowRepository.FollowUser(newFollow)
	return userIsFollowed
}