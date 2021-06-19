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

	err1 := repository.addUser(session, newFollow.FollowerId)
	err2 := repository.addUser(session, newFollow.FollowedId)
	if err1 != nil || err2 != nil {
		return false
	}
	result, err3 := session.Run("match (u1:User), (u2:User) where u1.Id = $followerId and " +
		"u2.Id = $followedId merge (u1)-[f:follow{close: FALSE, muted: FALSE, blocked : FALSE, request: $isPrivate}]->(u2) return f",
		map[string]interface{}{"followerId":newFollow.FollowerId, "followedId":newFollow.FollowedId, "isPrivate":newFollow.IsPrivate })
	if err3 != nil {
		return false
	}
	if result.Next() {
		println(&result)
		return true
	}
	return false
}

func (repository *FollowRepository) addUser(session neo4j.Session, userId string) error {
	_, err := session.Run("merge (u:User{Id:$followerId}) return u", map[string]interface{}{"followerId":userId,})
	if err != nil {
		return err
	}
	return nil
}

func (repository *FollowRepository) SetFollowRequestFalse(loggedUserId string, followerId string) bool{
	session := *repository.DatabaseSession
	result, err := session.Run("match (u1:User{Id:$followerId})" +
		"-[f:follow {request: TRUE}]->(u2:User{Id:$loggedUserId}) set f.request = false return f;",
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

func (repository *FollowRepository) SetFollowMutedTrue(loggedUserId string, followingId string) bool{
	session := *repository.DatabaseSession
	result, err := session.Run("match (u1:User{Id:$loggedUserId})" +
		"-[f:follow {muted: FALSE}]->(u2:User{Id:$followingId}) set f.muted = true return f;",
		map[string]interface{}{"loggedUserId":loggedUserId, "followingId":followingId,})
	if err != nil {
		return false
	}
	if result.Next() {
		println(result)
		return true
	}
	return false
}

func (repository *FollowRepository) BlockUser(loggedUserId string, userId string) bool{
	session := *repository.DatabaseSession
	err1 := repository.addUser(session, loggedUserId)
	err2 := repository.addUser(session, userId)
	if err1 != nil || err2 != nil {
		return false
	}
	result, err := session.Run("match (u1:User), (u2:User) where u1.Id = $loggedUserId and " +
		"u2.Id = $userId merge (u1)-[f:follow{close: FALSE, muted: FALSE, blocked : TRUE, request: FALSE}]->(u2) return f",
		map[string]interface{}{"loggedUserId":loggedUserId, "userId":userId,})
	if err != nil {
		return false
	}
	if result.Next() {
		println(result)
		return true
	}
	return false
}

func (repository *FollowRepository) RemoveFollowing(loggedUserId string, followingId string) bool{
	session := *repository.DatabaseSession
	result, err := session.Run(" match (u1:User{Id:$loggedUserId})" +
		"-[f:follow]->(u2:User{Id: $followingId}) detach delete f return u1,u2",
		map[string]interface{}{"loggedUserId":loggedUserId, "followingId":followingId,})
	if err != nil {
		return false
	}
	if result.Next() {
		println(result)
		return true
	}
	return false
}

func (repository *FollowRepository) RemoveFollower(loggedUserId string, followerId string) bool{
	session := *repository.DatabaseSession
	result, err := session.Run(" match (u1:User{Id:$followerId})" +
		"-[f:follow]->(u2:User{Id:$loggedUserId}) detach delete f return u1,u2",
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

func (repository *FollowRepository) FindAllFollowersIds(userId string) ([]string, error){
	session := *repository.DatabaseSession
	var followersIds []string
	result, err := session.Run("match (u)" +
		"-[f:follow]->(u1:User{Id:$userId}) return u.Id",
		map[string]interface{}{"userId":userId,})
	if err != nil {
		return nil, err
	}
	for result.Next() {
		id, _ := result.Record().GetByIndex(0).(string)
		followersIds = append(followersIds, id)
	}
	if len(followersIds) == 0 {
		return nil, fmt.Errorf("no followers found")
	}
	return followersIds, nil
}

func (repository *FollowRepository) FindAllFollowingsIds(userId string) ([]string, error){
	session := *repository.DatabaseSession
	var followingsIds []string
	result, err := session.Run("match (u1:User{Id:$userId})" +
		"-[f:follow]->(u) return u.Id",
		map[string]interface{}{"userId":userId,})
	if err != nil {
		return nil, err
	}
	for result.Next() {
		id, _ := result.Record().GetByIndex(0).(string)
		followingsIds = append(followingsIds, id)
	}
	if len(followingsIds) == 0 {
		return nil, fmt.Errorf("no followings found")
	}
	return followingsIds, nil
}

func (repository *FollowRepository) FindAllBlockedUsersIds(userId string) ([]string, error){
	session := *repository.DatabaseSession
	var blockedUsersIds []string
	result, err := session.Run("match (u1:User{Id:$userId})" +
		"-[f:follow{blocked:TRUE}]->(u:User) return u.Id",
		map[string]interface{}{"userId":userId,})
	if err != nil {
		return nil, err
	}
	for result.Next() {
		id, _ := result.Record().GetByIndex(0).(string)
		blockedUsersIds = append(blockedUsersIds, id)
	}
	if len(blockedUsersIds) == 0 {
		return nil, fmt.Errorf("no blocked users found")
	}
	return blockedUsersIds, nil
}

func (repository *FollowRepository) UserAlreadyFollowed(followerId string, followedId string) error {
	session := *repository.DatabaseSession
	result, err := session.Run("match (u1:User{Id:$followerId})-[f:follow]->(u2:User{Id:$followedId}) return f;",
		map[string]interface{}{"followerId":followerId, "followedId":followedId,})
	if err != nil {
		return nil
	}
	if result.Next() {
		println(result)
 		return fmt.Errorf("user is already followed")
	}
	return nil
}