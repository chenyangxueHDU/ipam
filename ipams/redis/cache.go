package redis

import "github.com/garyburd/redigo/redis"

func existsKey(key string) (bool, error) {
	conn := getConn()
	defer conn.Close()

	i, e := redis.Int(conn.Do(`EXISTS`, key))

	return i != 0, e
}

func set(key, value string) error {
	conn := getConn()
	defer conn.Close()

	_, err := conn.Do(`SET`, key, value)
	return err
}

func get(key string) (string, error) {
	conn := getConn()
	defer conn.Close()

	return redis.String(conn.Do(`GET`, key))
}
