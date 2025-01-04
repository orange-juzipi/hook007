package utils_test

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"hook007/pkg/utils"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	fmt.Println(utils.GeneratePassword("123456"))
}

func TestGenerateRandomString(t *testing.T) {
	fmt.Println(utils.GenerateRandomString(4))
}

func TestGenerateRandomNumber(t *testing.T) {
	openAccessKey, openSecret, err := utils.GenerateRandomAccessKeySecret(16, 16)
	if err != nil {
		fmt.Println("Error generating keys:", err)
		return
	}

	fmt.Printf("openAccessKey: %s\n", openAccessKey)
	fmt.Printf("openSecret: %s\n", openSecret)
}

func TestGenerateOrderID(t *testing.T) {
	fmt.Println(utils.GenerateOrderID())
}

func TestGenerateJwtSecret(t *testing.T) {
	secret, err := generateJWTSecret(32) // 生成 32 字节的随机 secret
	if err != nil {
		fmt.Println("Error generating secret:", err)
		return
	}
	fmt.Println("Generated JWT Secret:", secret)
}
func generateJWTSecret(length int) (string, error) {
	secret := make([]byte, length)
	_, err := rand.Read(secret)
	if err != nil {
		return "", err
	}

	// 返回 Base64 编码的字符串
	return base64.URLEncoding.EncodeToString(secret), nil
}
