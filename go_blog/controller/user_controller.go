package controller

import (
	"fmt"
	"goCode/go_blog/config"
	"goCode/go_blog/models"
	"goCode/go_blog/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// RegisterRequest 注册请求参数
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// AuthResponse 用户认证/注册成功返回
type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// LoginRequest 登录请求参数
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 用户注册
func (UserController *UserController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// 检查用户名是否已经存在
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err == nil {
		utils.BadRequest(c, "用户名已经存在："+req.Username)
		return
	}

	// 检查Email是否已经存在
	if err := config.DB.Where("email = ?").First(&user).Error; err == nil {
		utils.BadRequest(c, "Email已经存在："+req.Email)
		return
	}

	// 创建用户
	user = models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	// 生成Token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		fmt.Println("生成Token失败:", err)
		utils.InternalServerError(c, err.Error())
		return
	}

	// 注册成功返回
	utils.Success(c, AuthResponse{token, user})
}

// 用户登录
func (UserController *UserController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	// 查找用户
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.BadRequest(c, "用户名不存在")
		return
	}

	// 验证密码
	if !user.CheckPassword(req.Password) {
		utils.BadRequest(c, "密码错误")
		return
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}
	utils.Success(c, AuthResponse{token, user})
}

func (UserController *UserController) GetUserInfo(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "获取用户信息失败，请登录")
	}

	var user models.User
	if err := config.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		utils.BadRequest(c, "用户不存在")
		return
	}
	utils.Success(c, user)

}
