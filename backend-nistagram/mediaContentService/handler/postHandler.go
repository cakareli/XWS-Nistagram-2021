package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/dto"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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
