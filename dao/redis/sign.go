package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func Sign(userID int64) error {
	year, month, day := time.Now().Date()
	key := getRedisKey(strconv.FormatInt(userID, 10)) + fmt.Sprintf(":signin:%d/%d/%d", year, int(month), day)
	offset := day - 1
	isSigned, err := client.SetBit(context.Background(), key, int64(offset), 1).Result()
	if err != nil {
		return err
	}
	if isSigned == 1 {
		return errors.New("禁止重复签到")
	}
	return err
}
