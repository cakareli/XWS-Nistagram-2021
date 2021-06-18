package service

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/dto"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/repository"
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

func createPostFromPostUploadDTO(postUploadDto *dto.PostUploadDTO) (*model.Post, error){
	regularUser, err := getRegularUserFromUsername(postUploadDto.Username)
	if err != nil {
		return nil, err
	}
	var post model.Post
	post.Tags = postUploadDto.Tags
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
	decoder.Decode(&regularUser)

	return &regularUser, nil
}

