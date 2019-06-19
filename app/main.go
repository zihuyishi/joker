package main

import (
	"github.com/zihuyishi/joker/web"
	"github.com/zihuyishi/joker/web/utils"
)

func main() {
	config := &utils.Config{
		DBPass: "mysecretpassword",
		DBUser: "postgres",
		DBName: "joker",
		Addr: ":7001",

	}
	server := web.New(config)
	server.Serve()
}
