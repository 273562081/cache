package cache

import (
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

// RedisCache
type RedisCache struct {
	pool *redis.Pool
}

type Connection struct {
	redis.Conn
}

// New Redis Pool
func NewRedisPool(maxIdle, idleTimeout int, network, address, password, db string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return dial(network, address, password, db)
		},
	}
}

func dial(network, address, password, db string) (redis.Conn, error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, err := c.Do("AUTH", password); err != nil {
			c.Close()
			return nil, err
		}
	}
	if len(db) > 0 {
		if _, err := c.Do("SELECT", db); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, err
}

// Create redis cache by redis'pool.
func NewRedisCache(pool *redis.Pool) *RedisCache {
	return &RedisCache{
		pool: pool,
	}
}

// Get Redis Pool
func (rc *RedisCache) GetPool() *redis.Pool {
	return rc.pool
}

// Get redis connection from pool.
func (rc *RedisCache) GetConn() *Connection {
	return &Connection{rc.pool.Get()}
}

// Execute Command.
func (c *Connection) Execute(command string, args ...interface{}) (interface{}, error) {
	return c.Do(command, args...)
}

// Returns boolean indicating whether the reply is equal to 'OK'
func IsReplyOk(v string) bool {
	return (0 == strings.Compare("OK", v))
}
