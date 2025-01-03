package utils

import (
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	"strings"
)

const IDLength = 22

func NewID() string {
	uuidBytes := uuid.New()
	return b58(uuidBytes[:])
}

func b58(in []byte) string {
	val := base58.Encode(in)

	if len(val) == IDLength {
		return val
	}

	// Left pad with 1 as leading 0s are stripped
	// https://medium.com/concerning-pharo/understanding-base58-encoding-23e673e37ff6
	// https://gosamples.dev/string-padding/
	return strings.ReplaceAll(fmt.Sprintf("%22s", val), " ", "1")
}
