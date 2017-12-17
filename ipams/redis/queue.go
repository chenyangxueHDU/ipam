package redis

import "github.com/garyburd/redigo/redis"

func rpush(args ...interface{}) error {
	conn := getConn()
	defer conn.Close()

	_, err := conn.Do(`RPUSH`, args...)
	return err
}

func lpop(key string) (int, error) {
	conn := getConn()
	defer conn.Close()

	return redis.Int(conn.Do(`LPOP`, key))
}
