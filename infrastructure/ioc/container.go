package ioc

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/admin/Interface"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller/admin/user"
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

	// 注册控制层实现
	err := container.Provide(func(userApi user.Ctrl) Interface.IUserCtrl {
		return user.NewUserCtrl()
	})
	if err != nil {
		panic(err)
	}

	// 注册仓储实现

}
