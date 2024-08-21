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

func (r *RedisCluster) XDel(stream string, keys ...string) *redis.IntCmd {
	return r.RedisClient.XDel(context.Background(), stream, keys...)
}

func (r *RedisCluster) XLen(stream string) *redis.IntCmd {
	return r.RedisClient.XLen(context.Background(), stream)
}

// LPush 头部插入
func (r *RedisCluster) LPush(key string, value ...interface{}) *redis.IntCmd {
	return r.RedisClient.LPush(context.Background(), key, value...)
}

// RPush 尾部插入
func (r *RedisCluster) RPush(key string, value ...interface{}) *redis.IntCmd {
	return r.RedisClient.RPush(context.Background(), key, value...)
}

// LPop 头部弹出
func (r *RedisCluster) LPop(key string) *redis.StringCmd {
	return r.RedisClient.LPop(context.Background(), key)
}

// RPop 尾部弹出
func (r *RedisCluster) RPop(key string) *redis.StringCmd {
	return r.RedisClient.RPop(context.Background(), key)
}

func (r *RedisCluster) LLen(key string) *redis.IntCmd {
	return r.RedisClient.LLen(context.Background(), key)
}

// SAdd 将value添加到集合key中
func (r *RedisCluster) SAdd(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.SAdd(context.Background(), key, value)
}

// SRem 从集合key中删除value
func (r *RedisCluster) SRem(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.SRem(context.Background(), key, value)
}

// SMembers 返回集合所有元素
func (r *RedisCluster) SMembers(key string) *redis.StringSliceCmd {
	return r.RedisClient.SMembers(context.Background(), key)
}

// SIsMember 检查value是否是集合key成员
func (r *RedisCluster) SIsMember(key string, value interface{}) *redis.BoolCmd {
	return r.RedisClient.SIsMember(context.Background(), key, value)
}

func (r *RedisCluster) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return r.RedisClient.ZAdd(context.Background(), key, members...)
}

func (r *RedisCluster) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.RedisClient.ZRange(context.Background(), key, start, stop)
}

func (r *RedisCluster) ZRank(key string, member string) *redis.IntCmd {
	return r.RedisClient.ZRank(context.Background(), key, member)
}

func (r *RedisCluster) HGetKey(key string, field string) *redis.StringCmd {
	return r.RedisClient.HGet(context.Background(), key, field)
}

func (r *RedisCluster) HGetAllKey(key string) *redis.MapStringStringCmd {
	return r.RedisClient.HGetAll(context.Background(), key)
}

func (r *RedisCluster) HAddKey(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.HSet(context.Background(), key, value)
}

func (r *RedisCluster) HDelKey(stream string, keys ...string) *redis.IntCmd {
	return r.RedisClient.HDel(context.Background(), stream, keys...)
}

func (r *RedisCluster) SetBit(key string, offset int64, value int) *redis.IntCmd {
	return r.RedisClient.SetBit(context.Background(), key, offset, value)
}

func (r *RedisCluster) GetBit(key string, offset int64) *redis.IntCmd {
	return r.RedisClient.GetBit(context.Background(), key, offset)
}

func (r *RedisCluster) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	return r.RedisClient.PFAdd(context.Background(), key, els...)
}

func (r *RedisCluster) PFCount(key string) *redis.IntCmd {
	return r.RedisClient.PFCount(context.Background(), key)
}

func (r *RedisCluster) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	return r.RedisClient.GeoAdd(context.Background(), key, geoLocation...)
}

// GeoDist 计算两个成员之间距离
func (r *RedisCluster) GeoDist(key string, member1 string, member2 string, unit string) *redis.FloatCmd {
	return r.RedisClient.GeoDist(context.Background(), key, member1, member2, unit)
}

// GeoRadius 查找某个半径范围内所有成员
func (r *RedisCluster) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.RedisClient.GeoRadius(context.Background(), key, longitude, latitude, query)
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
