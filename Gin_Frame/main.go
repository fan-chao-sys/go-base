package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err)
	}

	// 请求
	router.GET("/ping", func(c *gin.Context) {})
	router.POST("/ping", func(c *gin.Context) {})
	router.PUT("/ping", func(c *gin.Context) {})
	router.PATCH("/ping", func(c *gin.Context) {})
	router.DELETE("/ping", func(c *gin.Context) {})
	router.HEAD("/ping", func(c *gin.Context) {})
	router.OPTIONS("/ping", func(c *gin.Context) {})
	router.Any("/login", func(c *gin.Context) { // 接受所有请求类型
		c.String(200, "login OK!")
	})

	// 路由分组 --访问路径： 127.0.0.1:/v1/v1User || 127.0.0.1:/v2/v1User
	{
		gv1 := router.Group(`v1`)
		gv1.GET("/v1User", func(c *gin.Context) {
			c.String(http.StatusOK, "v1路由")
		})
	}
	{
		gv2 := router.Group("v2")
		gv2.GET("/v2User", func(c *gin.Context) {
			c.String(http.StatusOK, "v2路由")
		})
	}

	// RESTFUL 风格：根据请求类型自动执行 增删改操作
	router.GET("/v3/:id", func(c *gin.Context) {})    // 查
	router.POST("/v3/:id", func(c *gin.Context) {})   // 新增
	router.PATCH("/v3/:id", func(c *gin.Context) {})  // 更新（客户端提供需要修改的数据）
	router.PUT("/v3/:id", func(c *gin.Context) {})    // 更新（客户端提供完整数据）
	router.DELETE("/v3/:id", func(c *gin.Context) {}) // 删除

	// 重定向
	// 重定向到外部
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	// 重定向到内部
	router.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/v1User")
	})
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// 访问静态文件
	router.Static("/assets", "./assets") // 访问文件目录-访问路径： 127.0.0.1:static/assets.txt
	router.StaticFS("/assets", http.Dir("assets"))
	router.StaticFile("/filePath", "./assets/assets.txt") // 访问指定固定文件

	// 请求打印内容方式: JSON、String、XML、YAML、ProtoBuf、Toml、Html
	router.GET("/favicon.ico", func(c *gin.Context) {
		u := &User{
			Name: "test",
			Age:  "18",
		}
		c.JSON(200, u)
		//c.JSON(200, gin.H{"favicon": "/favicon.ico"})
		//c.JSON(200, gin.H{"Name": u.Name,"Age": u.Age})

		c.String(200, "favicon")
		c.XML(200, gin.H{"favicon": "/favicon.ico"})
		c.YAML(200, gin.H{"favicon": "/favicon.ico"})
		c.ProtoBuf(200, gin.H{"favicon": "/favicon.ico"})
		c.TOML(200, gin.H{"favicon": "/favicon.ico"})
		c.HTML(200, "index.tmpl", gin.H{})
	})

	router.LoadHTMLGlob("templates/*") // 加载 HTML 模版
	router.GET("/htmlTest", func(c *gin.Context) {
		c.HTML(200, "user.impl", gin.H{
			"title": "html测试",
		})
	})

	// 正则表达式路由
	router.GET("/users/:id([0-9]+)", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "id:"+id)
	})

	// 路径参数方式
	// 路径绑定接参
	router.GET("/user/:name/:age", func(c *gin.Context) {
		use := &User{}
		er := c.ShouldBindHeader(use) // 对应User对象字段上定义的 header字段指明
		//er := c.ShouldBind(use)       // 解释接受是否正确
		//c.ShouldBindBodyWith(use, binding.JSON) // 对应同时传过来2个不同对像的参数整合一起时,如：user + user2 一起传参
		c.MustBindWith(use, binding.JSON) // 意思：指定的字段必须有值才能取，否则报错，直接退出!

		if er != nil {
			c.JSON(500, gin.H{"error": er, "data": nil})
		}

		c.JSON(200, &User{
			Name: c.Param("name"),
			Age:  c.Param("age"),
		})
	})
	// 路径对象接参 -- 传对象时
	router.GET("/user", func(c *gin.Context) {
		u := c.Query("name")
		a := c.Query("age")
		us := &User{
			Name: u,
			Age:  a,
		}
		c.JSON(200, us)
	})
	// 表单方式：
	router.GET("/user", func(c *gin.Context) {
		u := c.PostForm("name")
		a := c.PostForm("age")
		us := &User{
			Name: u,
			Age:  a,
		}
		c.JSON(200, us)
	})

	// 路由字段、对象验证器
	router.GET("/", func(c *gin.Context) {
		login := LoginInfo{}
		e := c.ShouldBind(&login) // LoginInfo 上字段的类型，定义 binding 自动校验字段的规则!
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
			return
		}
		c.JSON(http.StatusOK, login)
	})
	es := router.Run(":8080")
	if es != nil {
		panic(es)
	}

	// 自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookable date", bookableDate)
	}
	router.GET("/bookable", getBookable)
	router.Run(":8085")

	// 中间件
	// 全局统一调用中间件
	router.Use(Middleware()) // 每次任何接口调用前先 调用执行中间件逻辑
	// 单接口限用
	router.GET("/user/:name/:age", Middleware(), func(c *gin.Context) {})
	// 限定 分组 调用中间件
	rv6 := router.Group("/v6")
	rv6.Use(Middleware())
	// 单接口,多中间件 调用
	router.GET("/user/:name/:age", Middleware(), Middleware2(), func(c *gin.Context) {
		fmt.Println("self")
		// 多个中间件嵌套调用,执行流程: Middleware > Middleware2> self自身
	})

	// basicAuth 自带官方-基础认证-中间件
	// 设置user分组，需要专认证 用户账号, userAuth所有请求都会进行认证
	userAuth := router.Group("/user", gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	userAuth.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		c.String(http.StatusOK, user)
	})

	// HTTPS
	// 自有证书:Openssl 生成证书，如果只是提供API服务，可以用没有经过认证的证书，如果是网站，则需要认证证书
	router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")

	// 开源免费认证证书（Let's Encrypt）
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	fmt.Println(m)
	//log.Fatal(autotl.RunWithManager(r, &m))
}

// Middleware 检查认证 中间件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mw before")
		c.Next() // 调用下个中间件,如没有，及调用接口本身
		fmt.Println("mw after")
	}
}

func Middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mw2 before")
		c.Next() // 调用下个中间件,如没有，及调用接口本身
		fmt.Println("mw2 after")
	}
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

// Booking 包含绑定和验证的数据。
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

type LoginInfo struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"number"`
	Email    string `json:"email" form:"email" binding:"email"`
}

type User struct {
	Name string `json:"name" uri:"name" query:"name" form:"name" header:"name"` // uri 对应  ShouldBind
	Age  string `json:"age" uri:"age" query:"age" form:"age" header:"age"`
}

type User2 struct {
	Name  string `json:"name" uri:"name" query:"name" form:"name" header:"name"`
	Class string `json:"class" uri:"class" query:"class" form:"class" header:"class"`
}
