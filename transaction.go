package cache

import "github.com/garyburd/redigo/redis"

func (c *Connection) Discard() error {
	_, err := c.Execute("DISCARD")
	return err
}

func (c *Connection) Exec() ([]interface{}, error) {
	return redis.Values(c.Execute("EXEC"))
}

func (c *Connection) Multi() error {
	_, err := c.Execute("MULTI")
	return err
}

func (c *Connection) Watch() error {
	_, err := c.Execute("WATCH")
	return err
}

func (c *Connection) UnWatch() error {
	_, err := c.Execute("UNWATCH")
	return err
}
