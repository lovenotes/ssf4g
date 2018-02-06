package crypto

import (
	"crypto/rc4"
)

// Func - 使用EccRc4加密数据
func EncryptRc4(data []byte, key []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)

	if err != nil {
		return nil, err
	}

	destData := make([]byte, len(data))
	c.XORKeyStream(destData, data)

	return destData, nil
}
