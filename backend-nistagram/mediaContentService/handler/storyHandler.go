package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/dto"
	"XWS-Nistagram-2021/backend-nistagram/mediaContentService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (handler *StoryHandler) UpdateStoriesPrivacy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var privacyUpdateDTO dto.PrivacyUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&privacyUpdateDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.StoryService.UpdateStoriesPrivacy(&privacyUpdateDTO)
	if err != nil {
		if err.Error() == "Story does not exist!" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
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

func (handler *StoryHandler) ReportStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var storyReportDTO dto.StoryReportDTO
	err := json.NewDecoder(r.Body).Decode(&storyReportDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.StoryService.ReportStory(storyReportDTO)
	if err != nil {
		if err.Error() == "regular user is NOT found" {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *StoryHandler) GetAllStoryReports(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	storyReports := handler.StoryService.GetAllStoryReports()
	storyReportsJson, err := json.Marshal(storyReports)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(storyReportsJson)
	}
}

func (handler *StoryHandler) DeleteReportedStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	storyReportId,err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.StoryService.DeleteReportedStory(storyReportId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryHandler) DeleteStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	storyId,err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.StoryService.DeleteStory(storyId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}