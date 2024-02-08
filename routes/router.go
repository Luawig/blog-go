package routes

import (
	"blog-go/api/handler"
	"blog-go/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.GetConfig().Server.Mode)
	r := gin.Default()

	// Api group
	api := r.Group("/api")
	{
		// Article
		api.POST("article", handler.CreateArticle)
		api.GET("article/:id", handler.GetArticle)
		api.GET("articles", handler.GetArticleList)
		api.GET("articles/category/:id", handler.GetArticleListByCategory)
		api.GET("articles/:title", handler.GetArticleListByTitle)
		api.PUT("article/:id", handler.UpdateArticle)
		api.DELETE("article/:id", handler.DeleteArticle)

		// Category
		api.POST("category", handler.CreateCategory)
		api.GET("category/:id", handler.GetCategory)
		api.GET("categories", handler.GetCategoryList)
		api.PUT("category/:id", handler.UpdateCategory)
		api.DELETE("category/:id", handler.DeleteCategory)

		// Comment
		api.POST("comment", handler.CreateComment)
		api.GET("comment/:id", handler.GetComment)
		api.GET("comments", handler.GetCommentList)
		api.GET("comments/article/:id", handler.GetCommentListByArticle)
		api.PUT("comment/:id", handler.UpdateComment)
		api.DELETE("comment/:id", handler.DeleteComment)

		// User
		api.POST("user", handler.CreateUser)
		api.GET("user/:id", handler.GetUser)
		api.GET("users", handler.GetUserList)
		api.GET("users/:username", handler.GetUserListByUsername)
		api.PUT("user/:id", handler.UpdateUser)
		api.PUT("user/:id/password", handler.UpdateUserPassword)
		api.DELETE("user/:id", handler.DeleteUser)
	}

	err := r.Run(config.GetConfig().Server.Port)
	if err != nil {
		panic(err)
		return
	}
}
