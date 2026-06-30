package controller

import (
	"goCode/go_blog/config"
	"goCode/go_blog/models"
	"goCode/go_blog/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
}

type CreateCommentRequest struct {
	PostId  int    `json:"postId" binding:"required"`
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

// 发表评论
func (commentController *CommentController) CreateComment(c *gin.Context) {
	var request CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 验证登录用户
	userId, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "未登录")
		return
	}
	var user models.User
	if err := config.DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		utils.BadRequest(c, "用户不存在")
		return
	}

	// 验证帖子是否存在
	var postDb models.Post
	if err := config.DB.Where("id = ?", request.PostId).First(&postDb).Error; err != nil {
		utils.BadRequest(c, "不存在的帖子")
		return
	}

	var commentDb = models.Comment{
		Content: request.Content,
		PostID:  request.PostId,
		UserID:  user.ID,
	}

	if err := config.DB.Create(&commentDb).Error; err != nil {
		utils.InternalServerError(c, "创建评论失败")
		return
	}

	utils.Success(c, "评论成功")
}

// GetComments 获取评论列表
func (commentController *CommentController) GetComments(c *gin.Context) {
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
	var commentList []models.Comment
	if err := config.DB.Order("created_at desc").Offset(offset).Limit(pageSize).Find(&commentList).Error; err != nil {
		utils.BadRequest(c, "获取评论列表失败")
		return
	}

	// 查询总数
	var total int64
	if err := config.DB.Model(models.Comment{}).Count(&total).Error; err != nil {
		utils.BadRequest(c, "获取评论总数失败")
		return
	}
	utils.Success(c, gin.H{
		"data":      commentList,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})

}

// GetComment 获取单个评论
func (commentController *CommentController) GetComment(c *gin.Context) {
	commontId := c.Param("id")
	if commontId == "" {
		utils.BadRequest(c, "评论ID不能为空")
		return
	}

	var commentDb models.Comment
	if err := config.DB.Where("id = ?", commontId).First(&commentDb).Error; err != nil {
		utils.BadRequest(c, "评论不存在")
		return
	}

	utils.Success(c, commentDb)
}
