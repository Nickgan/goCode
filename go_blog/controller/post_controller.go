package controller

import (
	"goCode/go_blog/config"
	"goCode/go_blog/models"
	"goCode/go_blog/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
}

// CreatePostRequest 新建帖子
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=5,max=10"`    //帖子标题，最小长度5，最大长度10
	Content string `json:"content" binding:"required,min=5,max=300"` //帖子内容，最小长度5，最大长度300
}

// UpdatePostRequest 更新帖子
type UpdatePostRequest struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required,min=5,max=10"`    //帖子标题，最小长度5，最大长度10
	Content string `json:"content" binding:"required,min=5,max=300"` //帖子内容，最小长度5，最大长度300
}

// CreatePost 创建帖子
func (postController *PostController) CreatePost(c *gin.Context) {
	var createPostRequest CreatePostRequest
	if err := c.ShouldBindJSON(&createPostRequest); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// 查询当前登录用户
	userId, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未登录")
		return
	}
	var user models.User
	if err := config.DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		utils.BadRequest(c, "用户不存在")
		return
	}

	// 创建帖子
	post := models.Post{
		Title:   createPostRequest.Title,
		Content: createPostRequest.Content,
		UserID:  user.ID,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		utils.BadRequest(c, "创建帖子失败")
		return
	}

	utils.Success(c, "post")
}

// UpdatePost 更新帖子
func (postController *PostController) UpdatePost(c *gin.Context) {
	var updatePostRequest UpdatePostRequest
	if err := c.ShouldBindJSON(&updatePostRequest); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// 检查帖子是否存在
	var postDb models.Post
	if err := config.DB.Where("id = ?", updatePostRequest.ID).First(&postDb).Error; err != nil {
		utils.BadRequest(c, "帖子不存在")
		return
	}

	// 检查发帖人是否当前登录用户
	userId, exists := c.Get("user_id")
	if !exists || postDb.UserID != userId {
		utils.Unauthorized(c, "无权限更新帖子")
		return
	}

	// 更新文章
	postDb.Title = updatePostRequest.Title
	postDb.Content = updatePostRequest.Content
	if err := config.DB.Save(&postDb).Error; err != nil {
		utils.BadRequest(c, "更新帖子失败")
		return
	}

	utils.Success(c, postDb)
}

// DeletePost 删除帖子
func (postController *PostController) DeletePost(c *gin.Context) {
	postId := c.Param("id")
	if postId == "" {
		utils.BadRequest(c, "帖子ID不能为空")
		return
	}

	// 判断帖子是否存在
	var postDb models.Post
	if err := config.DB.Where("id = ?", postId).First(&postDb).Error; err != nil {
		utils.BadRequest(c, "帖子不存在")
		return
	}

	// 判断帖子主人是否是当前登录用户
	userId, exists := c.Get("user_id")
	if !exists || postDb.UserID != userId {
		utils.Unauthorized(c, "无权限删除帖子")
		return
	}

	// 删除帖子
	if err := config.DB.Delete(&postDb).Error; err != nil {
		utils.BadRequest(c, "删除帖子失败")
		return
	}

	utils.Success(c, "帖子删除成功")
}

// GetPosts 获取帖子列表
func (postController *PostController) GetPosts(c *gin.Context) {

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))           // 默认第一页
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10")) // 默认每页10条

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// 查询列表
	var postList []models.Post
	if err := config.DB.Order("created_at desc").Offset(offset).Limit(pageSize).Find(&postList).Error; err != nil {
		utils.BadRequest(c, "获取帖子列表失败")
		return
	}

	// 查询总数
	var total int64
	if err := config.DB.Model(models.Post{}).Count(&total).Error; err != nil {
		utils.BadRequest(c, "获取帖子总数失败")
		return
	}

	utils.Success(c, gin.H{
		"data":      postList,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetPost 获取单个
func (postController *PostController) GetPost(c *gin.Context) {
	postId := c.Param("id")
	if postId == "" {
		utils.BadRequest(c, "帖子ID不能为空")
		return
	}

	var postDb models.Post
	if err := config.DB.Where("id = ?", postId).First(&postDb).Error; err != nil {
		utils.BadRequest(c, "帖子不存在")
		return
	}

	utils.Success(c, postDb)
}
