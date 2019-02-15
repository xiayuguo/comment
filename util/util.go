package util

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
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
