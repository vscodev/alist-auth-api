package secrets

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// TokenBytes returns a secure random token of the specified number of bytes
func TokenBytes(n uint8) ([]byte, error) {
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// TokenHex returns a secure random token of the specified number of bytes, encoded as hex
func TokenHex(n uint8) (string, error) {
	b, err := TokenBytes(n)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// TokenBase64 returns a secure random token of the specified number of bytes, encoded as base64
func TokenBase64(n uint8) (string, error) {
	b, err := TokenBytes(n)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "="), nil
}
