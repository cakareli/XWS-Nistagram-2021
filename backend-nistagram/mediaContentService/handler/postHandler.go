package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type PostHandler struct {
	PostService *service.PostService
}

func(handler *PostHandler) Hello(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, "hello from controller!")
	handler.PostService.Hello()
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
		w.Write(regularUserPostsJson)
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
		w.Write(publicPostsJson)
	}
}
