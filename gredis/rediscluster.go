package gredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func (r *RedisCluster) InitRedisClient() *redis.ClusterClient {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		//NewClient: func(opt *redis.Options) *redis.Client {
		//	user, pass := userPassForAddr(opt.Addr)
		//	opt.Username = user
		//	opt.Password = pass
		//	return redis.NewClient(opt)
		//},
		Addrs:    r.RedisAddress,
		Password: r.RedisPasswd,
	})

	r.RedisClient = rdb
	return rdb
}

func (r *RedisCluster) GetKey(key string) *redis.StringCmd {
	get := r.RedisClient.Get(context.Background(), key)
	return get
}

func (r *RedisCluster) SetKey(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	set := r.RedisClient.Set(context.Background(), key, value, expiration)
	return set
}

func (r *RedisCluster) DelKey(key string) *redis.IntCmd {
	del := r.RedisClient.Del(context.Background(), key)
	return del
}

func (r *RedisCluster) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	expire := r.RedisClient.Expire(context.Background(), key, expiration)
	return expire
}

func (r *RedisCluster) XAdd(keys *redis.XAddArgs) *redis.StringCmd {
	return r.RedisClient.XAdd(context.Background(), keys)
}

func (r *RedisCluster) GetStreamsData(streamName string, groupName string, ss string) *redis.XStreamSliceCmd {
	sb := r.RedisClient.XReadStreams(context.Background(), ss)
	// TODO 需要获取所有id ACK一下
	r.RedisClient.XAck(context.Background(), streamName, groupName, ss)
	return sb
}

// Ping 测试联通性
func (r *RedisCluster) Ping() {
	err := r.RedisClient.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
}
