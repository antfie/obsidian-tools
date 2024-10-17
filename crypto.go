package main

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512(data []byte) []byte {
	hash := sha512.New512_256()
	hash.Write(data)
	return hash.Sum(nil)
}

func Sha512ToString(data []byte) string {
	return hex.EncodeToString(Sha512(data))
}
