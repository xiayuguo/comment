package controller

import (
	// "encoding/json"
	. "comment/database"
	// "comment/logger"
	. "comment/util"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CommentController struct{}

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

func (cc *CommentController) CreateComment(c *gin.Context) {
	var commentForm CommentForm
	c.Bind(&commentForm)
	now := time.Now().Unix()
	err := DB.C("comment").Insert(
		&Comment{
			UserId:     bson.ObjectIdHex(commentForm.UserId),
			ReplyId:    StringIdToObjectId(commentForm.ReplyId),
			Content:    commentForm.Content,
			CreateTime: now,
			UpdateTime: now,
		})
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	var comment struct {
		Content string `json:"content" form:"content"`
	}
	id := c.Param("id")
	c.Bind(&comment)
	condition := BM{"_id": bson.ObjectIdHex(id)}
	update := BM{"$set": BM{"update_time": time.Now().Unix(), "content": comment.Content}}
	err := DB.C("comment").Update(condition, update)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	bsonId := bson.ObjectIdHex(id)
	_, err := DB.C("comment").RemoveAll(BM{"$or": []BM{BM{"_id": bsonId}, BM{"reply_id": bsonId}}})
	if err != nil {
		panic(err)
	}
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
