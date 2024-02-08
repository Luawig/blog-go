package routes

import (
	"blog-go/api/handler"
	"blog-go/config"
	"blog-go/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.GetConfig().Server.Mode)
	r := gin.Default()

	// Auth group
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		// Article
		auth.PUT("article/:id", handler.UpdateArticle)
		auth.DELETE("article/:id", handler.DeleteArticle)

		// Category
		auth.PUT("category/:id", handler.UpdateCategory)
		auth.DELETE("category/:id", handler.DeleteCategory)

		// Comment
		auth.PUT("comment/:id", handler.UpdateComment)
		auth.DELETE("comment/:id", handler.DeleteComment)

		// User
		auth.PUT("user/:id", handler.UpdateUser)
		auth.PUT("user/:id/password", handler.UpdateUserPassword)
		auth.DELETE("user/:id", handler.DeleteUser)
	}

	// Public group
	public := r.Group("/api")
	{
		public.POST("login", handler.Login)

		// Article
		public.POST("article", handler.CreateArticle)
		public.GET("article/:id", handler.GetArticle)
		public.GET("articles", handler.GetArticleList)
		public.GET("articles/category/:id", handler.GetArticleListByCategory)
		public.GET("articles/:title", handler.GetArticleListByTitle)

		// Category
		public.POST("category", handler.CreateCategory)
		public.GET("category/:id", handler.GetCategory)
		public.GET("categories", handler.GetCategoryList)

		// Comment
		public.POST("comment", handler.CreateComment)
		public.GET("comment/:id", handler.GetComment)
		public.GET("comments", handler.GetCommentList)
		public.GET("comments/article/:id", handler.GetCommentListByArticle)

		// User
		public.POST("user", handler.CreateUser)
		public.GET("user/:id", handler.GetUser)
		public.GET("users", handler.GetUserList)
		public.GET("users/:username", handler.GetUserListByUsername)

	}

	err := r.Run(config.GetConfig().Server.Port)
	if err != nil {
		panic(err)
		return
	}
}
