package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin.Context的核心职责是处理请求，返回响应
// 支持的路由类型：静态路由，参数路由，通配符路由

// 用户是查询数据，用GET，参数放到查询参数后面
// 用户是提交数据，用POST，参数写到body之中

//RESTful风格路由，路由表示资源，方法表示对资源的处理
// get    /user/qxx   获取资源,查询
// delete /user/qxx	  删除
// put    /user/qxx   注册
// post  /user/qxx	  修改

func main() {
	server := gin.Default() //返回一个逻辑上的服务器Engine,可以启动多个
	// 静态路由
	server.GET("/hello", func(c *gin.Context) { //路由注册,路由规则与处理函数
		c.String(http.StatusOK, "hello gin")
	})

	server.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "hello post method")
	})

	//参数路由
	server.GET("/user/:name", func(c *gin.Context) {
		s := c.Param("name")
		c.String(http.StatusOK, "hello,这是参数路由:%s", s)
	})

	//通配符路由,*不能单独出现，比如/views/*/html
	server.GET("/views/*.html", func(c *gin.Context) {
		s := c.Param(".html")
		c.String(http.StatusOK, "hello,这是通配符路由:%s", s)
	})

	//查询参数,url跟着的参数/order?id=123
	server.GET("/order", func(c *gin.Context) {
		s := c.Query("id")
		c.String(http.StatusOK, "查询参数:"+s)

	})
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
