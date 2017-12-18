package common

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", `192.168.187.130:6379`)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetConn() redis.Conn {
	conn := pool.Get()
	if conn.Err() != nil {
		return nil
	}
	return conn
}

func ExistsKey(key string) (bool, error) {
	conn := GetConn()
	defer conn.Close()

	i, e := redis.Int(conn.Do(`EXISTS`, key))

	return i != 0, e
}
