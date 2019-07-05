package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zihuyishi/joker/web/utils"
	"net/http"
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

	// joker
	{
		g.GET("/joker/:id", r.jokerById)
		g.PUT("/joker", r.newJoker)
		g.GET("/random", r.randomJoker)
	}


	// tag
	{
		g.PUT("/tag", r.newTag)
		g.POST("/tag/tojoker", r.addTagToJoker)
	}

	// user
	{
		g.POST("/login", r.login)
		g.POST("/logout", r.logout)
		g.GET("/user/:id", r.userById)
		g.GET("/user", r.currentUser)
		g.POST("/user/password", r.changePassword)
	}
}

func (r *Router) codeResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

func (r *Router) jsonResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}