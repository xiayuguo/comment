package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//生成Logid
func GenerateLogid() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.FormatInt(time.Now().UnixNano(), 10) + fmt.Sprintf(":%d", r.Intn(10000))
}
