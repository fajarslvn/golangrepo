package routes

import (
	"encoding/json"
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
	
	post.ID = len(posts) + 1
	posts = append(posts, post)
	res.WriteHeader(http.StatusOK)

	result, _ := json.Marshal(posts)
	res.Write(result)
}