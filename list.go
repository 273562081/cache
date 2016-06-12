package cache

import "github.com/garyburd/redigo/redis"

func (c *Connection) LIndex(key string, index int) (string, error) {
	return redis.String(c.Execute("LINDEX", key, index))
}

func (c *Connection) LLen(key string) (int, error) {
	return redis.Int(c.Execute("LLEN", key))
}

func (c *Connection) LPop(key string) (string, error) {
	return redis.String(c.Execute("LPOP", key))
}

func (c *Connection) LPush(key string, args ...interface{}) (int, error) {
	var _args []interface{}
	_args = append(_args, key)
	_args = append(_args, args...)
	return redis.Int(c.Execute("LPUSH", _args...))
}

func (c *Connection) LRange(key string, start, stop int) ([]string, error) {
	return redis.Strings(c.Execute("LRANGE", key, start, stop))
}

func (c *Connection) LSet(key string, index int, value string) (bool, error) {
	reply, err := redis.String(c.Execute("LSET", key, index, value))
	return IsReplyOk(reply), err
}
