package cache

import (
	"github.com/garyburd/redigo/redis"
)

// HDEL Command.
func (c *Connection) HDel(key string, fields ...interface{}) (int, error) {
	return redis.Int(c.Execute("HDEL", fields...))
}

// HEXISTS Command.
func (c *Connection) HExists(key, field string) (bool, error) {
	reply, err := redis.Int(c.Execute("HEXISTS", key, field))
	return (reply == 1), err
}

// HGET Command.
func (c *Connection) HGet(key, field string) (string, error) {
	return redis.String(c.Execute("HGET", key, field))
}

// HGETALL Command.
func (c *Connection) HGetAll(key string) (map[string]string, error) {
	return redis.StringMap(c.Execute("HGETALL", key))
}

// HINCRBY Command.
func (c *Connection) HIncrBy(key, field string, increment int64) (int64, error) {
	return redis.Int64(c.Execute("HINCRBY", key, field, increment))
}

// HINCRBYFLOAT Command.
func (c *Connection) HIncrByFloat(key, field string, increment float64) (float64, error) {
	return redis.Float64(c.Execute("HINCRBYFLOAT", key, field, increment))
}

// HKEYS Command.
func (c *Connection) HKeys(key string) ([]string, error) {
	return redis.Strings(c.Execute("HKEYS", key))
}

// HLEN Command.
func (c *Connection) HLen(key string) (int, error) {
	return redis.Int(c.Execute("HLEN", key))
}

// HMGET Command.
func (c *Connection) HMGet(key string, args ...interface{}) ([]string, error) {
	var _args []interface{}
	_args = append(_args, key)
	_args = append(_args, args...)
	return redis.Strings(c.Execute("HMGET", _args...))
}

// HMSET Command.
func (c *Connection) HMSet(key string, args ...interface{}) (bool, error) {
	var _args []interface{}
	_args = append(_args, key)
	_args = append(_args, args...)
	reply, err := redis.String(c.Execute("HMSET", _args...))
	return IsReplyOk(reply), err
}

// HSET Command.
func (c *Connection) HSet(key, field, value string) (int, error) {
	return redis.Int(c.Execute("HSET", key, field, value))

}

// HSETNX Command.
func (c *Connection) HSetNX(key, field, value string) (bool, error) {
	reply, err := redis.Int(c.Execute("HSETNX", key, field, value))
	return (reply == 1), err
}

// HVALS Command.
func (c *Connection) HVals(key string) ([]string, error) {
	return redis.Strings(c.Execute("HVALS", key))
}
