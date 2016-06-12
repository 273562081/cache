package cache

import (
	"github.com/garyburd/redigo/redis"
)

// DEL Command.
// @param []interface{} keys which will be deleted.
// @return the count of deleted keys.
func (c *Connection) Del(keys ...interface{}) (int, error) {
	delCount, err := redis.Int(c.Execute("DEL", keys...))

	if err != nil {
		return 0, err
	}
	return delCount, nil
}

// EXISTS Command.
// @param string key.
// @return boolean and error indicating whether the key is existed.
func (c *Connection) Exists(key string) (bool, error) {
	reply, err := redis.Int(c.Execute("EXISTS", key))
	return (reply == 1), err
}

// EXPIRE Command.
// @param string key.
// @param uint seconds, zero means delete the key.
// @return boolean and error.
func (c *Connection) Expire(key string, seconds uint) (bool, error) {
	reply, err := redis.Int(c.Execute("EXPIRE", key, seconds))
	return (reply == 1), err
}

// EXPIREAT Command.
// @param string key.
// @param int timestamp, zero or negative number means delete the key.
// @return boolean and error.
func (c *Connection) ExpireAt(key string, timestamp int) (bool, error) {
	reply, err := redis.Int(c.Execute("EXPIREAT", key, timestamp))
	return (reply == 1), err
}

// KEYS Command.
func (c *Connection) Keys(pattern string) ([]string, error) {
	return redis.Strings(c.Execute("EXPIREAT", pattern))
}

// MOVE Command.
// @param string key.
// @param string db.
// @return boolean and error.
func (c *Connection) Move(key string, db string) (bool, error) {
	reply, err := redis.Int(c.Execute("MOVE", key, db))
	return (reply == 1), err
}

// PERSIST Command.
// @param string key.
// @return boolean and error.
func (c *Connection) Persist(key string) (bool, error) {
	reply, err := redis.Int(c.Execute("PERSIST", key))
	return (reply == 1), err
}

// PEXPIRE Command.
// @param string key.
// @param uint64 milliseconds, zero means delete the key.
// @return boolean and error.
func (c *Connection) PExpire(key string, milliseconds uint64) (bool, error) {
	reply, err := redis.Int(c.Execute("PEXPIRE", key, milliseconds))
	return (reply == 1), err
}

// PEXPIREAT Command.
// @param string key.
// @param int64 milliseconds timestamp, zero or negative number means delete the key.
// @return boolean and error.
func (c *Connection) PExpireAt(key string, timestamp int64) (bool, error) {
	reply, err := redis.Int(c.Execute("PEXPIREAT", key, timestamp))
	return (reply == 1), err
}

// PTTL Command.
// @param string key.
// @return error will be returned when error was reached.
// In general, returns negative one, negative two or positive number,
// negative one means that no specify time to live,
// negative two means that the key does not exist.
func (c *Connection) PTtl(key string) (int64, error) {
	return redis.Int64(c.Execute("PTTL", key))
}

// RANDOMKEY Command.
// @return random key and error.
func (c *Connection) RandomKey() (string, error) {
	return redis.String(c.Execute("RANDOMKEY"))
}

// RENAME Command.
// @param string key old key's name.
// @param string newKey new key's name, the new key will be replaced if it exists.
// @return boolean and error.
func (c *Connection) Rename(key, newkey string) (bool, error) {
	reply, err := redis.String(c.Execute("RENAME", key, newkey))
	return IsReplyOk(reply), err
}

// RENAMENX Command.
// @param string key old key's name.
// @param string newKey new key's name.
// @return boolean and error, error will be reached if the new key exists.
func (c *Connection) RenameNX(key, newkey string) (bool, error) {
	reply, err := redis.String(c.Execute("RENAMENX", key, newkey))
	return IsReplyOk(reply), err
}

// TTL Command.
// @param string key.
// @return error will be returned when error was reached.
// In general, returns negative one, negative two or positive number,
// negative one means that no specify time to live,
// negative two means that the key does not exist.
func (c *Connection) Ttl(key string) (int, error) {
	return redis.Int(c.Execute("TTL", key))
}

// TYPE Command.
// @param string key.
// @return the type of key's value.
// In general, returns 'none','string','list','set' or 'hash'.
// Empty string and error will be returned when error was reached.
func (c *Connection) TYPE(key string) (string, error) {
	return redis.String(c.Execute("TYPE", key))
}
