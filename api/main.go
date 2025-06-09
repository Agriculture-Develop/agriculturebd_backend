package main

import (
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/api/routes"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/ioc"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/logs"
)

func init() {

	config.Init()
	logs.Init()

	// 初始化容器
	ioc.Init()
}

func main() {
	// 启动路由
	routes.Router()
}
