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
		DBAddr: "some-postgres:5432",
		Addr: ":7001",

	}
	server := web.New(config)
	server.Serve()
}
