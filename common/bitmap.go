package common

func NewBitmap(key string, num int) error {
	conn := GetConn()
	defer conn.Close()

	_, err := conn.Do(`SETBIT`, key, 0, 1)
	return err
}
