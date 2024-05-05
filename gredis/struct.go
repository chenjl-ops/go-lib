package gredis

import "github.com/redis/go-redis/v9"

type RedisBase struct {
	RedisPasswd string
}

type Redis struct {
	RedisBase
	RedisAddress string
	RedisClient  *redis.Client
}

type RedisCluster struct {
	RedisBase
	RedisAddress []string
	RedisClient  *redis.ClusterClient
}

type RedisSentinel struct {
	RedisBase
	RedisMasterName      string
	RedisSentinelAddress []string
	RedisClient          *redis.Client
}
