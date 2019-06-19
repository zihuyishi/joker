package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zihuyishi/joker/web/code"
	"github.com/zihuyishi/joker/web/model"
	"net/http"
	"strconv"
)

type CodeResult struct {
	code int
}

func (r *Router) jokerById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		r.codeResponse(c, code.WrongParams)
		return
	}
	joker, err := r.ctx.Dao.FindJokerById(id)
	if err != nil {
		r.codeResponse(c, code.DBError)
		return
	}
	tags, err := r.ctx.Dao.FindJokerTags(joker.Id)
	if err == nil {
		joker.Tags = tags
	} else {
		fmt.Printf("get tags fail %s\n", err.Error())
	}
	c.JSON(http.StatusOK, joker)
}

func (r *Router) newJoker(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	if len(title) == 0 || len(content) == 0 {
		r.codeResponse(c, code.WrongParams)
		return
	}
	joker := &model.Joker{
		Title: title,
		Content: content,
	}
	err := r.ctx.Dao.InsertJoker(joker)
	if err != nil {
		r.codeResponse(c, code.DBError)
	}
	c.JSON(http.StatusOK, joker)
}

func (r *Router) codeResponse(c *gin.Context, code int) {
	res := &CodeResult{
		code,
	}
	c.JSON(http.StatusOK, res)
}

