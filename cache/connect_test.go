package cache

import (
	"testing"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func TestGetConn(t *testing.T) {
	conn := GetConn()
	if conn == nil {
		t.Fail()
	}

	conn.Do(`set`, `d`, 1<<2)

	s, e := redis.String(conn.Do(`get`, `d`))

	fmt.Printf(`%s,%v`, s, e)
}
