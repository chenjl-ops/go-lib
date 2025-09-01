package gredis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

/*
XAdd stream 流数据类型
HAdd hashmap 数据类型
LAdd list 列表数据类型
SAdd set 无序集合数据类型
ZAdd sorted set 有序集合数据类型

SetBit bitmap 位图数据类型
PFAdd Hyperloglog 一种基数估算算法
GEOAdd geospatial 地理空间 类型；例如 计算亮点之间距离、查询某个范围内的地点等
*/

func (r *Redis) InitRedisClient() *redis.Client {
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

// XAdd stream 数据类型
func (r *Redis) XAdd(keys *redis.XAddArgs) *redis.StringCmd {
	return r.RedisClient.XAdd(context.Background(), keys)
}

func (r *Redis) XDel(stream string, keys ...string) *redis.IntCmd {
	return r.RedisClient.XDel(context.Background(), stream, keys...)
}

func (r *Redis) XLen(stream string) *redis.IntCmd {
	return r.RedisClient.XLen(context.Background(), stream)
}

// LPush 头部插入
func (r *Redis) LPush(key string, value ...interface{}) *redis.IntCmd {
	return r.RedisClient.LPush(context.Background(), key, value...)
}

// RPush 尾部插入
func (r *Redis) RPush(key string, value ...interface{}) *redis.IntCmd {
	return r.RedisClient.RPush(context.Background(), key, value...)
}

// LPop 头部弹出
func (r *Redis) LPop(key string) *redis.StringCmd {
	return r.RedisClient.LPop(context.Background(), key)
}

// RPop 尾部弹出
func (r *Redis) RPop(key string) *redis.StringCmd {
	return r.RedisClient.RPop(context.Background(), key)
}

func (r *Redis) LLen(key string) *redis.IntCmd {
	return r.RedisClient.LLen(context.Background(), key)
}

// SAdd 将value添加到集合key中
func (r *Redis) SAdd(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.SAdd(context.Background(), key, value)
}

// SRem 从集合key中删除value
func (r *Redis) SRem(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.SRem(context.Background(), key, value)
}

// SMembers 返回集合所有元素
func (r *Redis) SMembers(key string) *redis.StringSliceCmd {
	return r.RedisClient.SMembers(context.Background(), key)
}

// SIsMember 检查value是否是集合key成员
func (r *Redis) SIsMember(key string, value interface{}) *redis.BoolCmd {
	return r.RedisClient.SIsMember(context.Background(), key, value)
}

func (r *Redis) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return r.RedisClient.ZAdd(context.Background(), key, members...)
}

func (r *Redis) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.RedisClient.ZRange(context.Background(), key, start, stop)
}

func (r *Redis) ZRank(key string, member string) *redis.IntCmd {
	return r.RedisClient.ZRank(context.Background(), key, member)
}

func (r *Redis) HGetKey(key string, field string) *redis.StringCmd {
	return r.RedisClient.HGet(context.Background(), key, field)
}

func (r *Redis) HGetAllKey(key string) *redis.MapStringStringCmd {
	return r.RedisClient.HGetAll(context.Background(), key)
}

func (r *Redis) HAddKey(key string, value interface{}) *redis.IntCmd {
	return r.RedisClient.HSet(context.Background(), key, value)
}

func (r *Redis) HDelKey(stream string, keys ...string) *redis.IntCmd {
	return r.RedisClient.HDel(context.Background(), stream, keys...)
}

func (r *Redis) SetBit(key string, offset int64, value int) *redis.IntCmd {
	return r.RedisClient.SetBit(context.Background(), key, offset, value)
}

func (r *Redis) GetBit(key string, offset int64) *redis.IntCmd {
	return r.RedisClient.GetBit(context.Background(), key, offset)
}

func (r *Redis) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	return r.RedisClient.PFAdd(context.Background(), key, els...)
}

func (r *Redis) PFCount(key string) *redis.IntCmd {
	return r.RedisClient.PFCount(context.Background(), key)
}

func (r *Redis) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	return r.RedisClient.GeoAdd(context.Background(), key, geoLocation...)
}

// GeoDist 计算两个成员之间距离
func (r *Redis) GeoDist(key string, member1 string, member2 string, unit string) *redis.FloatCmd {
	return r.RedisClient.GeoDist(context.Background(), key, member1, member2, unit)
}

// GeoRadius 查找某个半径范围内所有成员
func (r *Redis) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return r.RedisClient.GeoRadius(context.Background(), key, longitude, latitude, query)
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
