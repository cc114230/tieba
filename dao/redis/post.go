package redis

import (
	"context"
	"tieba/models"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//按分数从大到小的顺序h查询指定数量的元素
	return client.ZRevRange(context.Background(), key, start, end).Result()
}
