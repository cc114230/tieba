package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var (
	ErrIsAlreadyCheckin = errors.New("今天已经签到过了")
)

func Sign(userID int64) error {

	today := time.Now().Format("2006-01-02") // 获取当前日期
	year, month, day := time.Now().Date()
	userKey := getRedisKey(strconv.FormatInt(userID, 10)) + fmt.Sprintf(":signin:%d/%d/%d", year, int(month), day)
	// 检查用户今天是否已经签到
	isAlreadyCheckedIn, err := client.GetBit(context.Background(), userKey, getBitPosition(today)).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if isAlreadyCheckedIn == 1 {
		return ErrIsAlreadyCheckin
	}

	//签到
	_, err = client.SetBit(context.Background(), userKey, getBitPosition(today), 1).Result()
	if err != nil {
		return err
	}

	return err
}

// 获取指定日期在位图中的位置
func getBitPosition(date string) int64 {
	// Redis位图是从左到右的，即第0位表示最早日期，以此类推
	// 我们将每天的日期作为位图的一位，以便查找用户是否签到
	// 计算给定日期距离今天的天数
	t, _ := time.Parse("2006-01-02", date)
	today := time.Now()
	daysDiff := today.Sub(t).Hours() / 24
	return int64(daysDiff)
}
