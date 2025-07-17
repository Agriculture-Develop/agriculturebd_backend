package routes

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/middleware"
	"net/http"

	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/api/routes/admin"
	"github.com/Agriculture-Develop/agriculturebd/api/routes/auth"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/ioc"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var apiConf config.Api

func Router() {
	apiConf = config.Get().Api

	// gin模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(middleware.GinLogger(), middleware.GinRecovery(true), middleware.CORS())

	// 注册添加路由
	registerRoute(r)

	// 启动服务
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
	pprof.Register(r)

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 注入控制层依赖
	err := ioc.GetIocContainer().Invoke(func(
		authCtrl Interface.IAuthCtrl,
		userCtrl Interface.IUserCtrl,
		newsCtrl Interface.INewsCtrl,
		newsCategoryCtrl Interface.INewsCategoryCtrl,
	) {
		auth.AuthModels(v1.Group("auth"), authCtrl)
		admin.UserModels(v1.Group("admin"), userCtrl)
		admin.NewsModels(v1.Group("admin"), newsCtrl, newsCategoryCtrl)
	})
	if err != nil {
		panic(err)
	}

	// 获取静态文件
	r.StaticFS(apiConf.BaseUrl+"/files", http.Dir(config.Get().Api.StaticPath))

	return r
}
