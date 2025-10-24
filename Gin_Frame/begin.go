package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("localhost:8080") // 监听并在 0.0.0.0:8080 上启动服务

	// ----------------------------------------------------------------------------------- 路由
	// *普通路由
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"}) })
	router.POST("/login", func(c *gin.Context) {})
	router.Any("/login", func(c *gin.Context) {})

	// *路由分组
	// 简单的路由组: v1
	{
		v1 := router.Group("/v1")
		fmt.Print(v1)
		//v1.POST("/login", loginEndpoint)
		//v1.POST("/submit", submitEndpoint)
		//v1.POST("/read", readEndpoint)
	}
	// 简单的路由组: v2
	{
		v2 := router.Group("/v2")
		fmt.Print(v2)
		//v2.POST("/login", loginEndpoint)
		//v2.POST("/submit", submitEndpoint)
		//v2.POST("/read", readEndpoint)
	}
	// *RESTFUL
	//router.GET("/user", QueryFunc)         //查询
	//router.Post("/user", AddFunc)          // 新增
	//router.Delete("/user", DeleteFunc)     // 删除
	//router.PUT("/user", UpdateFunc)        // 更新（客户端提供完整数据）
	//router.PATCH("/user", PatchUpdateFunc) // 更新（客户端提供需要修改的数据）

	// *重定向
	// 重定向到外部
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	// 重定向到内部
	router.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// *静态资源
	router.Static("/assets", "./assets") // 文件目录
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico") // 单独的文件

	// ---------------------------------------------------------------------------------------- 输出
	// *XML/JSON/TOML/YAML/ProtoBuf
	router.GET("/json", func(c *gin.Context) {
		// JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello JSON",
		})
	})
	router.GET("/xml", func(c *gin.Context) {
		// XML 响应
		c.XML(http.StatusOK, gin.H{
			"message": "Hello XML",
		})
	})
	router.GET("/yaml", func(c *gin.Context) {
		// YAML 响应
		c.YAML(http.StatusOK, gin.H{
			"message": "Hello YAML",
		})
	})
	router.GET("/protobuf", func(c *gin.Context) {
		// ProtoBuf 响应
		c.ProtoBuf(http.StatusOK, "")
	})
	// *HTML
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// ------------------------------------------------------------------------------------------ 参数
	// *参数绑定

}
