package main

import (
	"goCode/go_gin/middleware"
	"goCode/go_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.LoggerMiddleware())

	// 公开路由
	r.POST("/api/login", login)
	r.GET("/api/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "public endpoint"})
	})

	// 需要认证的路由
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware()) // 认证中间件
	{
		api.GET("protected", func(c *gin.Context) {
			UserId, _ := c.Get("userID") //	 userID
			Username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{
				"message":  "protected endpoint",
				"userId":   UserId,
				"username": Username,
			})
		})
	}

	//================= 静态文件服务 ==============
	//提供静态文件服务
	r.Static("/static", "./static")

	// 提供文件系统
	r.StaticFS("/files", http.Dir("./static"))

	r.Run(":80")

}

// ========== 登录接口 ==========
func login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证用户名密码（示例，实际应从数据库查询）
	if req.Username == "admin" && req.Password == "admin123" {
		token, err := utils.GenerateToken(1, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user": gin.H{
				"id":       1,
				"username": req.Username,
			},
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid credentials",
	})
}
