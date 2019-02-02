package main


import (

	"strconv"
	"strings"
	. "comment/config"
	"github.com/gin-gonic/gin"
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