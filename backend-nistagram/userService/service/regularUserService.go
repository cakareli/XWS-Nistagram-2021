package service

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"XWS-Nistagram-2021/backend-nistagram/userService/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"go.mongodb.org/mongo-driver/bson"

)

type RegularUserService struct {
	RegularUserRepository *repository.RegularUserRepository
}

func (service *RegularUserService) Register(regularUserRegistrationDto dto.RegularUserRegistrationDTO) error {
	fmt.Println("Creating regular user")

	if service.RegularUserRepository.ExistByUsername(regularUserRegistrationDto.Username) {
		return fmt.Errorf("username is already taken")
	}

	var regularUser = createRegularUserFromRegularUserRegistrationDTO(&regularUserRegistrationDto)
	createdUserId, err := service.RegularUserRepository.Register(regularUser)
	if err != nil {
		return err
	}
	postBody, _ := json.Marshal(map[string]string{
		"userId": createdUserId,
		"email": regularUserRegistrationDto.Email,
		"password": regularUserRegistrationDto.Password,
		"username": regularUserRegistrationDto.Username,
		"name": regularUserRegistrationDto.Name,
		"surname": regularUserRegistrationDto.Surname,
	})
	requestUrl := fmt.Sprintf("http://%s:%s/register", os.Getenv("AUTHENTICATION_SERVICE_DOMAIN"), os.Getenv("AUTHENTICATION_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *RegularUserService) UpdatePersonalInformations(regularUserUpdateDto dto.RegularUserUpdateDTO) error {
	fmt.Println("Updating regular user")

	if service.RegularUserRepository.ExistByUsername(regularUserUpdateDto.Username) {
		id, _ := primitive.ObjectIDFromHex(regularUserUpdateDto.Id)
		if service.RegularUserRepository.UsernameChanged(regularUserUpdateDto.Username, id) {
			return fmt.Errorf("username is already taken")
		}
	}

	var regularUser = createRegularUserFromRegularUserUpdateDTO(&regularUserUpdateDto)
	err := service.RegularUserRepository.UpdatePersonalInformations(regularUser)
	if err != nil {
		return err
	}

	postBody, _ := json.Marshal(map[string]string{
		"_id": regularUserUpdateDto.Id,
		"email": regularUserUpdateDto.Email,
		"username": regularUserUpdateDto.Username,
		"name": regularUserUpdateDto.Name,
		"surname": regularUserUpdateDto.Surname,
	})
	requestUrl := fmt.Sprintf("http://%s:%s/update", os.Getenv("AUTHENTICATION_SERVICE_DOMAIN"), os.Getenv("AUTHENTICATION_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *RegularUserService) FindUserById(userId primitive.ObjectID) (*model.RegularUser, error){
	fmt.Print("Searching for logged user...")
	regularUser, err := service.RegularUserRepository.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	return regularUser, err
}

func (service *RegularUserService) CreateRegularUserPostDTOByUsername(username string) (*dto.RegularUserPostDTO, error){
	regularUser, err := service.RegularUserRepository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	regularUserPostDto := createRegularUserPostDTOFromRegularUser(regularUser)
	return regularUserPostDto, nil
}

func (service *RegularUserService) FindRegularUserByUsername(username string) (*dto.RegularUserProfileDataDTO, error){
	regularUser, err := service.RegularUserRepository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	regularUserPostDto := createRegularUserProfileDataDto(regularUser)
	return regularUserPostDto, nil
}

func (service *RegularUserService) GetUserSearchResults(searchInput string) ([]model.RegularUser, error){
	searchPublicRegularUser,err := service.RegularUserRepository.GetAllPublicRegularUsers()
	if err != nil {
		return nil, err
	}
	searchPublicRegularUserModel := CreateUserFromDocuments(searchPublicRegularUser)
	searchPublicRegularUserResults := service.RegularUserRepository.GetUserSearchResults(searchInput, searchPublicRegularUserModel)

	return searchPublicRegularUserResults, nil
}

func (service *RegularUserService) GetAllPublicRegularUsers() ([]dto.RegularUserDTO, error){
	allRegularUsers,err := service.RegularUserRepository.GetAllPublicRegularUsers()
	if err != nil {
		return nil, err
	}

	allRegularUsersModel := CreateUserFromDocuments(allRegularUsers)

	allRegularUsersDto := createRegularUserDtoFromRegularUser(allRegularUsersModel)
	return allRegularUsersDto,nil
}

func CreateUserFromDocuments(UserDocuments []bson.D) []model.RegularUser {
	var users []model.RegularUser
	for i := 0; i < len(UserDocuments); i++ {
		var user model.RegularUser
		bsonBytes, _ := bson.Marshal(UserDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &user)
		users = append(users, user)
	}
	return users
}
func (service *RegularUserService) FindUsersByIds(usersIds []string) (*[]dto.UserFollowDTO, error){
	var users []model.RegularUser
	for i:=0; i < len(usersIds); i++ {
		id, _ := primitive.ObjectIDFromHex(usersIds[i])
		regularUser, err := service.RegularUserRepository.FindUserById(id)
		if err != nil {
			return nil, err
		}
		users = append(users, *regularUser)
	}

	userFollowDTOs := createUserFollowDTOsFromRegularUsers(users)
	return userFollowDTOs, nil
}

func createRegularUserPostDTOFromRegularUser(regularUser *model.RegularUser) *dto.RegularUserPostDTO {
	var regularUserPostDto dto.RegularUserPostDTO
	regularUserPostDto.Id = regularUser.Id.Hex()
	regularUserPostDto.PrivacyType = &regularUser.ProfilePrivacy.PrivacyType

	return &regularUserPostDto
}

func createRegularUserFromRegularUserRegistrationDTO(regularUserDto *dto.RegularUserRegistrationDTO) *model.RegularUser{
	profilePrivacy := model.ProfilePrivacy{
		PrivacyType: model.PrivacyType(0),
		AllMessageRequests: true,
		TagsAllowed: true,
	}
	var regularUser model.RegularUser
	regularUser.Name = regularUserDto.Name
	regularUser.Surname = regularUserDto.Surname
	regularUser.Username = regularUserDto.Username
	regularUser.Password = regularUserDto.Password
	regularUser.Email = regularUserDto.Email
	regularUser.PhoneNumber = regularUserDto.PhoneNumber
	regularUser.BirthDate = regularUserDto.BirthDate
	regularUser.Biography = regularUserDto.Biography
	regularUser.WebSite = regularUserDto.WebSite
	regularUser.ProfilePrivacy = profilePrivacy
	regularUser.IsDisabled = false
	regularUser.UserRole = model.UserRole(0)
	regularUser.UserType = model.UserType(0)
	regularUser.Gender = regularUserDto.Gender

	return &regularUser
}

func createRegularUserFromRegularUserUpdateDTO(userUpdateDto *dto.RegularUserUpdateDTO) *model.RegularUser{
	id, _ := primitive.ObjectIDFromHex(userUpdateDto.Id)
	var regularUser model.RegularUser
	regularUser.Id = id
	regularUser.Name = userUpdateDto.Name
	regularUser.Surname = userUpdateDto.Surname
	regularUser.Username = userUpdateDto.Username
	regularUser.Email = userUpdateDto.Email
	regularUser.PhoneNumber = userUpdateDto.PhoneNumber
	regularUser.Gender= userUpdateDto.Gender
	regularUser.BirthDate = userUpdateDto.BirthDate
	regularUser.Biography = userUpdateDto.Biography
	regularUser.WebSite = userUpdateDto.WebSite

	return &regularUser
}

func createRegularUserDtoFromRegularUser(allRegularUsers []model.RegularUser) []dto.RegularUserDTO{

	var regularUser []dto.RegularUserDTO
	for i := 0; i < len(allRegularUsers); i++{
		var regularUserIteration dto.RegularUserDTO
		regularUserIteration.Id = allRegularUsers[i].Id
		regularUserIteration.Username = allRegularUsers[i].Username
		regularUserIteration.Name = allRegularUsers[i].Name
		regularUserIteration.Surname = allRegularUsers[i].Surname
		regularUser = append(regularUser, regularUserIteration)
	}
	return regularUser
}

func createRegularUserProfileDataDto(regularUser *model.RegularUser) *dto.RegularUserProfileDataDTO{
	var regularUserProfileDataDto dto.RegularUserProfileDataDTO

	regularUserProfileDataDto.Name  = regularUser.Name
	regularUserProfileDataDto.Surname = regularUser.Surname
	regularUserProfileDataDto.Username = regularUser.Username
	regularUserProfileDataDto.Biography = regularUser.Biography
	regularUserProfileDataDto.WebSite = regularUser.WebSite

	return &regularUserProfileDataDto
}
func createUserFollowDTOsFromRegularUsers(regularUsers []model.RegularUser) *[]dto.UserFollowDTO {
	var userFollowDTOs []dto.UserFollowDTO
	for i := 0; i < len(regularUsers); i++ {
		var userFollowDto dto.UserFollowDTO
		userFollowDto.Username = regularUsers[i].Username
		userFollowDto.UserId = regularUsers[i].Id.Hex()
		userFollowDTOs = append(userFollowDTOs, userFollowDto)
	}

	return &userFollowDTOs
}
