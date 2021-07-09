package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/dto"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) CreateNewPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postUploadDto dto.PostUploadDTO
	err := json.NewDecoder(r.Body).Decode(&postUploadDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.CreateNewPost(postUploadDto)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) GetAllRegularUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserPosts := handler.PostService.GetAllRegularUserPosts(username)
	regularUserPostsJson, err := json.Marshal(regularUserPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(regularUserPostsJson)
	}
}

func (handler *PostHandler) GetLocationSearchResults(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	searchInput := param["searchInput"]
	searchPosts := handler.PostService.GetLocationSearchResults(searchInput)
	searchPostsJson, err := json.Marshal(searchPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(searchPostsJson)
	}
}

func (handler *PostHandler) GetUserSearchResults(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	searchInput := param["searchInput"]
	searchPosts := handler.PostService.GetUserSearchResults(searchInput)
	searchPostsJson, err := json.Marshal(searchPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(searchPostsJson)
	}
}

func (handler *PostHandler) GetTagSearchResults(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	searchInput := param["searchInput"]
	searchPosts := handler.PostService.GetTagSearchResults(searchInput)
	searchPostsJson, err := json.Marshal(searchPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(searchPostsJson)
	}
}

func (handler *PostHandler) GetAllPublicPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	publicPosts := handler.PostService.GetAllPublicPosts()
	publicPostsJson, err := json.Marshal(publicPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(publicPostsJson)
	}
}

func (handler *PostHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var commentDTO dto.CommentDTO
	err := json.NewDecoder(r.Body).Decode(&commentDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.CommentPost(commentDTO)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) UpdatePostsPrivacy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var privacyUpdateDTO dto.PrivacyUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&privacyUpdateDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.UpdatePostsPrivacy(&privacyUpdateDTO)
	if err != nil {
		if err.Error() == "Post does not exist!" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postLikeDTO dto.PostLikeDTO
	err := json.NewDecoder(r.Body).Decode(&postLikeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.LikePost(postLikeDTO)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) DislikePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postLikeDTO dto.PostLikeDTO
	err := json.NewDecoder(r.Body).Decode(&postLikeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.DislikePost(postLikeDTO)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) GetAllLikedPostsByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	likedPosts, err := handler.PostService.GetAllLikedPostsByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	likedPostsJson, err := json.Marshal(likedPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(likedPostsJson)
	}
}

func (handler *PostHandler) GetAllDislikedPostsByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	dislikedPosts, err := handler.PostService.GetAllDislikedPostsByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dislikedPostsJson, err := json.Marshal(dislikedPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(dislikedPostsJson)
	}
}

func (handler *PostHandler) GetAllSavedPostsByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	dislikedPosts, err := handler.PostService.GetAllSavedPostsByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dislikedPostsJson, err := json.Marshal(dislikedPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(dislikedPostsJson)
	}
}

func (handler *PostHandler) GetUsersFeed(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")
	var usersIds []string
	err := json.NewDecoder(r.Body).Decode(&usersIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	postDtos, err := handler.PostService.GetUsersFeed(usersIds)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postDtos)
}

func (handler *PostHandler) ReportPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postReportDTO dto.PostReportDTO
	err := json.NewDecoder(r.Body).Decode(&postReportDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.ReportPost(postReportDTO)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PostHandler) GetAllPostReports(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	postReports := handler.PostService.GetAllPostReports()
	postReportsJson, err := json.Marshal(postReports)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(postReportsJson)
	}
}

func (handler *PostHandler) DeleteReportedPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	postReportId,err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.PostService.DeleteReportedPost(postReportId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	postId,err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.PostService.DeletePost(postId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
