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

	err := service.FollowRepository.UserAlreadyFollowed(newFollow.FollowerId, newFollow.FollowedId)
	if err != nil {
		return false
	}
	userIsFollowed := service.FollowRepository.CreateFollowing(newFollow)
	return userIsFollowed
}

func (service *FollowService) AcceptFollowRequest(loggedUserId string, followerId string) bool {
	fmt.Println("accepting follow request...")

	userIsAccepted := service.FollowRepository.SetFollowRequestFalse(loggedUserId, followerId)
	return userIsAccepted
}

func (service *FollowService) RemoveFollower(loggedUserId string, followerId string) bool {
	fmt.Println("removing follower...")

	userIsRemoved := service.FollowRepository.RemoveFollowing(loggedUserId, followerId)
	return userIsRemoved
}

