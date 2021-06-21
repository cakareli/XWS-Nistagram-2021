package service

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/dto"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
)

type PostService struct {
	PostRepository *repository.PostRepository
}

func (service *PostService) GetAllRegularUserPosts(username string) []model.Post {
	regularUserPostDocuments := service.PostRepository.GetAllByUsername(username)

	regularUserPosts := CreatePostsFromDocuments(regularUserPostDocuments)
	return regularUserPosts
}

func (service *PostService) GetLocationSearchResults(searchInput string) []model.Post {
	searchPublicPostDocuments := service.PostRepository.GetAllPublic()

	searchPosts := CreatePostsFromDocuments(searchPublicPostDocuments)
	searchPostDocuments := service.PostRepository.GetLocationSearchResults(searchInput, searchPosts)

	return searchPostDocuments
}

func (service *PostService) GetUserSearchResults(searchInput string) []model.Post {
	searchPublicPostDocuments := service.PostRepository.GetAllPublic()

	searchPosts := CreatePostsFromDocuments(searchPublicPostDocuments)
	searchPostDocuments := service.PostRepository.GetUserSearchResults(searchInput, searchPosts)

	return searchPostDocuments
}

func (service *PostService) GetTagSearchResults(searchInput string) []model.Post {
	searchPublicPostDocuments := service.PostRepository.GetAllPublic()

	searchPosts := CreatePostsFromDocuments(searchPublicPostDocuments)
	searchPostDocuments := service.PostRepository.GetTagSearchResults(searchInput, searchPosts)

	return searchPostDocuments
}

func (service *PostService) GetAllPublicPosts() []model.Post {
	publicPostsDocuments := service.PostRepository.GetAllPublic()

	publicPosts := CreatePostsFromDocuments(publicPostsDocuments)
	return publicPosts
}

func (service *PostService) CreateNewPost(postUploadDto dto.PostUploadDTO) error {
	fmt.Println("Creating new post")

	post, err := createPostFromPostUploadDTO(&postUploadDto)
	if err != nil {
		return err
	}
	err = service.PostRepository.Create(post)
	if err != nil {
		return err
	}
	return nil
}

func CreatePostsFromDocuments(PostsDocuments []bson.D) []model.Post {
	var publicPosts []model.Post
	for i := 0; i < len(PostsDocuments); i++ {
		var post model.Post
		bsonBytes, _ := bson.Marshal(PostsDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &post)
		publicPosts = append(publicPosts, post)
	}
	return publicPosts
}

func (service *PostService) CommentPost(commentDTO dto.CommentDTO) error {
	fmt.Println("Commenting post...")

	comment, err := createCommentFromCommentDTO(&commentDTO)
	if err != nil {
		return err
	}
	postId, _ := primitive.ObjectIDFromHex(commentDTO.PostId)
	post, err := service.PostRepository.FindPostById(postId)
	if err != nil {
		return err
	}

	appendedComments := append(post.Comment, *comment)
	post.Comment = appendedComments
	err = service.PostRepository.Update(post)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) LikePost(postLikeDTO dto.PostLikeDTO) error {
	fmt.Println("Liking post...")

	postId, _ := primitive.ObjectIDFromHex(postLikeDTO.PostId)
	post, err := service.PostRepository.FindPostById(postId)
	if err != nil {
		return err
	}
	userLikedAndDisliked, err := getRegularUserLikedAndDislikedPostsByUsername(postLikeDTO.Username)
	if err != nil {
		return err
	}

	if(!contains(userLikedAndDisliked.LikedPostsIds, postLikeDTO.PostId) && !contains(userLikedAndDisliked.DislikedPostsIds, postLikeDTO.PostId)){
		post.Likes = post.Likes + 1
		updateUserLikedPosts(postLikeDTO, "yes")
	}
	if(!contains(userLikedAndDisliked.LikedPostsIds, postLikeDTO.PostId) && contains(userLikedAndDisliked.DislikedPostsIds, postLikeDTO.PostId)){
		post.Dislikes = post.Dislikes -1
		post.Likes = post.Likes + 1
		err := updateUserLikedPosts(postLikeDTO, "yes")
		if( err != nil){
			fmt.Println(err)
		}

		err = updateUserDislikedPosts(postLikeDTO, "no")
		if( err != nil){
			fmt.Println(err)
		}
	}
	if(contains(userLikedAndDisliked.LikedPostsIds, postLikeDTO.PostId)){
		post.Likes = post.Likes -1
		err := updateUserLikedPosts(postLikeDTO, "no")
		if( err != nil){
			fmt.Println(err)
		}
	}

	err = service.PostRepository.Update(post)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) DislikePost(postLikeDTO dto.PostLikeDTO) error {
	fmt.Println("Disliking post...")

	postId, _ := primitive.ObjectIDFromHex(postLikeDTO.PostId)
	post, err := service.PostRepository.FindPostById(postId)
	if err != nil {
		return err
	}
	userLikedAndDisliked, err := getRegularUserLikedAndDislikedPostsByUsername(postLikeDTO.Username)
	if err != nil {
		return err
	}

	if (!contains(userLikedAndDisliked.DislikedPostsIds, postLikeDTO.PostId) && !contains(userLikedAndDisliked.LikedPostsIds, postLikeDTO.PostId)) {
		post.Dislikes = post.Dislikes + 1
		updateUserDislikedPosts(postLikeDTO, "yes")
	}
	if (!contains(userLikedAndDisliked.DislikedPostsIds, postLikeDTO.PostId) && contains(userLikedAndDisliked.LikedPostsIds, postLikeDTO.PostId)) {
		post.Likes = post.Likes - 1
		post.Dislikes = post.Dislikes + 1
		err := updateUserDislikedPosts(postLikeDTO, "yes")
		if (err != nil) {
			fmt.Println(err)
		}

		err = updateUserLikedPosts(postLikeDTO, "no")
		if (err != nil) {
			fmt.Println(err)
		}
	}
	if (contains(userLikedAndDisliked.DislikedPostsIds, postLikeDTO.PostId)) {
		post.Dislikes = post.Dislikes - 1
		err := updateUserDislikedPosts(postLikeDTO, "no")
		if (err != nil) {
			fmt.Println(err)
		}
	}

	err = service.PostRepository.Update(post)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) UpdatePostsPrivacy(privacyUpdateDTO *dto.PrivacyUpdateDTO) error {
	fmt.Println("updating posts privacy...")
	userPostsDocuments := service.PostRepository.FindAllPostsByUserId(privacyUpdateDTO.Id)
	userPosts := CreatePostsFromDocuments(userPostsDocuments)

	var privacyType model.PrivacyType
	if privacyUpdateDTO.PrivacyType == "0" {
		privacyType = model.PrivacyType(0)
	} else {
		privacyType = model.PrivacyType(1)
	}
	for i:=0; i < len(userPosts); i++ {
		err := service.PostRepository.UpdatePostPrivacy(privacyType, userPosts[i].Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *PostService) GetAllLikedPostsByUsername(username string) ([]model.Post,error) {
	userLikedAndDisliked, err := getRegularUserLikedAndDislikedPostsByUsername(username)
	if err != nil {
		return nil, err
	}

	publicPostsDocuments := service.PostRepository.GetAllPostsByIds(userLikedAndDisliked.LikedPostsIds)
	likedPosts := CreatePostsFromDocuments(publicPostsDocuments)

	return likedPosts, nil
}

func (service *PostService) GetAllDislikedPostsByUsername(username string) ([]model.Post,error) {
	userLikedAndDisliked, err := getRegularUserLikedAndDislikedPostsByUsername(username)
	if err != nil {
		return nil, err
	}

	publicPostsDocuments := service.PostRepository.GetAllPostsByIds(userLikedAndDisliked.DislikedPostsIds)
	dislikedPosts := CreatePostsFromDocuments(publicPostsDocuments)

	return dislikedPosts, nil
}

func createPostFromPostUploadDTO(postUploadDto *dto.PostUploadDTO) (*model.Post, error){
	regularUser, err := getRegularUserFromUsername(postUploadDto.Username)
	if err != nil {
		return nil, err
	}
	var post model.Post
	post.Hashtags = postUploadDto.Hashtags
	var tags []model.RegularUser
	for i := 0; i < len(postUploadDto.Tags); i++{
		var regUser *model.RegularUser
		regUser, err = getRegularUserFromUsername(postUploadDto.Tags[i])
		regUser.Username = postUploadDto.Tags[i]
		tags = append(tags,*regUser)

	}
	post.Tags = tags
	post.Description = postUploadDto.Description
	post.MediaPaths = postUploadDto.MediaPaths
	post.UploadDate = postUploadDto.UploadDate
	post.Location = postUploadDto.Location
	post.RegularUser = *regularUser
	post.RegularUser.Username = postUploadDto.Username
	post.Likes = 0
	post.Dislikes = 0
	if len(postUploadDto.MediaPaths) > 1 {
		post.MediaContentType = model.MediaContentType(2)
	} else {
		post.MediaContentType = model.MediaContentType(0)
	}
	return &post, nil
}

func createCommentFromCommentDTO(commentDTO *dto.CommentDTO) (*model.Comment, error){
	regularUser, err := getRegularUserFromUsername(commentDTO.Username)
	if err != nil {
		return nil, err
	}
	var comment model.Comment
	comment.RegularUser = *regularUser
	comment.RegularUser.Username = commentDTO.Username
	comment.Text = commentDTO.Text

	return &comment, nil
}

func getRegularUserFromUsername(username string) (*model.RegularUser, error) {
	requestUrl := fmt.Sprintf("http://%s:%s/by-username/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), username)
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var regularUser model.RegularUser
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&regularUser)

	return &regularUser, nil
}

func getRegularUserLikedAndDislikedPostsByUsername(username string) (*dto.UserLikedAndDislikedDTO, error) {
	requestUrl := fmt.Sprintf("http://%s:%s/liked-and-disliked/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), username)
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var userLikesAndDislikes dto.UserLikedAndDislikedDTO
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&userLikesAndDislikes)

	return &userLikesAndDislikes, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func updateUserLikedPosts(postLikeDTO dto.PostLikeDTO, isAdd string) error {
	postBody, _ := json.Marshal(map[string]string{
		"username": postLikeDTO.Username,
		"postId": postLikeDTO.PostId,
		"isAdd" : isAdd,
	})
	requestUrl := fmt.Sprintf("http://%s:%s/update-liked-posts", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func updateUserDislikedPosts(postLikeDTO dto.PostLikeDTO, isAdd string) error {
	postBody, _ := json.Marshal(map[string]string{
		"username": postLikeDTO.Username,
		"postId": postLikeDTO.PostId,
		"isAdd" : isAdd,
	})
	requestUrl := fmt.Sprintf("http://%s:%s/update-disliked-posts", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}