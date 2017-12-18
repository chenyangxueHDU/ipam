package ipams

import "errors"

var (
	ErrPoolEmpty = errors.New(`No available ip address in the pool. `)
)
