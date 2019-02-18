package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// GenerateLogid 生成 logid
func GenerateLogid() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.FormatInt(time.Now().UnixNano(), 10) + fmt.Sprintf(":%d", r.Intn(10000))
}

// S is a shortcut for []interface{}
type S []interface{}

// BM is a shortcut for bson.M
type BM bson.M

// Join 拼接 String 和 Integer
func Join(slice S, sep string) string {
	var tmp []string
	for _, v := range slice {
		switch t := v.(type) {
		case string:
			tmp = append(tmp, t)
		case int:
			tmp = append(tmp, strconv.Itoa(t))
		default:
			// 异常处理
			panic(fmt.Sprintf("Error: I don't support type %T!\n", t))
		}
	}
	return strings.Join(tmp, sep)
}

// StringIDToObjectID convert string id to ObjectId
func StringIDToObjectID(id string) *bson.ObjectId {
	if id != "" && bson.IsObjectIdHex(id) {
		bsonObjectID := bson.ObjectIdHex(id)
		return &bsonObjectID
	}
	return nil
}

// Get 根据特定请求uri，发起get请求返回响应
func Get(uri string, router *gin.Engine) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func PostForm(uri string, param url.Values, router *gin.Engine) []byte {

	// 构造post请求
	req := httptest.NewRequest("POST", uri, strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

// PostJSON 根据特定请求uri和参数param，以Json形式传递参数，发起post请求返回响应
func PostJSON(uri string, param map[string]interface{}, router *gin.Engine) []byte {
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(param)

	// 构造post请求，json数据以请求body的形式传递
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	req.Header.Set("Content-Type", "application/json")

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
