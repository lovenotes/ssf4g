package crypto

import (
	"crypto/rc4"
)

func EccRc4(cryptokey []byte, cryptodata []byte) ([]byte, error) {
	c, err := rc4.NewCipher(cryptokey)

	if err != nil {
		return nil, err
	}

	destData := make([]byte, len(cryptodata))
	c.XORKeyStream(destData, cryptodata)

	return destData, nil
}
