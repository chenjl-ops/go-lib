package gredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func (r *RedisSentinel) InitRedisClient() *redis.Client {
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

func (r *RedisSentinel) XDel(stream string, keys ...string) *redis.IntCmd {
	return r.RedisClient.XDel(context.Background(), stream, keys...)
}

func (r *RedisSentinel) XLen(stream string) *redis.IntCmd {
	return r.RedisClient.XLen(context.Background(), stream)
}

// LPush 头部插入
func (r *RedisSentinel) LPush(key string, value ...interface{}) *redis.IntCmd {
	return r.RedisClient.LPush(context.Background(), key, value...)
}

// RPush 尾部插入
func (r *RedisSentinel) RPush(key string, value ...interface{}) *redis.IntCmd {
	return r.RedisClient.RPush(context.Background(), key, value...)
}

// LPop 头部弹出
func (r *RedisSentinel) LPop(key string) *redis.StringCmd {
	return r.RedisClient.LPop(context.Background(), key)
}

// RPop 尾部弹出
func (r *RedisSentinel) RPop(key string) *redis.StringCmd {
	return r.RedisClient.RPop(context.Background(), key)
}

func (r *RedisSentinel) LLen(key string) *redis.IntCmd {
	return r.RedisClient.LLen(context.Background(), key)
}

// SAdd 将value添加到集合key中
func (r *RedisSentinel) SAdd(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.SAdd(context.Background(), key, value)
}

// SRem 从集合key中删除value
func (r *RedisSentinel) SRem(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.SRem(context.Background(), key, value)
}

// SMembers 返回集合所有元素
func (r *RedisSentinel) SMembers(key string) *redis.StringSliceCmd {
	return r.RedisClient.SMembers(context.Background(), key)
}

// SIsMember 检查value是否是集合key成员
func (r *RedisSentinel) SIsMember(key string, value interface{}) *redis.BoolCmd {
	return r.RedisClient.SIsMember(context.Background(), key, value)
}

func (r *RedisSentinel) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return r.RedisClient.ZAdd(context.Background(), key, members...)
}

func (r *RedisSentinel) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.RedisClient.ZRange(context.Background(), key, start, stop)
}

func (r *RedisSentinel) ZRank(key string, member string) *redis.IntCmd {
	return r.RedisClient.ZRank(context.Background(), key, member)
}

func (r *RedisSentinel) HGetKey(key string, field string) *redis.StringCmd {
	return r.RedisClient.HGet(context.Background(), key, field)
}

func (r *RedisSentinel) HGetAllKey(key string) *redis.MapStringStringCmd {
	return r.RedisClient.HGetAll(context.Background(), key)
}

func (r *RedisSentinel) HAddKey(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.HSet(context.Background(), key, value)
}

func (r *RedisSentinel) HDelKey(stream string, keys ...string) *redis.IntCmd {
	return r.RedisClient.HDel(context.Background(), stream, keys...)
}

func (r *RedisSentinel) SetBit(key string, offset int64, value int) *redis.IntCmd {
	return r.RedisClient.SetBit(context.Background(), key, offset, value)
}

func (r *RedisSentinel) GetBit(key string, offset int64) *redis.IntCmd {
	return r.RedisClient.GetBit(context.Background(), key, offset)
}

func (r *RedisSentinel) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	return r.RedisClient.PFAdd(context.Background(), key, els...)
}

func (r *RedisSentinel) PFCount(key string) *redis.IntCmd {
	return r.RedisClient.PFCount(context.Background(), key)
}

func (r *RedisSentinel) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	return r.RedisClient.GeoAdd(context.Background(), key, geoLocation...)
}

// GeoDist 计算两个成员之间距离
func (r *RedisSentinel) GeoDist(key string, member1 string, member2 string, unit string) *redis.FloatCmd {
	return r.RedisClient.GeoDist(context.Background(), key, member1, member2, unit)
}

// GeoRadius 查找某个半径范围内所有成员
func (r *RedisSentinel) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.RedisClient.GeoRadius(context.Background(), key, longitude, latitude, query)
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
