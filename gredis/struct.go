package gredis

import "github.com/redis/go-redis/v9"

type Redis struct {
	RedisPasswd  string
	RedisAddress string
	RedisClient  *redis.Client
}

type RedisCluster struct {
	RedisPasswd  string
	RedisAddress []string
	RedisClient  *redis.ClusterClient
}

type RedisSentinel struct {
	RedisPasswd          string
	RedisMasterName      string
	RedisSentinelAddress []string
	RedisClient          *redis.Client
}
