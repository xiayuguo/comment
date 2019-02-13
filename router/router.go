package router

import (
	. "comment/controller"

	"github.com/gin-gonic/gin"
)

func RouterGroupBoth(r *gin.RouterGroup, path string, handler gin.HandlerFunc) {
	r.GET(path, handler)
	r.POST(path, handler)
}

func RouterBoth(r *gin.Engine, path string, handler gin.HandlerFunc) {
	r.GET(path, handler)
	r.POST(path, handler)
}

func Init(r *gin.Engine) {
	initComment(r)
}

//获取评论相关接口
func initComment(r *gin.Engine) {
	comment := r.Group("/comment")
	{
		comment.POST("/", (&CommentController{}).CreateComment)
		comment.GET("/:id", (&CommentController{}).GetComment)
		comment.PUT("/:id", (&CommentController{}).UpdateComment)
		comment.DELETE("/:id", (&CommentController{}).DeleteComment)
		comment.POST("/:id/like", (&CommentController{}).LikeComment)
		comment.POST("/:id/dislike", (&CommentController{}).DislikeComment)
		comment.DELETE("/:id/like", (&CommentController{}).DeleteLikeComment)
		comment.DELETE("/:id/dislike", (&CommentController{}).DeleteDislikeComment)
		comment.GET("/:id/likes", (&CommentController{}).LikesComment)
		comment.GET("/:id/dislikes", (&CommentController{}).DislikesComment)
	}
	r.GET("/comments", (&CommentController{}).GetAllComment)
}
