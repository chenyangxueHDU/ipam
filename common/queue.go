package common

import "github.com/garyburd/redigo/redis"

// Push key,val1,val2
func Push(args ...interface{}) error {
	conn := GetConn()
	defer conn.Close()

	_, err := conn.Do(`RPUSH`, args...)
	return err
}

func Pop(key string) (int, error) {
	conn := GetConn()
	defer conn.Close()

	return redis.Int(conn.Do(`LPOP`, key))
}
