package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zihuyishi/joker/web/code"
	"github.com/zihuyishi/joker/web/model"
	"strconv"
)

func (r *Router) login(c *gin.Context) {
	var msg model.User
	err := c.ShouldBind(&msg)
	if err != nil {
		r.codeResponse(c, code.WrongParams)
		return
	}

	user, err := r.ctx.Dao.VerifyUserPassword(msg.Name, msg.Password)
	if err != nil {
		fmt.Printf("verify password error: %s", err.Error())
		r.codeResponse(c, code.WrongPassword)
		return
	}

	session := sessions.Default(c)
	session.Set("uid", user.Id)
	err = session.Save()
	if err != nil {
		r.codeResponse(c, code.DBError)
		return
	}
	user.Password = ""
	r.jsonResponse(c, user)
}

func (r *Router) logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		r.codeResponse(c, code.DBError)
		return
	}
	r.codeResponse(c, code.Ok)
}

func (r *Router) userById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		r.codeResponse(c, code.WrongParams)
		return
	}
	user, err := r.ctx.Dao.FindUserById(id)
	if err != nil {
		r.codeResponse(c, code.DBError)
		return
	}
	user.Password = ""
	r.jsonResponse(c, user)
}

func (r *Router) currentUser(c *gin.Context) {
	session := sessions.Default(c)
	var uid int64
	v := session.Get("uid")
	if v == nil {
		r.codeResponse(c, code.USER_NOT_LOGIN)
		return
	} else {
		uid = v.(int64)
	}
	user, err := r.ctx.Dao.FindUserById(uid)
	if err != nil {
		r.codeResponse(c, code.DBError)
		return
	}
	user.Password = ""
	r.jsonResponse(c, user)
}

type ChangePasswordMsg struct {
	Password string `form:"password"`
	Newpassword string `form:"newpassword"`
}

func (r *Router) changePassword(c *gin.Context) {
	session := sessions.Default(c)
	var uid int64
	v := session.Get("uid")
	if v == nil {
		r.codeResponse(c, code.USER_NOT_LOGIN)
		return
	} else {
		uid = v.(int64)
	}

	var msg ChangePasswordMsg
	err := c.ShouldBind(&msg)

	if err != nil {
		r.codeResponse(c, code.WrongParams)
		return
	}
	err = r.ctx.Dao.ChangePassword(uid, msg.Newpassword)
	if err != nil {
		r.codeResponse(c, code.DBError)
		return
	}
	r.codeResponse(c, code.Ok)
}
