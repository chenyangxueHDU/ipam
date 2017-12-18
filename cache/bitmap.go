package cache

func NewBitmap(key string) error {
	conn := GetConn()
	defer conn.Close()

	_, err := conn.Do(`SETBIT`, key, 0, 0)
	return err
}
