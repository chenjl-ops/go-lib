package gredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func (r *RedisSentinel) NewRedisClient() *redis.Client {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		SentinelAddrs: r.RedisSentinelAddress,
		MasterName:    r.RedisMasterName,
		Password:      r.RedisPasswd,
	})

	r.RedisClient = rdb
	return rdb
}

func (r *RedisSentinel) GetKey(key string) *redis.StringCmd {
	get := r.RedisClient.Get(context.Background(), key)
	return get
}

func (r *RedisSentinel) SetKey(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	set := r.RedisClient.Set(context.Background(), key, value, expiration)
	return set
}

func (r *RedisSentinel) DelKey(key string) *redis.IntCmd {
	del := r.RedisClient.Del(context.Background(), key)
	return del
}

func (r *RedisSentinel) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	expire := r.RedisClient.Expire(context.Background(), key, expiration)
	return expire
}

func (r *RedisSentinel) XAdd(keys *redis.XAddArgs) *redis.StringCmd {
	return r.RedisClient.XAdd(context.Background(), keys)
}

func (r *RedisSentinel) GetStreamsData(streamName string, groupName string, ss string) *redis.XStreamSliceCmd {
	sb := r.RedisClient.XReadStreams(context.Background(), ss)
	// TODO 需要获取所有id ACK一下
	r.RedisClient.XAck(context.Background(), streamName, groupName, ss)
	return sb
}

// Ping 测试联通性
func (r *RedisSentinel) Ping() {
	_, err := r.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
