package cache

import "github.com/garyburd/redigo/redis"

func (c *Connection) FlushDb() error {
	_, err := c.Execute("FLUSHDB")
	return err
}

// FLUSHALL Command.
// @return error, returns non-nil error if error was reached.
func (c *Connection) FlushAll() error {
	_, err := c.Execute("FLUSHALL")
	return err
}

func (c *Connection) Save() (bool, error) {
	reply, err := redis.String(c.Execute("SAVE"))
	return IsReplyOk(reply), err
}
