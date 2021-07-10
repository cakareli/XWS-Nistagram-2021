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

func (service *FollowService) UnmuteFollowing(loggedUserId string, followingId string) bool {
	fmt.Println("unmuting user...")

	userIsUnmuted := service.FollowRepository.SetFollowMutedFalse(loggedUserId, followingId)
	return userIsUnmuted
}

func (service *FollowService) TurnNotificationsForUserOn(loggedUserId string, followingId string) bool {
	fmt.Println("turning notifications for user on...")

	userNotificationsTurnedOn := service.FollowRepository.SetFollowNotificationsTrue(loggedUserId, followingId)
	return userNotificationsTurnedOn
}

func (service *FollowService) TurnNotificationsForUserOff(loggedUserId string, followingId string) bool {
	fmt.Println("turning notifications for user off...")

	userNotificationsTurnedOff := service.FollowRepository.SetFollowNotificationsFalse(loggedUserId, followingId)
	return userNotificationsTurnedOff
}

func (service *FollowService) AddToCloseFollowers(loggedUserId string, followingId string) bool {
	fmt.Println("adding user to close followers...")

	userIsInClose := service.FollowRepository.SetFollowCloseTrue(loggedUserId, followingId)
	return userIsInClose
}

func (service *FollowService) RemoveFromCloseFollowers(loggedUserId string, followingId string) bool {
	fmt.Println("removing user from close followers...")

	userIsRemoved := service.FollowRepository.SetFollowCloseFalse(loggedUserId, followingId)
	return userIsRemoved
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

func (service *FollowService) FindAllUserFollowers(loggedUserId string) ([]dto.UserDTO, error) {
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

func (service *FollowService) FindAllUserFollowings(loggedUserId string) ([]dto.UserDTO, error) {
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

func (service *FollowService) DeleteUser(id string) error{
	deleted := service.FollowRepository.DeleteUser(id)
	if !deleted {
		return nil
	}
	return nil
}

func (service *FollowService) FindAllPostsForFeed(loggedUserId string) ([]dto.PostDTO, error) {
	fmt.Println("getting all feed users...")

	feedUsers, err := service.FollowRepository.FindAllFeedUsersIds(loggedUserId)
	if err != nil {
		return nil, err
	}
	userDTOs, err2 := service.getUserDTOsFromUserIds(feedUsers)
	if err2 != nil {
		return nil, err2
	}
	var userIds []string
	userIds = appendArray(userIds, userDTOs)
	userIds = append(userIds, loggedUserId)
	postDTOs, err3 := service.getPostsFromFeedUserIds(userIds)
	if err3 != nil {
		return nil, err3
	}
	return postDTOs, nil
}

func (service *FollowService) FindAllUserBlockedUsers(loggedUserId string) ([]dto.UserDTO, error) {
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

func (service *FollowService) FindAllUserMutedUsers(loggedUserId string) ([]dto.UserDTO, error) {
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

func (service *FollowService) FindAllUserFollowRequests(loggedUserId string) ([]dto.UserDTO, error) {
	fmt.Println("getting all follow requests users...")

	followRequestsIds, err := service.FollowRepository.FindAllUserFollowRequestsIds(loggedUserId)
	if err != nil {
		return nil, err
	}
	userDTOs, err2 := service.getUserDTOsFromUserIds(followRequestsIds)
	if err2 != nil {
		return nil, err2
	}
	return userDTOs, nil
}

func (service *FollowService) FindAllUserCloseFollowers(loggedUserId string) ([]dto.UserDTO, error) {
	fmt.Println("getting all close followers...")

	closeFollowers, err := service.FollowRepository.FindAllUserCloseFollowers(loggedUserId)
	if err != nil {
		return nil, err
	}
	userDTOs, err2 := service.getUserDTOsFromUserIds(closeFollowers)
	if err2 != nil {
		return nil, err2
	}
	return userDTOs, nil
}

func (service *FollowService) FindAllFollowersWithNotificationsTurnedOn(loggedUserId string) ([]string, error) {
	fmt.Println("getting all followers with notifications turned on...")

	followersIds, err := service.FollowRepository.FindAllFollowersWithNotificationsTurnedOn(loggedUserId)
	if err != nil {
		return nil, err
	}
	return followersIds, nil
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

func (service *FollowService) getPostsFromFeedUserIds(userIds []string) ([]dto.PostDTO, error){
	var postDTOs []dto.PostDTO
	postBody, _ := json.Marshal(userIds)
	requestUrl := fmt.Sprintf("http://%s:%s/feed-posts", os.Getenv("MEDIA_CONTENT_SERVICE_DOMAIN"), os.Getenv("MEDIA_CONTENT_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&postDTOs)
	return postDTOs, nil
}

func (service *FollowService) getStoriesFromFeedUserIds(userIds []string) ([]dto.StoryDTO, error){
	var storyDTOs []dto.StoryDTO
	postBody, _ := json.Marshal(userIds)
	requestUrl := fmt.Sprintf("http://%s:%s/feed-stories", os.Getenv("MEDIA_CONTENT_SERVICE_DOMAIN"), os.Getenv("MEDIA_CONTENT_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&storyDTOs)
	return storyDTOs, nil
}

func appendArray(firstSlice []string, secondSlice []dto.UserDTO) []string{
	for i := 0; i < len(secondSlice); i++{
		firstSlice = append(firstSlice,secondSlice[i].UserId)
	}
	return firstSlice
}