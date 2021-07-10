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
	StoryRepository repository.StoryRepository
	NotificationRepository *repository.NotificationRepository
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

func (service *PostService) FindPostById(id primitive.ObjectID) (dto.PostDTO,error) {
	var post dto.PostDTO
	searchPostDocument, err := service.PostRepository.FindPostById(id)
	if err != nil{
		return post,err
	}
	postDto := createPostDTOFromPost(*searchPostDocument)

	return *postDto, nil
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

func (service *PostService) GetAllPostReports() []dto.ReportedPostDTO {
	postReports := service.PostRepository.GetAllPostReports()

	publicPostsModel := createPostReportsFromDocuments(postReports)

	publicPostsDto := createPostReportedPostDtoFromModel(publicPostsModel)

	return publicPostsDto
}

func (service *PostService) CreateNewPost(postUploadDto dto.PostUploadDTO) error {
	fmt.Println("Creating new post")

	post, err := createPostFromPostUploadDTO(&postUploadDto)
	if err != nil {
		return err
	}
	postId, err1 := service.PostRepository.Create(post)
	if err1 != nil {
		return err1
	}

	followersIds, err2 := getUserFollowersWithNotificationsTurnedOn(post.RegularUser.Id)
	if err2 != nil {
		return err2
	}

	notification := CreateNotificationFromEvent(postId, post.RegularUser.Id, followersIds, model.NotificationType(0))
	err3 := service.NotificationRepository.CreateNotification(notification)
	if err3 != nil {
		return err3
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

func createPostReportsFromDocuments(ReportDocuments []bson.D) []model.InappropriateContentPost {
	var reportPosts []model.InappropriateContentPost
	for i := 0; i < len(ReportDocuments); i++ {
		var report model.InappropriateContentPost
		bsonBytes, _ := bson.Marshal(ReportDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &report)
		reportPosts = append(reportPosts, report)
	}
	return reportPosts
}

func createPostReportedPostDtoFromModel(ReportedDocumentsModel []model.InappropriateContentPost) []dto.ReportedPostDTO{
	var reporedPostsDTO []dto.ReportedPostDTO
	for i := 0 ;i < len(ReportedDocumentsModel); i++{
		var reportDTO dto.ReportedPostDTO
		reportDTO.Id = ReportedDocumentsModel[i].Id.Hex()
		reportDTO.Text = ReportedDocumentsModel[i].Text
		reportDTO.Post = ReportedDocumentsModel[i].Post
		reporedPostsDTO = append(reporedPostsDTO, reportDTO)
	}
	return reporedPostsDTO
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

	var usersToNotify []string
	usersToNotify = append(usersToNotify, post.RegularUser.Id)
	notification := CreateNotificationFromEvent(commentDTO.PostId, comment.RegularUser.Id, usersToNotify, model.NotificationType(2))
	err3 := service.NotificationRepository.CreateNotification(notification)
	if err3 != nil {
		return err3
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

		regularUser, err := getRegularUserFromUsername(postLikeDTO.Username)
		if err != nil {
			return err
		}

		var usersToNotify []string
		usersToNotify = append(usersToNotify, post.RegularUser.Id)
		notification := CreateNotificationFromEvent(postLikeDTO.PostId, regularUser.Id, usersToNotify, model.NotificationType(3))
		err3 := service.NotificationRepository.CreateNotification(notification)
		if err3 != nil {
			return err3
		}
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

		regularUser, err := getRegularUserFromUsername(postLikeDTO.Username)
		if err != nil {
			return err
		}

		var usersToNotify []string
		usersToNotify = append(usersToNotify, post.RegularUser.Id)
		notification := CreateNotificationFromEvent(postLikeDTO.PostId, regularUser.Id, usersToNotify, model.NotificationType(3))
		err3 := service.NotificationRepository.CreateNotification(notification)
		if err3 != nil {
			return err3
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

		regularUser, err := getRegularUserFromUsername(postLikeDTO.Username)
		if err != nil {
			return err
		}

		var usersToNotify []string
		usersToNotify = append(usersToNotify, post.RegularUser.Id)
		notification := CreateNotificationFromEvent(postLikeDTO.PostId, regularUser.Id, usersToNotify, model.NotificationType(4))
		err3 := service.NotificationRepository.CreateNotification(notification)
		if err3 != nil {
			return err3
		}
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

		regularUser, err := getRegularUserFromUsername(postLikeDTO.Username)
		if err != nil {
			return err
		}

		var usersToNotify []string
		usersToNotify = append(usersToNotify, post.RegularUser.Id)
		notification := CreateNotificationFromEvent(postLikeDTO.PostId, regularUser.Id, usersToNotify, model.NotificationType(4))
		err3 := service.NotificationRepository.CreateNotification(notification)
		if err3 != nil {
			return err3
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

func (service *PostService) GetAllSavedPostsByUsername(username string) ([]model.Post,error) {
	userSavedPosts, err := getRegularUserSavedPostsByUsername(username)
	if err != nil {
		return nil, err
	}
	var postIds []string
	for i := 0; i < len(userSavedPosts); i++{
		postIds = append(postIds, userSavedPosts[i].PostId)
	}
	publicPostsDocuments := service.PostRepository.GetAllPostsByIds(postIds)
	savedPosts := CreatePostsFromDocuments(publicPostsDocuments)

	return savedPosts, nil
}

func (service *PostService) GetUsersFeed(usersIds []string) (*[]dto.PostDTO, error) {
	var posts []model.Post
	for i:=0; i < len(usersIds); i++ {
		post := service.PostRepository.FindAllPostsByUserId(usersIds[i])
		postDocument := CreatePostsFromDocuments(post)
		posts = appendPosts(posts, postDocument)
	}

	postDTOs := createPostDTOsFromPosts(posts)
	return postDTOs, nil
}

func (service *PostService) DeleteReportedPost(id primitive.ObjectID) error{
	err := service.PostRepository.DeleteReportedPost(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) DeletePost(id primitive.ObjectID) error{
	err := service.PostRepository.DeletePost(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) DeleteUserMediaContent(id primitive.ObjectID) error{
	postDocuments := service.PostRepository.FindAllPostsByUserId(id.Hex())
	postModel := CreatePostsFromDocuments(postDocuments)
	err := service.PostRepository.DeleteUserPostMediaContent(postModel)
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) ReportPost(postReportDTO dto.PostReportDTO) error {
	fmt.Println("Disliking post...")

	postId, _ := primitive.ObjectIDFromHex(postReportDTO.PostId)
	post, err := service.PostRepository.FindPostById(postId)
	if err != nil {
		return err
	}

	regularUser, err := getRegularUserFromUsername(postReportDTO.Username)
	if err != nil {
		return err
	}
	var inappropriateContentPost model.InappropriateContentPost
	inappropriateContentPost.Post = *post
	inappropriateContentPost.RegularUser = *regularUser
	inappropriateContentPost.Text = postReportDTO.Text

	err = service.PostRepository.CreatePostReport(&inappropriateContentPost)
	if err != nil {
		return err
	}
	return nil
}

func createPostDTOsFromPosts(posts []model.Post) *[]dto.PostDTO{
	var postDTOs []dto.PostDTO
	for i := 0; i < len(posts); i++{
		var postDTO dto.PostDTO
		postDTO.Id = posts[i].Id.Hex()
		postDTO.Hashtags = posts[i].Hashtags
		postDTO.Tags = posts[i].Tags
		postDTO.Description = posts[i].Description
		postDTO.MediaPaths = posts[i].MediaPaths
		postDTO.UploadDate = posts[i].UploadDate
		postDTO.MediaContentType = posts[i].MediaContentType
		postDTO.RegularUser = posts[i].RegularUser
		postDTO.Likes = posts[i].Likes
		postDTO.Dislikes = posts[i].Dislikes
		postDTO.Location = posts[i].Location
		postDTO.Comment = posts[i].Comment
		postDTOs = append(postDTOs, postDTO)
	}
	return &postDTOs
}

func createPostDTOFromPost(posts model.Post) *dto.PostDTO{
	var postDTO dto.PostDTO
	postDTO.Id = posts.Id.Hex()
	postDTO.Hashtags = posts.Hashtags
	postDTO.Tags = posts.Tags
	postDTO.Description = posts.Description
	postDTO.MediaPaths = posts.MediaPaths
	postDTO.UploadDate = posts.UploadDate
	postDTO.MediaContentType = posts.MediaContentType
	postDTO.RegularUser = posts.RegularUser
	postDTO.Likes = posts.Likes
	postDTO.Dislikes = posts.Dislikes
	postDTO.Location = posts.Location
	postDTO.Comment = posts.Comment
	return &postDTO
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

func getUserFollowersWithNotificationsTurnedOn(userId string) ([]string, error) {
	requestUrl := fmt.Sprintf("http://%s:%s/followers-with-notifications/%s", os.Getenv("FOLLOW_SERVICE_DOMAIN"), os.Getenv("FOLLOW_SERVICE_PORT"), userId)
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var followersIds []string
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&followersIds)

	return followersIds, nil
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

func getRegularUserSavedPostsByUsername(username string) ([] model.SavedPost, error) {
	requestUrl := fmt.Sprintf("http://%s:%s/saved-posts/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), username)
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var userSavedPosts []model.SavedPost
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&userSavedPosts)

	return userSavedPosts, nil
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

func appendPosts(allPosts []model.Post, newPosts []model.Post) []model.Post{
	for i := 0; i<len(newPosts); i++{
		allPosts = append(allPosts, newPosts[i])
	}
	return allPosts
}