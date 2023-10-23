package service

import (
	cRand "crypto/rand"
	"encoding/binary"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Base62Encode(number uint64) string {
	length := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}

func GetShortURL() string {
	b := make([]byte, 16)
	cRand.Read(b)
	return Base62Encode(binary.BigEndian.Uint64(b))
}
