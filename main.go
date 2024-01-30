package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"tieba/controller"
	"tieba/dao/mysql"
	"tieba/dao/redis"
	"tieba/es"
	"tieba/job"
	"tieba/kafka"
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

	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: tieba config.yaml")
		return
	}
	// 加载配置
	if err := settings.Init(os.Args[1]); err != nil {
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

	// 初始化kafka
	kafka.Init(settings.Conf.KafkaConfig)
	defer kafka.Close()

	// 初始化es
	if err := es.Init(settings.Conf.EsConfig); err != nil {
		fmt.Printf("init es failed, err:%v\n", err)
		return
	}

	go func() {
		err := job.ReadFromKafkaToES(&settings.EsConfig{
			Index:     "post",
			Addresses: []string{"http://localhost:9200"},
		}, context.Background())
		if err != nil {
			panic(err)
		}
	}()

	// 注册路由
	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
