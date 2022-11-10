package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterBytes     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIndexBits = 6                      // 6 bits to represent a letter index
	letterIndexMask = 1<<letterIndexBits - 1 // All 1-bits, as many as letterIdxBits
	letterIndexMax  = 63 / letterIndexBits   // # of letter indices fitting in 63 bits
)

func GetRandomString(size int) string {
	src := rand.NewSource(time.Now().UnixNano())
	bytes := make([]byte, size)

	for bytesIndex, cache, remain := size-1, src.Int63(), letterIndexMax; bytesIndex >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIndexMax
		}
		if letterIndex := int(cache & letterIndexMask); letterIndex < len(letterBytes) {
			bytes[bytesIndex] = letterBytes[letterIndex]
			bytesIndex--
		}

		cache >>= letterIndexBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&bytes))
}
