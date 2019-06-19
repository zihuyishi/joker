package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/zihuyishi/joker/web/dao"
)

type Context struct {
	Dao *dao.Dao
	G *gin.Engine
}

