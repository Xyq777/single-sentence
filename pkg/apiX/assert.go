package apiX

import (
	"github.com/gin-gonic/gin"
	"net/http"
	e "single-sentence/pkg/errorCode"
)

func AssertId(c *gin.Context, code int) (ID uint) {
	result, exist := c.Get("ID")
	if !exist {
		code = e.ErrorUserNotFound
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
		})
		return 0
	}
	ID, ok := result.(uint)
	if !ok {
		code = e.InvalidParams
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "不合法的ID",
		})
		return 0

	}
	return ID
}
func AssertUsername(c *gin.Context, code int) (username string) {
	result, exist := c.Get("username")
	if !exist {
		code = e.ErrorUserNotFound
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
		})
		return ""
	}
	username, ok := result.(string)
	if !ok {
		code = e.InvalidParams
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "不合法的username",
		})
		return ""

	}
	return username
}
