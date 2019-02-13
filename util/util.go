package util

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 生成 Logid
func GenerateLogid() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.FormatInt(time.Now().UnixNano(), 10) + fmt.Sprintf(":%d", r.Intn(10000))
}

// S is a shortcut for []interface{}
type S []interface{}

// BM is a shortcut for bson.M
type BM bson.M

// 拼接 String 和 Integer
func Join(slice S, sep string) string {
	var tmp []string
	for _, v := range slice {
		switch t := v.(type) {
		case string:
			tmp = append(tmp, v.(string))
		case int:
			tmp = append(tmp, strconv.Itoa(v.(int)))
		default:
			// 异常处理
			fmt.Printf("Error: I don't support type %T!\n", t)
		}
	}
	return strings.Join(tmp, sep)
}

// convert string id to ObjectId
func StringIdToObjectId(id string) *bson.ObjectId {
	if id != "" && bson.IsObjectIdHex(id) {
		bsonObjectId := bson.ObjectIdHex(id)
		return &bsonObjectId
	}
	return nil
}
