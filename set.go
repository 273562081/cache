package cache

import "github.com/garyburd/redigo/redis"

func (c *Connection) SAdd(key string, members ...interface{}) (int, error) {
	args := []interface{}{key}
	args = append(args, members...)
	return redis.Int(c.Execute("SADD", args...))
}

func (c *Connection) SCard(key string) (int, error) {
	return redis.Int(c.Execute("SCARD", key))
}

func (c *Connection) SDiff(keys ...interface{}) ([]string, error) {
	return redis.Strings(c.Execute("SDIFF", keys...))
}

func (c *Connection) SDiffStore(destination string, keys ...interface{}) ([]string, error) {
	args := []interface{}{destination}
	args = append(args, keys...)
	return redis.Strings(c.Execute("SDIFFSTORE", args...))
}

func (c *Connection) SInter(keys ...interface{}) ([]string, error) {
	return redis.Strings(c.Execute("SINTER", keys...))
}

func (c *Connection) SInterStore(destination string, keys ...interface{}) ([]string, error) {
	args := []interface{}{destination}
	args = append(args, keys...)
	return redis.Strings(c.Execute("SINTERSTORE", args...))
}

func (c *Connection) SIsMember(key string, member string) (bool, error) {
	reply, err := redis.String(c.Execute("SISMEMBER", key, member))
	return IsReplyOk(reply), err
}

func (c *Connection) SMembers(key string) ([]string, error) {
	return redis.Strings(c.Execute("SMEMBERS", key))
}

func (c *Connection) SMove(source, destination, member string) (bool, error) {
	reply, err := redis.Int(c.Execute("SMOVE", source, destination, member))
	return (reply == 1), err
}

func (c *Connection) SPop(key string) (string, error) {
	return redis.String(c.Execute("SPOP", key))
}

func (c *Connection) SRandMember(key string) (string, error) {
	return redis.String(c.Execute("SRANDMEMBER", key))
}

func (c *Connection) SRandMembers(key string, count int) ([]string, error) {
	return redis.Strings(c.Execute("SRANDMEMBER", key, count))
}

func (c *Connection) SReM(key string, members ...interface{}) (int, error) {
	args := []interface{}{key}
	args = append(args, members...)
	return redis.Int(c.Execute("SREM", args))
}

func (c *Connection) SUnion(keys ...interface{}) ([]string, error) {
	return redis.Strings(c.Execute("SUNION", keys...))
}

func (c *Connection) SUnionStore(destination string, keys ...interface{}) ([]string, error) {
	args := []interface{}{destination}
	args = append(args, keys...)
	return redis.Strings(c.Execute("SUNIONSTORE", args...))
}
