package crypto

import (
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

func Sha1HashTime(data string) string {
	return Sha1Hash(time.Now().String() + data)
}

func Sha1Hash(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
