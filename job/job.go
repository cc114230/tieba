package job

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/appengine/log"
	"tieba/es"
	"tieba/kafka"
	settings "tieba/setting"
)

// 帖子数据流处理任务

// Msg 定义kafka中接受到的数据
type Msg struct {
	Type     string                   `json:"type"`
	Database string                   `json:"database"`
	Table    string                   `json:"table"`
	IsDdl    bool                     `json:"isDdl"`
	Data     []map[string]interface{} `json:"data"`
}

func ReadFromKafkaToES(cfg *settings.EsConfig, ctx context.Context) (err error) {
	// 1. 从kafka中获取MySQL中的数据变更消息
	for {
		m, err := kafka.Reader.ReadMessage(ctx)
		if errors.Is(err, context.Canceled) {
			return nil
		}
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		// 2. 将数据写入ES
		msg := new(Msg)
		if err := json.Unmarshal(m.Value, msg); err != nil {
			log.Errorf(ctx, "Unmarshal msg from kafka failed, err:%v", err)
			continue
		}
		//fmt.Printf("msg:%#v", msg)
		if msg.Type == "INSERT" {
			// 往es中新增文档
			for idx := range msg.Data {
				fmt.Println(msg.Data[idx])
				indexDocument(cfg, msg.Data[idx])
			}
		} else {
			// 往es中更新文档
			for idx := range msg.Data {
				updateDocument(cfg, msg.Data[idx])
			}
		}
	}
	return
}

// indexDocument 索引文档
func indexDocument(cfg *settings.EsConfig, d map[string]interface{}) {
	postID := d["post_id"].(string)
	fmt.Println(postID)
	//fmt.Println(cfg.Index)
	//fmt.Println(cfg.Addresses)
	// 添加文档
	resp, err := es.Client.Index(cfg.Index).
		Id(postID).
		Document(d).
		Do(context.Background())
	if err != nil {
		zap.L().Error("indexing document failed", zap.Error(err))
		return
	}
	fmt.Printf("result:%#v\n", resp.Result)
}

// updateDocument 更新文档
func updateDocument(cfg *settings.EsConfig, d map[string]interface{}) {
	postID := d["post_id"].(string)
	resp, err := es.Client.Update(cfg.Index, postID).
		Doc(d). // 使用结构体变量更新
		Do(context.Background())
	if err != nil {
		zap.L().Error("update document failed", zap.Error(err))
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}
