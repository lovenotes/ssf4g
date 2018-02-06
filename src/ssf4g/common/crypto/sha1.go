package crypto

import (
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

// Func - 使用Sha1Hash加密含时间信息的数据
func EncryptSha1HashTime(data string) string {
	return EncryptSha1Hash(time.Now().String() + data)
}

// Func - 使用Sha1Hash加密数据
func EncryptSha1Hash(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
