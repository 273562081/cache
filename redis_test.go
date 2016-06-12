package cache_test

import (
	"github.com/clevergo/cache"
	"github.com/garyburd/redigo/redis"
	"strings"
	"testing"
)

var (
	keyPrefix   = "rediscache_test_"
	maxIdle     = 100
	idleTimeout = 300
	network     = "tcp"
	address     = ":6379"
	password    = ""
	db          = ""
)

func TestAll(t *testing.T) {
	pool := cache.NewRedisPool(maxIdle, idleTimeout, network, address, password, db)
	defer pool.Close()

	c := cache.NewRedisCache(pool)

	conn := c.GetConn()

	err := conn.FlushAll()
	if err != nil {
		t.Error("Failed to FLUSHALL: %s", err.Error())
	}

	var key = keyPrefix + "string_key"

	// GET command.
	getReply, err := conn.Get(key)
	if (err == nil) || (err != redis.ErrNil) {
		t.Error("Get a nonexistent key, unexpected result: %s, %v", getReply, err)
	}

	// Set Command.
	ok, err := conn.Set(key, "string_key")
	if err != nil {
		t.Error("Set command error: %s, %v", ok, err)
	}

	getReply, err = conn.Get(key)
	if 0 != strings.Compare("string_key", getReply) {
		t.Error("Get a nonexistent key, unexpected result: %s, %v", getReply, err)
	}
}
