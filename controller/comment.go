package controller

import (
	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func (cc *CommentController) GetComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": c.Param("id"),
	})
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) ReplyComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) LikeComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) DislikeComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) DeleteLikeComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) DeleteDislikeComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) LikesComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) DislikesComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) GetAllComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}
