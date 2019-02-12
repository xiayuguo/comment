package controller

import (
	"gopkg.in/mgo.v2/bson"
)

// 用户表
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`            // 用户id
	Avatar   string        `json:"avatar" bson:"avatar"`     //头像
	Username string        `json:"username" bson:"username"` //用户名
}

// 评论表
type Comment struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	ReplyId    bson.ObjectId `json:"reply_id" bson:"reply_id"`       // 被评论的id
	UserId     bson.ObjectId `json:"user_id" bson:"user_id"`         // 用户的id
	Content    string        `json:"content" bson:"content"`         //评论内容
	CreateTime int64         `json:"create_time" bson:"create_time"` //评论时间
	UpdateTime int64         `json:"update_time" bson:"update_time"` //更新时间
}

// 喜欢表
type Like struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`                // 赞踩id
	CommentId bson.ObjectId `json:"comment_id" bson:"comment_id"` // 评论id
	UserId    bson.ObjectId `json:"user_id" bson:"user_id"`       // 用户id
	IsLike    bool          `json:"is_like" bson:"is_like"`      // 是否喜欢
}
