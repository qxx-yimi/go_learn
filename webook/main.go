package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go_learn/webook/internal/repository"
	"go_learn/webook/internal/repository/dao"
	"go_learn/webook/internal/service"
	"go_learn/webook/internal/web"
	"go_learn/webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

// 跨域问题
// 我们的请求是从localhost:3000 发送到 localhost:8080
// 协议，域名，端口任意一个不同，都是跨域请求
// 解决跨域问题的关键是preflight，第一个请求是preflight，接着才是业务请求，第一个请求会告诉origin，headers等条件，在middleware中指定允许相应的origin和headers
// 即可通过接着的业务请求

//bcrypt是一个号称最安全的加密算法
//优点:
//不需要你自己去生成盐值
//不需要额外存储盐值
//可以通过控制cost来控制加密性能
//同样的文本，加密后的结果不同

func main() {
	db := initDB()
	server := initWebServer()

	u := initUser(db)
	u.RegisterRoutes(server)

	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	//解决跨域问题,所有的请求要经过这里，middleware，决定哪些跨域请求可以得到响应处理，所有业务关心的事情AOP，切面
	// use 作用于全部路由
	server.Use(cors.New(cors.Config{
		// AllowOrigins:           []string{"http://localhost:3000"},
		// AllowMethods:           []string{"POST","GET"},
		AllowHeaders:     []string{"Content-type", "Authorization"},
		AllowCredentials: true, //cookie之类的
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	//cookie存储在哪里
	store := cookie.NewStore([]byte("secret"))
	//cookie的名字叫mysession
	server.Use(sessions.Sessions("mysession", store))
	//登录校验
	server.Use(middleware.NewLoginMiddlewareBuilder().Build())
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		// 只在初始化过程panic，panic相当于整个goroutine结束
		// 一旦初始化过程出错，应用就不要启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
