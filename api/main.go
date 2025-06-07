package main

import (
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/api/routes"
	database "github.com/Agriculture-Develop/agriculturebd/infrastructure/dao/init"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/ioc"
)

func init() {
	// 初始化配置
	config.Init()

	// 初始数据库
	database.Init()

	// 初始化容器
	ioc.Init()
}

func main() {

	// 启动路由
	routes.Router()
}
