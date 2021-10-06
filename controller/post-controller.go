package controller

import (
	"encoding/json"
	"net/http"

	"github.com/fajarslvn/go_rest_api/entity"
	"github.com/fajarslvn/go_rest_api/errorz"
	"github.com/fajarslvn/go_rest_api/service"
)

type controller struct {}

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError )
		json.NewEncoder(res).Encode(errorz.ServiceError{Message: "Error getting the posts"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func (*controller) AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError )
		json.NewEncoder(res).Encode(errorz.ServiceError{Message: "Error unmarshaling the request"})
		return
	}
 
	err1 := postService.Validate(&post)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError )
		json.NewEncoder(res).Encode(errorz.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError )
		json.NewEncoder(res).Encode(errorz.ServiceError{Message: "error saving post "})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}