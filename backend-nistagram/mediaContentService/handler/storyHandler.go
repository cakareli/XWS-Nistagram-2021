package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/dto"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type StoryHandler struct {
	StoryService *service.StoryService
}

func (handler *StoryHandler) CreateNewStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var storyUploadDTO dto.StoryUploadDTO
	err := json.NewDecoder(r.Body).Decode(&storyUploadDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.StoryService.CreateNewStory(storyUploadDTO)
	if err != nil {
		if err.Error() == "Regular user is NOT found!" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *StoryHandler) GetAllRegularUserStories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserStories := handler.StoryService.GetAllRegularUserStories(username)
	regularUserStoriesJson, err := json.Marshal(regularUserStories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(regularUserStoriesJson)
	}
}