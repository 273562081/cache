package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// APPEND Command.
// @param string key.
// @param string value.
// @return string value. Returns the length of value,
// The key's value will be set as value if it does not exist.
func (c *Connection) Append(key, value string) (int, error) {
	return redis.Int(c.Execute("APPEND", key, value))
}

// GET Command.
// @return string value.
// Non-nil error will be returned, if the value's type is not string.
// If the key does not exist, returns nil and nil error.
func (c *Connection) Get(key string) (string, error) {
	return redis.String(c.Execute("GET", key))
}

// GETRANGE Command.
// @param string key.
// @param int start.
// @param int end.
// @return substring of the key's value.
func (c *Connection) GetRange(key string, start, end int) (string, error) {
	return redis.String(c.Execute("GETRANGE", key, start, end))
}

// GETSET Command.
// @param string key.
// @param string value the new value.
// @return the old value, ErrNil will be return if the key does not exist.
func (c *Connection) GetSet(key, value string) (string, error) {
	return redis.String(c.Execute("GETSET", key, value))
}

// INCR Command.
// @param string key.
// @return int64 and error.
// Key's value will be set as zero and return if the key does not exist.
func (c *Connection) Incr(key string) (int64, error) {
	return redis.Int64(c.Execute("INCR", key))
}

// MGET Command.
func (c *Connection) MGet(keys ...interface{}) ([]string, error) {
	return redis.Strings(c.Execute("MGET", keys))
}

// MSET Command.
func (c *Connection) MSet(args ...interface{}) error {
	_, err := c.Execute("MSET", args)
	return err
}

// SET Command.
// @param string key.
// @param string value.
// @param []string args, args[0] can be set as 'EX seconds', 'PX milliseconds', 'NX' OR 'XX'.
// @return boolean and error.
func (c *Connection) Set(key, value string, args ...[]string) (bool, error) {
	var reply string
	var err error

	if len(args) == 0 {
		reply, err = redis.String(c.Execute("SET", key, value))
	} else {
		reply, err = redis.String(c.Execute("SET", key, value, args[0]))
	}

	if err != nil {
		return false, err
	}

	if IsReplyOk(reply) && (err == nil) {
		return true, nil
	}

	return false, fmt.Errorf(reply)
}

// SETRANGE Command.
// @param string key.
// @param int offset.
// @param string value.
// @return the length of modified key's value.
func (c *Connection) SetRange(key string, offset int, value string) (int64, error) {
	return redis.Int64(c.Execute("SETRANGE", key, offset, value))
}

// STRLEN Command.
// @return length of string value.
// Non-nil error will be returned, if the value's type is not string.
// If the key does not exist, returns zero and nil error.
func (c *Connection) Strlen(key string) (int, error) {
	return redis.Int(c.Execute("STRLEN", key))
}
