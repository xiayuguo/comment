package controller

import (
	"gopkg.in/mgo.v2/bson"
)

// User 用户表
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`  // 用户id
	Avatar   string        `json:"avatar" bson:"avatar"`     //头像
	Username string        `json:"username" bson:"username"` //用户名
}

// Comment 评论表
type Comment struct {
	ID         bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	ReplyID    *bson.ObjectId `json:"reply_id" bson:"reply_id"`       // 被评论的id
	UserID     bson.ObjectId  `json:"user_id" bson:"user_id"`         // 用户的id
	Content    string         `json:"content" bson:"content"`         //评论内容
	CreateTime int64          `json:"create_time" bson:"create_time"` //评论时间
	UpdateTime int64          `json:"update_time" bson:"update_time"` //更新时间
}

// Like 喜欢表
type Like struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`      // 赞踩id
	CommentID bson.ObjectId `json:"comment_id" bson:"comment_id"` // 评论id
	UserID    bson.ObjectId `json:"user_id" bson:"user_id"`       // 用户id
	IsLike    bool          `json:"is_like" bson:"is_like"`       // 是否喜欢
}

// CommentForm 评论表单
type CommentForm struct {
	UserID  string `json:"user_id" form:"user_id"`
	ReplyID string `json:"reply_id" form:"reply_id"`
	Content string `json:"content" form:"content"`
}

// UpdateCommentForm 评论更新表单
type UpdateCommentForm struct {
	Content string `json:"content" form:"content"`
}

// LikeForm 喜欢表单
type LikeForm struct {
	UserID string `json:"user_id" form:"user_id"`
}
