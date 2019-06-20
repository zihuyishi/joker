package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zihuyishi/joker/web/code"
	"github.com/zihuyishi/joker/web/model"
)

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
	r.jsonResponse(c, joker)
}

func (r *Router) newJoker(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	if len(title) == 0 || len(content) == 0 {
		r.codeResponse(c, code.WrongParams)
		return
	}
	joker := &model.Joker{
		Title:   title,
		Content: content,
	}
	err := r.ctx.Dao.InsertJoker(joker)
	if err != nil {
		r.codeResponse(c, code.DBError)
		return
	}
	r.jsonResponse(c, joker)
}

func (r *Router) randomJoker(c *gin.Context) {
	strCount := c.DefaultQuery("count", "5")
	count, err := strconv.Atoi(strCount)
	if err != nil {
		r.codeResponse(c, code.WrongParams)
		return
	}
	jokers, err := r.ctx.Dao.RandomJoker(count)
	if err != nil {
		fmt.Printf("random joker fail: %s", err.Error())
		r.codeResponse(c, code.DBError)
		return
	}
	r.jsonResponse(c, jokers)
}

func (r *Router) batchNewJoker(c *gin.Context) {
	contents := c.PostFormArray("contents")
	for i := 0; i < len(contents); i++ {
		joker := &model.Joker{
			Title: "none",
			Content: contents[i],
		}
		err := r.ctx.Dao.InsertJoker(joker)
		if err != nil {
			r.codeResponse(c, code.DBError)
			return
		}
	}
	r.codeResponse(c, code.Ok)
}
