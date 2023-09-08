package logic

import "tieba/dao/redis"

func Sign(userID int64) error {
	return redis.Sign(userID)
}
