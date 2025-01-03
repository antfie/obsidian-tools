package crypto

import (
	"github.com/btcsuite/btcd/btcutil/base58"
	"golang.org/x/crypto/blake2b"
	"io"
	"os"
	"path"
)

func HashDataToString(data []byte) (string, error) {
	hash, err := blake2b.New512([]byte{})

	if err != nil {
		return "", err
	}

	hash.Write(data)

	return base58.Encode(hash.Sum(nil)), err
}

// https://crypto.stackexchange.com/a/89559
// Apparently blake2b is faster than SHA-512 and 'collision resistant'
func HashFile(filePath string) (string, error) {
	file, err := os.Open(path.Clean(filePath))

	if err != nil {
		return "", err
	}

	hash, err := blake2b.New512([]byte{})

	if err != nil {
		return "", err
	}

	buffer := make([]byte, 4096)

	for {
		size, err := file.Read(buffer)

		if err != nil && err != io.EOF {
			return "", err
		}

		if err == io.EOF {
			break
		}

		hash.Write(buffer[0:size])
	}

	err = file.Close()

	if err != nil {
		return "", err
	}

	// For 64 bytes of BLAKE2b hash data, we expect 87 to 88 characters of Base-58
	return base58.Encode(hash.Sum(nil)), err
}
