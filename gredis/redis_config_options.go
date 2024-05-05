package gredis

const redisDefaultsPasswd = "redis"

func NewRedis(opts ...RedisConfigOptions) (*Redis, error) {
	result := &Redis{
		RedisAddress: "127.0.0.1:6379",
		RedisPasswd:  redisDefaultsPasswd,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result, nil
}

func NewRedisCluster(opts ...RedisClusterConfigOptions) (*RedisCluster, error) {
	result := &RedisCluster{
		RedisPasswd:  redisDefaultsPasswd,
		RedisAddress: []string{"127.0.0.1:6379"},
	}

	for _, opt := range opts {
		opt(result)
	}
	return result, nil
}

func NewRedisSentinel(opts ...RedisSentinelConfigOptions) (*RedisSentinel, error) {
	result := &RedisSentinel{
		RedisPasswd:          redisDefaultsPasswd,
		RedisMasterName:      "test",
		RedisSentinelAddress: []string{"127.0.0.1:6379"},
	}

	for _, opt := range opts {
		opt(result)
	}
	return result, nil
}

type RedisConfigOptions func(*Redis)

type RedisClusterConfigOptions func(*RedisCluster)

type RedisSentinelConfigOptions func(*RedisSentinel)

func WithRedisPassword(password string) RedisConfigOptions {
	return func(r *Redis) {
		r.RedisPasswd = password
	}
}

func WithRedisClusterPassword(password string) RedisClusterConfigOptions {
	return func(r *RedisCluster) {
		r.RedisPasswd = password
	}
}

func WithRedisSentinelPassword(password string) RedisSentinelConfigOptions {
	return func(r *RedisSentinel) {
		r.RedisPasswd = password
	}
}

func WithRedisAddr(addr string) RedisConfigOptions {
	return func(r *Redis) {
		r.RedisAddress = addr
	}
}

func WithRedisClusterAddr(addr []string) RedisClusterConfigOptions {
	return func(r *RedisCluster) {
		r.RedisAddress = addr
	}
}

func WithRedisSentinelAddr(addr []string) RedisSentinelConfigOptions {
	return func(r *RedisSentinel) {
		r.RedisSentinelAddress = addr
	}
}

func WithRedisSentinelMasterName(name string) RedisSentinelConfigOptions {
	return func(r *RedisSentinel) {
		r.RedisMasterName = name
	}
}
