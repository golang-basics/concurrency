package digestion

import (
	"crypto/md5"
)

type MD5Result map[string][md5.Size]byte

// result represents the product of reading and summing a file using MD5.
type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}
