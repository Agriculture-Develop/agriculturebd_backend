package routes

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/logs"
	middleware2 "github.com/Agriculture-Develop/agriculturebd/interfaces/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	apiConf = config.Get().Api
)

func Router() {
	// gin模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	//TODO 待抽象
	r.Use(logs.GinLogger, logs.GinRecovery(true), middleware2.RateLimitMiddleware(), middleware2.CORS())

	// 注册添加路由
	registerRoute(r)

	pprof.Register(r)
	// 测试
	addr := fmt.Sprintf("%s:%d", apiConf.Host, apiConf.Post)
	fmt.Printf("服务在%s上成功启动", addr)
	err := r.Run(addr)
	if err != nil {
		fmt.Println(err)
	}

}

func registerRoute(r *gin.Engine) *gin.Engine {

	v1 := r.Group(apiConf.BaseUrl) // v1版

	// 测试
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//AuthModule(v1.Group("auth")) // 登录注册相关路由
	//
	//UserModule(v1.Group("user")) // 用户
	//
	//JudgeModule(v1.Group("judge")) // 在线编译
	//
	//ProgramModule(v1.Group("problem/program")) // 编程题
	//
	//TagModule(v1.Group("tag")) // 标签
	//
	//AdminModels(v1.Group("admin")) // 管理员

	// video.VideoModels(v1) // 视频
	// video.CollectUploadRoutes(v1) // 上传文件
	// course.CourseModels(v1) //

	// 获取静态文件
	//r.StaticFS(baseUrl+"/avatar", http.Dir(config.Get().Api.StaticPath)) // 用户头像文件夹

	return r
}
