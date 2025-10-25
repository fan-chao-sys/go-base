package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		// 解析验证token
		// 用户信息存入context
		user := User{}
		c.Set("user", user)
		c.Next()
		c.String(200, tokenString)
	}
}

// LatencyLogger 耗时统计中间件
func LatencyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Printf("%s %s cost: %v", c.Request.Method, c.Request.URL.Path, latency)
	}
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{Code: 0, Data: data})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(200, Response{Code: code, Message: msg})
}

func main() {
	r := gin.Default()
	r.Use(AuthMiddleware())
}
