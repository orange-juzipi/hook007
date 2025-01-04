package token_test

import (
	"fmt"
	"hook007/config"
	"hook007/pkg/token"
	"testing"
)

func TestMain(m *testing.M) {
	config.ConfigPath = "../../"
	config.Init()

	m.Run()
}

func TestXxx(t *testing.T) {
	a := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjAsIlJvbGUiOiJxdyIsImV4cCI6MTcxODU3NzkzMiwibmJmIjoxNzE4NDA1MTMyLCJpYXQiOjE3MTg0MDUxMzJ9.Nex8k1h2u7OGjNFlBhK_zPN-Gr5lNZh0XUrjUUT7lFM"
	str := token.New(config.Get().JWT.Secret)
	fmt.Println(str.JwtParse(a))
}
