package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/fajarslvn/go_rest_api/entity"
	"github.com/fajarslvn/go_rest_api/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError )
		res.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError )
		res.Write([]byte(`{"error": "Error unmarshaling the request"}`))
		return
	}
	
	post.ID = rand.Int63()
	repo.Save(&post)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post )
}