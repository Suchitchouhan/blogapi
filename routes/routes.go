// routes/routes.go

package routes

import (
	"blogapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Blog routes
	blog := r.Group("/blog")
	{
		blog.GET("/", controllers.GetBlogPosts)
		blog.GET("/:id", controllers.GetBlogPost)
		blog.POST("/", controllers.CreateBlogPost)
		blog.PUT("/:id", controllers.UpdateBlogPost)
		blog.DELETE("/:id", controllers.DeleteBlogPost)
	}

	return r
}
