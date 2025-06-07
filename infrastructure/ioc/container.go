package ioc

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/admin/user"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/auth"
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
	var err error

	// 注册控制层实现
	err = container.Provide(func(userApi user.Ctrl) Interface.IUserCtrl {
		return user.NewUserCtrl()
	})
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(userApi user.Ctrl) Interface.IAuthCtrl {
		return auth.NewAuthCtrl()
	})
	if err != nil {
		panic(err)
	}

	// 注册仓储实现

}
