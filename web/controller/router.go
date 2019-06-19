package controller

import (
	"github.com/zihuyishi/joker/web/utils"
)

type Router struct {
	ctx *utils.Context
}

func New(ctx *utils.Context) *Router {
	r := Router {
		ctx,
	}
	return &r
}

func (r *Router) LoadRoutes() {
	g := r.ctx.G
	g.GET("/joker/:id", r.jokerById)
	g.PUT("/joker", r.newJoker)
}