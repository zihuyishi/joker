package main

import (
	"fmt"
	"os"
	"github.com/zihuyishi/joker/web"
	"github.com/zihuyishi/joker/web/utils"
)

func main() {
	path, exists := os.LookupEnv("JOKER_CONFIG")
	if !exists {
		path = "./config/config.json"
	}
	config, err := utils.LoadConfigFromFile(path)
	if err != nil {
		fmt.Printf("read config error: %s", err.Error())
		return
	}

	server := web.New(config)
	server.Serve()
}
