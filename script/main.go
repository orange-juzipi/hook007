package main

import (
	"hook007/config"
)

func main() {
	config.ConfigPath = "../"
	config.Init()

	if err := ConnectGen(config.ConfigPath); err != nil {
		panic(err)
	}
}
