package vars

import "errors"

var (
	ErrAlreadyExist = errors.New("already exist")
	JwtSecret = []byte("kefteme")
)
