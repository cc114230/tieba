package main

import (
	"fmt"
	"github.com/spf13/viper"
	"tieba/controller"
	"tieba/dao/mysql"
	"tieba/dao/redis"
	"tieba/logger"
	"tieba/pkg/snowflake"
	"tieba/router"
	settings "tieba/setting"
)

// @title 贴吧项目接口文档
// @version 1.0
// @contact.name 米兰的小铁匠
// @contact.url 261077486@qq.com
// @host 127.0.0.1:8888
// @BasePath /api/v1
func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 初始化mysql连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	//初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	if err := redis.Init(settings.Conf.RedisConfig); err == nil {
		fmt.Println("redis connect success")
	}
	defer redis.Close()

	if err := snowflake.Init(fmt.Sprintf("%s", viper.GetString("Start_time")),
		viper.GetInt64("machineID")); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
