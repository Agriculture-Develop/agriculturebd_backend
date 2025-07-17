package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
	"sync"

	"github.com/spf13/viper"
)

const appName = "conf"

var (
	conf             *Default
	configPath       = "api/config/resources/app.yaml" // 默认值
	configSecretPath = "api/config/resources/app_secret.yaml"
	once             sync.Once
)

func Init() {
	once.Do(func() {
		// 初始化配置文件
		conf = new(Default)
		conf.InitConfig()
	})
}

func (app *Default) InitConfig() {

	vip := viper.New()
	// 添加命令行参数修改的配置项
	flag.StringVar(&configPath, "cp", "api/config/resources/app.yaml", "config path")
	flag.Parse()

	// 加载普通配置
	vip.SetConfigType("yaml")
	vip.SetConfigFile(configPath)

	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := vip.Unmarshal(app); err != nil {
		panic(err)
	}

	// 加载敏感配置，覆盖已有字段
	secretVip := viper.New()
	secretVip.SetConfigType("yaml")
	secretVip.SetConfigFile(configSecretPath)
	if err := secretVip.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := secretVip.Unmarshal(app); err != nil {
		panic(err)
	}

	// 动态监听普通配置
	vip.WatchConfig()
	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := vip.Unmarshal(app); err != nil {
			zap.L().Error("Error unmarshaling config: %s", zap.Error(err))
		}
	})
}

func Get() *Default {
	return conf
}
