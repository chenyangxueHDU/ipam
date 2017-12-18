package common

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

	conn.Do(`set`, `d`, []byte{1,0,0})

	s, e := redis.String(conn.Do(`get`, `d`))

	fmt.Printf(`%s,%v`, s, e)
}