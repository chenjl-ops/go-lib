package gredis

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type Client interface {
	Get(key string) *redis.StringCmd
	Set(key string, value string, expiration time.Duration) *redis.StatusCmd
	Del(key string) *redis.IntCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	XAdd(keys *redis.XAddArgs) *redis.StringCmd
	GetStreamsData(streamName string, groupName string, ss string) *redis.XStreamSliceCmd
	Ping()
}
