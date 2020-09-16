package gozzzproducer

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// RedisConn is redis connection struct
type RedisConn struct {
	conn *redis.Client
	ctx  context.Context
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
		ctx:  context.Background(),
	}
}

// SetZSetValue set ZSet value and score
func (rc *RedisConn) SetZSetValue(key string, value string, score float64) (retErr error) {
	zAdd := rc.conn.ZAdd(rc.ctx, key, &redis.Z{Score: score, Member: value})
	if err := zAdd.Err(); err != nil {
		retErr = err
	}
	return
}

// SetHashValue set Hash field and value
func (rc *RedisConn) SetHashValue(key string, field string, value string) (retErr error) {
	hSet := rc.conn.HSet(rc.ctx, key, field, value)
	if err := hSet.Err(); err != nil {
		retErr = err
	}
	return
}
