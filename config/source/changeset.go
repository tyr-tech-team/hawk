package source

import (
	"crypto/md5"
	"fmt"
)

// Sum - returns the md5 checksum of the ChangeSet data
func Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
