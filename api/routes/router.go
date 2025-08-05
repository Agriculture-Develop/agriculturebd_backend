package routes

import (
	"context"
	"errors"
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/routes/public"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	gracefulRun(fmt.Sprintf("%s:%d", apiConf.Host, apiConf.Post), r)
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
		supplyDemandCtrl Interface.ISupplyDemandCtrl,
	) {
		auth.AuthModels(v1.Group("auth"), authCtrl)
		admin.UserModels(v1.Group("admin"), userCtrl)
		admin.NewsModels(v1.Group("admin"), newsCtrl, newsCategoryCtrl)
		public.SupplyDemandModels(v1.Group("public"), supplyDemandCtrl)
		public.UserModels(v1.Group("public"), userCtrl)
	})
	if err != nil {
		panic(err)
	}

	// 获取静态文件
	r.StaticFS(apiConf.BaseUrl+"/files", http.Dir(config.Get().Api.StaticPath))

	return r
}

func gracefulRun(addr string, handler http.Handler) {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// 启动服务
	go func() {
		log.Printf("Server is running at %s\n", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
