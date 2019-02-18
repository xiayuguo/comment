package controller

import (
	"comment/database"
	"comment/util"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// DB is shortcut of database.DB
var DB = database.DB

type (
	// BM is shortcut of database.DB
	BM util.BM
	// CommentController 声明
	CommentController struct{}
)

// GetComment 获取某个评论
func (cc *CommentController) GetComment(c *gin.Context) {
	var result Comment
	id := c.Param("id")
	err := DB.C("comment").FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "not found",
		})
	} else {
		c.JSON(200, gin.H{
			"message": result,
		})
	}
}

// CreateComment 创建评论
func (cc *CommentController) CreateComment(c *gin.Context) {
	var commentForm CommentForm
	c.Bind(&commentForm)
	now := time.Now().Unix()
	err := DB.C("comment").Insert(
		&Comment{
			UserID:     bson.ObjectIdHex(commentForm.UserID),
			ReplyID:    util.StringIDToObjectID(commentForm.ReplyID),
			Content:    commentForm.Content,
			CreateTime: now,
			UpdateTime: now,
		})
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// UpdateComment 更新评论
func (cc *CommentController) UpdateComment(c *gin.Context) {
	var comment UpdateCommentForm
	id := c.Param("id")
	c.Bind(&comment)
	condition := BM{"_id": bson.ObjectIdHex(id)}
	update := BM{"$set": BM{"update_time": time.Now().Unix(), "content": comment.Content}}
	err := DB.C("comment").Update(condition, update)
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// DeleteComment 删除评论
func (cc *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	bsonID := bson.ObjectIdHex(id)
	_, err := DB.C("comment").RemoveAll(BM{"$or": []BM{BM{"_id": bsonID}, BM{"reply_id": bsonID}}})
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// LikeComment 评论点赞
func (cc *CommentController) LikeComment(c *gin.Context) {
	var likeform LikeForm
	commentID := c.Param("id")
	c.Bind(&likeform)
	err := DB.C("like").Insert(
		&Like{
			CommentID: bson.ObjectIdHex(commentID),
			UserID:    bson.ObjectIdHex(likeform.UserID),
			IsLike:    true,
		})
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// DislikeComment 踩评论
func (cc *CommentController) DislikeComment(c *gin.Context) {
	var likeform LikeForm
	commentID := c.Param("id")
	c.Bind(&likeform)
	err := DB.C("like").Insert(
		&Like{
			CommentID: bson.ObjectIdHex(commentID),
			UserID:    bson.ObjectIdHex(likeform.UserID),
			IsLike:    false,
		})
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// DeleteLikeComment 取消喜欢评论
func (cc *CommentController) DeleteLikeComment(c *gin.Context) {
	var likeform LikeForm
	CommentID := c.Param("id")
	c.Bind(&likeform)
	commentBsonID := bson.ObjectIdHex(CommentID)
	userBsonID := bson.ObjectIdHex(likeform.UserID)
	err := DB.C("like").Remove(BM{"$and": []BM{
		BM{"user_id": userBsonID},
		BM{"comment_id": commentBsonID},
		BM{"is_like": true},
	}})
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// DeleteDislikeComment 取消踩评论
func (cc *CommentController) DeleteDislikeComment(c *gin.Context) {
	var likeform LikeForm
	CommentID := c.Param("id")
	c.Bind(&likeform)
	commentBsonID := bson.ObjectIdHex(CommentID)
	userBsonID := bson.ObjectIdHex(likeform.UserID)
	err := DB.C("like").Remove(BM{"$and": []BM{
		BM{"user_id": userBsonID},
		BM{"comment_id": commentBsonID},
		BM{"is_like": false},
	}})
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// LikesComment 获取喜欢评论总数
func (cc *CommentController) LikesComment(c *gin.Context) {
	CommentID := c.Param("id")
	count, err := DB.C("like").FindId(bson.ObjectIdHex(CommentID)).Select(BM{"is_like": true}).Count()
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": count,
	})
}

// DislikesComment 获取踩评论总数
func (cc *CommentController) DislikesComment(c *gin.Context) {
	CommentID := c.Param("id")
	count, err := DB.C("like").FindId(bson.ObjectIdHex(CommentID)).Select(BM{"is_like": false}).Count()
	if err != nil {
		c.Set("err", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": count,
	})
}

// GetAllComment 获取所有评论
func (cc *CommentController) GetAllComment(c *gin.Context) {
	var result []Comment
	err := DB.C("comment").Find(nil).All(&result)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "not found",
		})
	} else {
		c.JSON(200, gin.H{
			"message": result,
		})
	}
}
