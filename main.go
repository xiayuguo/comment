package main

import (
	. "comment/config"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(strings.Join([]string{Global.Host, strconv.Itoa(Global.Port)}, ":"))
}
