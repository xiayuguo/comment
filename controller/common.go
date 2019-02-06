package controller

// 赞踩相关数据
type LikeInfo struct {
	IsLiked      bool  `json:"is_liked"`
	IsDisliked   bool  `json:"is_disliked"`
	LikeTotal    int64 `json:"like_total"`
	DislikeTotal int64 `json:"dislike_total"`
}

// 评论数据
type CommentInfo struct {
	Username string   `json:"username"`    //用户名
	Avatar   string   `json:"avatar"`      //头像
	Time     string   `json:"commenttime"` //评论时间
	Content  string   `json:"content"`     //评论内容
	LikeInfo LikeInfo `json:"like_info"`   // 赞踩相关的信息
}
