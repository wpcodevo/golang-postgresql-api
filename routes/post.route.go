package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-postgresql-api/controllers"
)

type PostRoutes struct {
	postController controllers.PostController
}

func NewRoutePost(postController controllers.PostController) PostRoutes {
	return PostRoutes{postController}
}

func (pc *PostRoutes) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.POST("/", pc.postController.CreatePost)
	router.GET("/", pc.postController.GetAllPosts)
	router.PATCH("/:postId", pc.postController.UpdatePost)
	router.GET("/:postId", pc.postController.GetPostById)
	router.DELETE("/:postId", pc.postController.DeletePostById)
}
