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
	XDel(stream string, keys ...string) *redis.IntCmd
	XLen(stream string) *redis.IntCmd
	LPush(key string, value ...interface{}) *redis.IntCmd
	RPush(key string, value ...interface{}) *redis.IntCmd
	LPop(key string) *redis.StringCmd
	RPop(key string) *redis.StringCmd
	LLen(key string) *redis.IntCmd
	SAdd(key string, value interface{}) *redis.IntCmd
	SRem(key string, value interface{}) *redis.IntCmd
	SMembers(key string) *redis.StringSliceCmd
	SIsMember(key string, value interface{}) *redis.BoolCmd
	ZAdd(key string, members ...redis.Z) *redis.IntCmd
	ZRange(key string, start, stop int64) *redis.StringSliceCmd
	ZRank(key string, member string) *redis.IntCmd
	HGetKey(key string, field string) *redis.StringCmd
	HGetAllKey(key string) *redis.MapStringStringCmd
	HAddKey(key string, value interface{}) *redis.IntCmd
	HDelKey(stream string, keys ...string) *redis.IntCmd
	SetBit(key string, offset int64, value int) *redis.IntCmd
	GetBit(key string, offset int64) *redis.IntCmd
	PFAdd(key string, els ...interface{}) *redis.IntCmd
	PFCount(key string) *redis.IntCmd
	GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd
	GeoDist(key string, member1 string, member2 string, unit string) *redis.FloatCmd
	GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GetStreamsData(streamName string, groupName string, ss string) *redis.XStreamSliceCmd
	Ping()
}
