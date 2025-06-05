package routes

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/api/routes/admin"
	"github.com/Agriculture-Develop/agriculturebd/api/routes/admin/Interface"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/ioc"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	apiConf config.Api
)

func Router() {

	apiConf = config.Get().Api

	// gin模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	//TODO 待抽象
	//r.Use(middleware.GinLogger(), middleware.GinRecovery(true), middleware.RateLimitMiddleware(), middleware.CORS())

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
		userCtrl Interface.IUserCtrl,

	) {

		admin.Models(v1.Group("admin"), userCtrl)
		//product.RegisterProductRoutes(v1.Group("products"), productApi)
	})
	if err != nil {
		panic(err)
	}

	// 获取静态文件
	//r.StaticFS(baseUrl+"/avatar", http.Dir(config.Get().Api.StaticPath)) // 用户头像文件夹

	return r
}
