package gredis

func NewRedisBase() RedisBase {
	return RedisBase{
		RedisPasswd: "redis",
	}
}

func NewRedis(opts ...RedisConfigOptions) (*Redis, error) {
	redisBase := NewRedisBase()
	result := &Redis{
		RedisAddress: "127.0.0.1:6379",
		RedisBase:    redisBase,
	}

	for _, opt := range opts {
		opt(result)
	}

	return result, nil
}

func NewRedisCluster(opts ...RedisClusterConfigOptions) (*RedisCluster, error) {
	redisBase := NewRedisBase()
	result := &RedisCluster{
		RedisBase:    redisBase,
		RedisAddress: []string{"127.0.0.1:6379"},
	}

	for _, opt := range opts {
		opt(result)
	}
	return result, nil
}

func NewRedisSentinel(opts ...RedisSentinelConfigOptions) (*RedisSentinel, error) {
	redisBase := NewRedisBase()
	result := &RedisSentinel{
		RedisBase:            redisBase,
		RedisMasterName:      "test",
		RedisSentinelAddress: []string{"127.0.0.1:6379"},
	}

	for _, opt := range opts {
		opt(result)
	}
	return result, nil
}

type RedisBaseConfigOptions func(*RedisBase)

type RedisConfigOptions func(*Redis)

type RedisClusterConfigOptions func(*RedisCluster)

type RedisSentinelConfigOptions func(*RedisSentinel)

func WithRedisBasePassword(password string) RedisBaseConfigOptions {
	return func(r *RedisBase) {
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
