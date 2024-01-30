package kafka

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	settings "tieba/setting"
)

var Reader *kafka.Reader

func Init(cfg *settings.KafkaConfig) {
	// readByReader 通过Reader接收消息
	// 创建Reader
	Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{fmt.Sprintf("%s", cfg.Brokers[0])},
		GroupID:  cfg.GroupID,
		Topic:    cfg.Topic,
		MaxBytes: 10e6, // 10MB
		//StartOffset: kafka.FirstOffset,
	})
}

func Close() {
	// 程序退出前关闭Reader
	if err := Reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
