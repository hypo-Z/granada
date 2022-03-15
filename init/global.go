package init

import (
	"github.com/goworkeryyt/go-config"
	"github.com/goworkeryyt/go-config/env"
	"github.com/goworkeryyt/go-core/db"
	"github.com/goworkeryyt/go-core/global"
	"github.com/goworkeryyt/go-core/redis"
	"github.com/goworkeryyt/go-core/zap"
	"granada1/model/table"
)

func init()  {
	// 获取程序运行环境，默认会读取 resources/active.yaml 文件中配置的运行环境
	global.ENV = env.Active()

	// 获取全局配置,默认根据运行环境加载对应配置文件
	global.CONFIG = goconfig.GlobalConfig()

	// 初始化zap日志
	global.LOG = zap.Zap()

	// 初始化数据库连接
	global.DB = db.Gorm()

	// 创建关系表
	table.AutoMigrate()

	// 初始化 redis 客户端
	global.REDIS = redis.Redis()

	// 初始化 mqtt
	//global.MQTT = mqtt.DefaultMqtt("111111")

	// 获取配置文件原始内容,这样方便在程序中全局拿到自己定义的配置子项
	global.VP = global.CONFIG.Viper
}
