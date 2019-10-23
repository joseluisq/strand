package strand

import (
	"crypto/rand"
	"fmt"
	"math"
)

// RandomBytes generates secure random bytes
func RandomBytes(len int) ([]byte, error) {
	if len <= 0 {
		return []byte(nil), nil
	}

	size := math.Ceil(float64(len) / 2.0)
	byts := make([]byte, int(size))

	_, err := rand.Read(byts)

	if err != nil {
		return nil, err
	}

	return byts, nil
}

// RandomString generates a secure random string
func RandomString(len int) (str string, err error) {
	if len <= 0 {
		return "", nil
	}

	byts, err := RandomBytes(len)

	if err != nil {
		return "", err
	}

	str = fmt.Sprintf("%x", byts)

	return str, nil
}
