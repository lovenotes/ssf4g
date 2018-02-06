package utility

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Func - 随机[0,n)一个整数
func RandNum(n int) int {
	if n <= 0 {
		return 0
	}

	return rand.Intn(n)
}

// Func - 随机[base,base+n]一个整数
func RandBaseNum(base int, n int) int {
	if n < 0 {
		return base
	}

	return base + rand.Intn(n+1)
}
