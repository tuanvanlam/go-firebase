package main

import (
	"encoding/json"
	"go-firebase/entity"
	"go-firebase/repository"
	"math/rand"
	"net/http"
)

var (
	repo = repository.NewPostRepository()
)

func getPosts(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, _ = response.Write([]byte(`{"error": "Error getting the posts"}`))
	}

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(posts)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, _ = response.Write([]byte(`{"error": "Error marshalling the request"}`))
		return
	}

	post.ID = rand.Int63()
	_, _ = repo.Save(&post)

	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(post)
}
