package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/zihuyishi/joker/web/controller"
	"github.com/zihuyishi/joker/web/dao"
	"github.com/zihuyishi/joker/web/utils"
	"log"
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

	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("jokersession", store))

	router := controller.New(&ctx)
	router.LoadRoutes()
	err := g.Run(web.config.Addr)
	if err != nil {
		log.Fatalf("start fail: %s", err.Error())
	}
}