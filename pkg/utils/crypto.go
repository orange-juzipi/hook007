package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// generateRandomBytes 生成指定长度的随机字节
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomAccessKeySecret 生成随机的 openAccessKey 和 openSecret
func GenerateRandomAccessKeySecret(keyLength, secretLength int) (string, string, error) {
	keyBytes, err := generateRandomBytes(keyLength)
	if err != nil {
		return "", "", err
	}
	key := base64.URLEncoding.EncodeToString(keyBytes)

	secretBytes, err := generateRandomBytes(secretLength)
	if err != nil {
		return "", "", err
	}
	secret := base64.URLEncoding.EncodeToString(secretBytes)

	return key, secret, nil
}
