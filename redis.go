package gozzzproducer

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// RedisConn is redis connection struct
type RedisConn struct {
	conn *redis.Client
}

// NewRedisConn create redis connection
func NewRedisConn(address string, password string, db int) *RedisConn {
	redisConn := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return &RedisConn{
		conn: redisConn,
	}
}

// SetZSetValue set ZSet value and score
func (rc *RedisConn) SetZSetValue(key string, value string, score float64) (retErr error) {
	ctx := context.Background()
	zAdd := rc.conn.ZAdd(ctx, key, &redis.Z{Score: score, Member: value})
	if err := zAdd.Err(); err != nil {
		retErr = err
	}
	return
}

// GetHashValue get hash value
func (rc *RedisConn) GetHashValue(key string, field string) (value string, retErr error) {
	ctx := context.Background()
	hGet := rc.conn.HGet(ctx, key, field)
	if err := hGet.Err(); err != nil {
		retErr = err
		return
	}
	value = hGet.Val()
	return
}

// SetHashValue set Hash field and value
func (rc *RedisConn) SetHashValue(key string, field string, value string) (retErr error) {
	ctx := context.Background()
	hSet := rc.conn.HSet(ctx, key, field, value)
	if err := hSet.Err(); err != nil {
		retErr = err
	}
	return
}
