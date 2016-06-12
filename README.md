# RedisCache
RedisCache is redis command package for Go.


# Documentation
## Installation
```
go get github.com/clevergo/cache
```

## Example
```
package main

import (
	"github.com/clevergo/cache"
	"fmt"
)

func main() {
    // Create a redis pool.
	pool := cache.NewRedisPool(100, 300, "tcp", ":6379", "", "")
	defer pool.Close()

    // Create RedisCache
	rc := cache.NewRedisCache(pool)

    var reply interface{}
    var err error

    // Get connection.
    conn := rc.GetConn()
    defer conn.Close()
    
	reply, err = conn.Execute("SET", "name", "RedisCache")
    fmt.Printf("SET: %v,%v\n", reply, err)
    
    reply, err = conn.Get("name")
    fmt.Printf("GET: %v,%v\n", reply, err)
    
    // Of course, you can do it like the following
    reply, err = conn.Execute("GET", "name")
    fmt.Printf("GET: %v,%v\n", reply, err)
}
```


# Supported Command
RedisCache supports most of redis's commands, and it is convenient to execute the other command by invoking the **Execute()** method.

## Server Command
- FLUSHDB
- FLUSHALL
- SAVE

## Connection Command
- AUTH
- ECHO
- PING
- QUIT
- SELECT

## Key Command
- DEL
- EXISTS
- EXPIRE
- EXPIREAT
- KEYS
- MOVE
- PERSIST
- PEXPIRE
- PEXPIREAT
- PTTL
- RANDOMKEY
- RENAME
- RENAMENX
- TTL
- TYPE

## String Command
- APPEND
- GET
- GETRANGE
- GETSET
- INCR
- MGET
- MSET
- SET
- SETRANGE
- STRLEN

## Hash Command
- HDEL
- HEXISTS
- HGET
- HGETALL
- HINCRBY
- HINCRBYFLOAT
- HKEYS
- HLEN
- HMGET
- HMSET
- HSET
- HSETNX
- HVALS

## List Command
- LINDEX
- LLEN
- LPOP
- LPUSH
- LRANGE
- LSET

## Set Command
- SADD
- SCARD
- SDIFF
- SDIFFSTORE
- SINTER
- SINTERSTORE
- SISMEMBER
- SMEMBERS
- SMOVE
- SPOP
- SRANDMEMBER
- SREM
- SUNION
- SUNIONSTORE

## Transaction Command
- DISCARD
- EXEC
- MULTI
- WATCH
- UNWATCH
