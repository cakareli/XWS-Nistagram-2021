package service

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/dto"
	"XWS-Nistagram-2021/backend-nistagram/followService/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	userIsFollowed := service.FollowRepository.AddFollowing(newFollow)
	return userIsFollowed
}

func (service *FollowService) AcceptFollowRequest(loggedUserId string, followerId string) bool {
	fmt.Println("accepting follow request...")

	userIsAccepted := service.FollowRepository.SetFollowRequestFalse(loggedUserId, followerId)
	return userIsAccepted
}

func (service *FollowService) MuteFollowing(loggedUserId string, followingId string) bool {
	fmt.Println("muting user...")

	userIsMuted := service.FollowRepository.SetFollowMutedTrue(loggedUserId, followingId)
	return userIsMuted
}

func (service *FollowService) BlockUser(loggedUserId string, userId string) bool {
	fmt.Println("blocking user...")

	err1 := service.FollowRepository.UserAlreadyFollowed(loggedUserId, userId)
	if err1 != nil {
		service.RemoveFollowing(loggedUserId, userId)
	}
	err2 := service.FollowRepository.UserAlreadyFollowed(userId, loggedUserId)
	if err2 != nil {
		service.RemoveFollowing(userId, loggedUserId)
	}
	userIsBlocked := service.FollowRepository.BlockUser(loggedUserId, userId)
	return userIsBlocked
}

func (service *FollowService) UnblockUser(loggedUserId string, userId string) bool {
	fmt.Println("unblocking user...")

	userIsUnblocked := service.FollowRepository.RemoveFollowing(loggedUserId, userId)
	return userIsUnblocked
}

func (service *FollowService) RemoveFollowing(loggedUserId string, followingId string) bool {
	fmt.Println("removing following...")

	FollowIsRemoved := service.FollowRepository.RemoveFollowing(loggedUserId, followingId)
	return FollowIsRemoved
}

func (service *FollowService) RemoveFollower(loggedUserId string, followerId string) bool {
	fmt.Println("removing follower...")

	FollowIsRemoved := service.FollowRepository.RemoveFollower(loggedUserId, followerId)
	return FollowIsRemoved
}

func (service *FollowService) FindAllFollowers(loggedUserId string) ([]dto.UserDTO, error) {
	fmt.Println("getting all followers...")

	followersIds, err := service.FollowRepository.FindAllUserFollowersIds(loggedUserId)
	if err != nil {
		return nil, err
	}
	userDTOs, err2 := service.getUserDTOsFromUserIds(followersIds)
	if err2 != nil {
		return nil, err2
	}
	return userDTOs, nil
}

func (service *FollowService) FindAllFollowings(loggedUserId string) ([]dto.UserDTO, error) {
	fmt.Println("getting all followings...")

	followingsIds, err := service.FollowRepository.FindAllUserFollowingsIds(loggedUserId)
	if err != nil {
		return nil, err
	}
	userDTOs, err2 := service.getUserDTOsFromUserIds(followingsIds)
	if err2 != nil {
		return nil, err2
	}
	return userDTOs, nil
}

func (service *FollowService) FindAllBlockedUsers(loggedUserId string) ([]dto.UserDTO, error) {
	fmt.Println("getting all blocked users...")

	blockedUsersIds, err := service.FollowRepository.FindAllUserBlockedUsersIds(loggedUserId)
	if err != nil {
		return nil, err
	}
	userDTOs, err2 := service.getUserDTOsFromUserIds(blockedUsersIds)
	if err2 != nil {
		return nil, err2
	}
	return userDTOs, nil
}

func (service *FollowService) FindAllMutedUsers(loggedUserId string) ([]dto.UserDTO, error) {
	fmt.Println("getting all muted users...")

	mutedUsersIds, err := service.FollowRepository.FindAllUserMutedUsersIds(loggedUserId)
	if err != nil {
		return nil, err
	}
	userDTOs, err2 := service.getUserDTOsFromUserIds(mutedUsersIds)
	if err2 != nil {
		return nil, err2
	}
	return userDTOs, nil
}

func (service *FollowService) getUserDTOsFromUserIds(userIds []string) ([]dto.UserDTO, error) {
	var userDTOs []dto.UserDTO
	postBody, _ := json.Marshal(userIds)
	requestUrl := fmt.Sprintf("http://%s:%s/by-users-ids", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&userDTOs)
	return userDTOs, nil
}