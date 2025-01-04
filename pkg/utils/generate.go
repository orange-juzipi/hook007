package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	saltPassword = "ag7i5hI4sF2HJ3ihs3iioIh09"
)

// GeneratePassword 生成加密密码
func GeneratePassword(password string) string {
	m := md5.New()
	m.Write([]byte(password))
	mByte := m.Sum(nil)

	h := hmac.New(sha256.New, []byte(saltPassword))
	h.Write(mByte)
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateRandomString 生成随机数
func GenerateRandomString(n uint) string {
	const str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var build strings.Builder
	for i := 0; i < int(n); i++ {
		build.WriteByte(str[rand.IntN(len(str))])
	}
	return build.String()
}

// 生成订单号
func GenerateOrderID() string {
	uid, err := uuid.NewV7()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s-%s", time.Now().Format(time.DateOnly), strings.Split(uid.String(), "-")[4])
}
