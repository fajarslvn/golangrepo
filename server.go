package main

import (
	"fmt"
	"net/http"

	"github.com/fajarslvn/go_rest_api/controller"
	router "github.com/fajarslvn/go_rest_api/http"
	"github.com/fajarslvn/go_rest_api/repository"
	"github.com/fajarslvn/go_rest_api/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService service.PostService = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter router.Router = router.NewMuxRouter() 
)

func main() { 
	const port string = ":8000" 

	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
} 