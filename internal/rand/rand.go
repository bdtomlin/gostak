package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	bytesRead, err := rand.Read(bytes)
	if err != nil {
		return nil, fmt.Errorf("Bytes: %w", err)
	}
	if bytesRead != n {
		return nil, fmt.Errorf("Bytes: %w", err)
	}

	return bytes, nil
}

func String(numBytes int) (string, error) {
	bytes, err := Bytes(numBytes)
	if err != nil {
		return "", fmt.Errorf("String: %w", err)
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
