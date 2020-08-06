package controller

import (
	"encoding/json"
	"go-firebase/entity"
	"go-firebase/errors"
	"go-firebase/service"
	"net/http"
)

type controller struct {
}

var (
	postSerivce service.PostService
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postSerivce = service
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := postSerivce.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(posts)
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error marshalling the request"})
		return
	}

	err = postSerivce.Validate(&post)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(response).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	result, err := postSerivce.Create(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(result)
}
