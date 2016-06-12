package cache

import (
	"github.com/garyburd/redigo/redis"
	"strings"
)

// SELECT Command. Select database.
// @return error, returns non-nil error if error was reached.
func (c *Connection) Select(db string) error {
	_, err := c.Execute("SELECT", db)
	return err
}

// AUTH Command.
// @param string password.
// @return boolean and error, returns non-nil error if error was reached.
// Returns true if the password is correct.
func (c *Connection) Auth(password string) (bool, error) {
	reply, err := redis.String(c.Execute("AUTH", password))
	return IsReplyOk(reply), err
}

// ECHO Command.
// @param string message.
// @return string and error, returns non-nil error if error was reached.
// Returns the message normally.
func (c *Connection) Echo(message string) (string, error) {
	return redis.String(c.Execute("ECHO", message))
}

// PING Command.
// @return boolean and error, return true normally.
// Returns non-nil error if error was reached.
func (c *Connection) Ping() (bool, error) {
	reply, err := redis.String(c.Execute("PING"))
	return (0 == strings.Compare("PONG", reply)), err
}

// QUIT Command.
func (c *Connection) Quit() {
	c.Execute("QUIT")
}
