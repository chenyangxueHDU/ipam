package redis

import (
	"testing"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func TestGetConn(t *testing.T) {
	conn := getConn()
	if conn == nil {
		t.Fail()
	}

	conn.Do(`set`, `hello`, `world`)

	s, e := redis.String(conn.Do(`get`, `hello`))

	fmt.Printf(`%s,%v`, s, e)
}