package util

import (
	"fmt"
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

// 拼接 String 和 Integer
func Join(slice []interface{}, sep string) string {
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
