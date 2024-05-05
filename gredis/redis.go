package gredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func (r *Redis) NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.RedisAddress,
		Password: r.RedisPasswd,
	})

	r.RedisClient = rdb
	return rdb
}

func (r *Redis) GetKey(key string) *redis.StringCmd {
	get := r.RedisClient.Get(context.Background(), key)
	return get
}

func (r *Redis) SetKey(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	set := r.RedisClient.Set(context.Background(), key, value, expiration)
	return set
}

func (r *Redis) DelKey(key string) *redis.IntCmd {
	del := r.RedisClient.Del(context.Background(), key)
	return del
}

func (r *Redis) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	expire := r.RedisClient.Expire(context.Background(), key, expiration)
	return expire
}

func (r *Redis) XAdd(keys *redis.XAddArgs) *redis.StringCmd {
	return r.RedisClient.XAdd(context.Background(), keys)
}

func (r *Redis) GetStreamsData(streamName string, groupName string, ss string) *redis.XStreamSliceCmd {
	sb := r.RedisClient.XReadStreams(context.Background(), ss)
	// TODO 需要获取所有id ACK一下
	r.RedisClient.XAck(context.Background(), streamName, groupName, ss)
	return sb
}

// Ping 测试联通性
func (r *Redis) Ping() {
	_, err := r.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
