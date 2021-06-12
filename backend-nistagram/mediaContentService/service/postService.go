package service

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/model"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type PostService struct {
	PostRepository *repository.PostRepository
}

func (service *PostService) Hello () {
	fmt.Printf("Hello from service!")
	service.PostRepository.Hello()
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

func CreatePostsFromDocuments(PostsDocuments []bson.D) []model.Post {
	var publicPosts []model.Post
	for i := 0; i < len(PostsDocuments); i++ {
		var post model.Post
		bsonBytes, _ := bson.Marshal(PostsDocuments[i])
		bson.Unmarshal(bsonBytes, &post)
		publicPosts = append(publicPosts, post)
	}
	return publicPosts
}