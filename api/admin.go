package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	e "single-sentence/pkg/errorCode"
	"single-sentence/service"
	"strconv"
)

func DeleteUser(c *gin.Context) {
	var code int
	code = e.SUCCESS
	isAdmin, exist := c.Get("isAdmin")
	if exist {
		if isAdmin.(bool) == true {
			userId, err := strconv.Atoi(c.Param("userId"))
			if err != nil {
				log.Println(err)
				code = e.InvalidParams
				c.JSON(http.StatusNotFound, gin.H{
					"status": code,
					"msg":    e.GetMsg(code),
					"data":   "不合法的用户id",
					"error":  err.Error(),
				})
				return
			}
			err = service.UserDelete(c, uint(userId))
			if err != nil {
				code = e.ERROR
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": code,
					"msg":    e.GetMsg(code),
					"data":   "删除用户信息失败",
					"error":  err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			return

		}
	}
}
func DeleteSentence(c *gin.Context) {
	var code int
	code = e.SUCCESS
	sentenceId, err := strconv.Atoi(c.Param("sentenceId"))
	if err != nil {
		code = e.InvalidParams
		c.JSON(http.StatusNotFound, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "不合法的句子ID",
		})
		return
	}

	err = service.SentenceDelete(uint(sentenceId))
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "句子删除失败",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
	})
}
