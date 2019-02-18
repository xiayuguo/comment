package main

import (
	"comment/router"
	"comment/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

var r *gin.Engine

func TestGetComment(t *testing.T) {
	// 初始化请求地址
	uri := "/comments"

	// 发起Get请求
	body := util.Get(uri, r)
	fmt.Printf("response:%v\n", string(body))

	// 判断响应是否和预期一致
	// if string(body) != "success" {
	// 	t.Errorf("响应字符串不符, body:%v\n", string(body))
	// }
}

func init() {
	r = gin.Default()
	router.Init(r)
}
