package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zihuyishi/joker/web/code"
	"github.com/zihuyishi/joker/web/model"
	"strconv"
)

func (r *Router) newTag(c *gin.Context) {
	name := c.PostForm("name")
	if len(name) == 0 {
		r.codeResponse(c, code.WrongParams)
		return
	}
	tag := &model.Tag{
		Name: name,
	}
	err := r.ctx.Dao.InsertTag(tag)
	if err != nil {
		r.codeResponse(c, code.DBError)
		return
	}
	r.jsonResponse(c, tag)
}

func (r *Router) addTagToJoker(c *gin.Context) {
	name := c.PostForm("name")
	strJokerId := c.PostForm("jokerId")
	jokerId, err := strconv.ParseInt(strJokerId, 10, 64)
	if err != nil || len(name) == 0 {
		r.codeResponse(c, code.WrongParams)
		return
	}

	err = r.ctx.Dao.AddNameTagToJoker(name, jokerId)
	if err != nil {
		fmt.Printf("add tag fail: %s", err.Error())
		r.codeResponse(c, code.DBError)
		return
	}
	r.codeResponse(c, code.Ok)
}

