package main

import (
	"fmt"
	"go-firebase/controller"
	router "go-firebase/http"
	"go-firebase/repository"
	"go-firebase/service"
	"net/http"
)

var (
	postRepository = repository.NewFirestoreRepository()
	postService    = service.NewPostService(postRepository)
	postController = controller.NewPostController(postService)
	httpRouter     = router.NewChiRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(response http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(response, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
