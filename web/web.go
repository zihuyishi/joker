package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/zihuyishi/joker/web/controller"
	"github.com/zihuyishi/joker/web/dao"
	"github.com/zihuyishi/joker/web/utils"
)

type Web struct {
	config *utils.Config
}

func New(config *utils.Config) *Web {
	web := Web {
		config,
	}
	return &web
}

func (web *Web) Serve() {
	config := web.config
	pgConfig := &pg.Options{
		Addr: config.DBAddr,
		User: config.DBUser,
		Password: config.DBPass,
		Database: config.DBName,
	}
	d := dao.New(pgConfig)
	g := gin.Default()
	g.Static("/public", "./public")
	ctx := utils.Context{
		Dao: d,
		G:   g,
	}
	router := controller.New(&ctx)
	router.LoadRoutes()
	g.Run(web.config.Addr)
}