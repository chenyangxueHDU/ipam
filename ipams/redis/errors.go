package redis

import "errors"

var (
	ErrKeyNotExists = errors.New(`The redis key is not exists. `)
)
