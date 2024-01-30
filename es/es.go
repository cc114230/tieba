package es

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	settings "tieba/setting"
)

var Client *elasticsearch.TypedClient

func Init(conf *settings.EsConfig) (err error) {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{ //"http://localhost:9200",
			fmt.Sprintf("%s", conf.Addresses[0]),
		},
	}
	// 创建客户端连接
	Client, err = elasticsearch.NewTypedClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewTypedClient failed, err:%v\n", err)
		return err
	}
	return
}
