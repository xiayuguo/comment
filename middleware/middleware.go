package middleware

import (
	"comment/logger"
	"comment/util"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ErrorMap = map[string]string{
	"UNDEFINED_ERROR":    "未知错误",
	"SERVER_INNER_ERROR": "服务器内部错误",
}

func RecoverHandler(c *gin.Context) {
	if err := recover(); err != nil {
		fmt.Println("panic occur", err)
		c.String(http.StatusOK, "server inner error")
		return
	}
}

//预处理中间件
func RequestMiddleHandler(c *gin.Context) {
	defer RecoverHandler(c)
	c.Set("start_time", time.Now())
	//生成logid
	logid := util.GenerateLogid()
	logger.Debug(logid, "************** Client : ", c.Request.RemoteAddr, " [", c.Request.Method, "] ", c.Request.URL.Path, " *****************")
	logger.Debug(logid, "request : [", c.Request.URL, "]")
	//设置logid
	c.Set("logid", logid)
	c.Next()
}

//响应客户端中间件
func ResponseMiddleHandler(c *gin.Context) {
	defer RecoverHandler(c)
	c.Next()
	logid, _ := c.Get("logid")
	err, exist := c.Get("err")
	//如果发生错误，则返回错误信息
	if exist {
		errMap := map[string]interface{}{"status": 0, "code": err.(string), "message": ErrorMap[err.(string)]}
		c.JSON(http.StatusOK, errMap)
	} else { //否则返回正确结果
		res, exist := c.Get("res")
		if exist {
			c.JSON(http.StatusOK, res)
			resByte, err := json.Marshal(res)
			if err != nil {
				logger.Error("json marshal failed")
			}
			logger.Debug(logid, "Response : ", string(resByte))
		}
	}
	startTime, _ := c.Get("start_time")
	timeEnd := time.Now()
	duration := timeEnd.Sub(startTime.(time.Time))
	logger.Debug(logid, "************Cost Duration : ", duration.String(), " ***********\r\n\r\n")
}

//中间件初始化函数
func Init(r *gin.Engine) {
	//加载预处理中间件
	r.Use(RequestMiddleHandler)
	//加载响应处理中间件
	r.Use(ResponseMiddleHandler)
}
