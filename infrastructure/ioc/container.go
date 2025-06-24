package ioc

import (
	authSvc "github.com/Agriculture-Develop/agriculturebd/domain/auth/service"
	newsSvc "github.com/Agriculture-Develop/agriculturebd/domain/news/service"
	"github.com/Agriculture-Develop/agriculturebd/domain/user/service"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/bootstrap"
	authRepo "github.com/Agriculture-Develop/agriculturebd/infrastructure/repository/auth"
	newsRepo "github.com/Agriculture-Develop/agriculturebd/infrastructure/repository/news"
	userRepo "github.com/Agriculture-Develop/agriculturebd/infrastructure/repository/user"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/cache"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/sms"
	newsCtrl "github.com/Agriculture-Develop/agriculturebd/interfaces/controller/admin/news"
	userCtrl "github.com/Agriculture-Develop/agriculturebd/interfaces/controller/admin/user"
	authCtrl "github.com/Agriculture-Develop/agriculturebd/interfaces/controller/auth"
	"go.uber.org/dig"
)

var container *dig.Container

func Init() {
	container = dig.New()

	BuildContainerList()
}

func GetIocContainer() *dig.Container {
	return container
}

// BuildContainerList IOC 注入列表
func BuildContainerList() {
	// 注册组件实现
	mustProvide(sms.NewAliYunSms)
	mustProvide(bootstrap.NewDb)
	mustProvide(cache.NewCache)

	// 注册仓储层实现
	mustProvide(authRepo.NewAuthRepo)
	mustProvide(userRepo.NewUserRepo)

	mustProvide(newsRepo.NewNewsRepo)
	mustProvide(newsRepo.NewNewsCategoryRepo)

	// 注册服务层实现
	mustProvide(authSvc.NewAuthSvc)
	mustProvide(service.NewUserSvc)

	mustProvide(newsSvc.NewNewsService)
	mustProvide(newsSvc.NewNewsCategoryService)

	// 注册控制层实现
	mustProvide(userCtrl.NewUserCtrl)
	mustProvide(authCtrl.NewAuthCtrl)

	mustProvide(newsCtrl.NewCtrl)
	mustProvide(newsCtrl.NewCategoryCtrl)
}

func mustProvide(constructor interface{}) {
	if err := container.Provide(constructor); err != nil {
		panic(err)
	}
}
