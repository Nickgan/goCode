package routes

import (
	"goCode/go_blog/controller"
	"goCode/go_blog/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes() *gin.Engine {
	r := gin.New()

	var userController = &controller.UserController{}
	var postController = &controller.PostController{}
	var commentController = &controller.CommentController{}

	// 使用中间件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandlerMiddleware())
	r.Use(gin.Recovery())

	// API路由分组
	api := r.Group("/api/v1")
	{

		// =========================需要认证的路由=========================
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware())
		{
			// 用户信息
			authenticated.GET("/profile", userController.GetUserInfo)

			// 文章相关路由
			posts := authenticated.Group("/posts")
			{
				posts.POST("/createPost", postController.CreatePost) // 发帖
				posts.GET("/update", postController.UpdatePost)      // 更新帖子
				posts.GET("/delete", postController.DeletePost)      // 删除
			}

			// 评论相关路由
			comments := authenticated.Group("/comments")
			{
				comments.POST("/createComment", commentController.CreateComment) // 发表评论
			}
		}

		// =========================公开路由（无需认证）=========================

		// 认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/login", userController.Login)
			auth.POST("/register", userController.Register)
		}

		// 公开路由
		public := api.Group("")
		{

			// 文章相关路由
			postsPublic := public.Group("/posts")
			{
				postsPublic.GET("/getPosts", postController.GetPosts) // 获取帖子列表
				postsPublic.GET("/getPost", postController.GetPost)   // 获取单个帖子
			}

			// 评论相关路由
			commentsPublic := public.Group("/comments")
			{
				commentsPublic.GET("/getComments", commentController.GetComments) // 获取评论列表
				commentsPublic.GET("/getComment", commentController.GetComment)   // 获取单个评论
			}

		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK", "status": "healthy"})
	})

	return r
}
