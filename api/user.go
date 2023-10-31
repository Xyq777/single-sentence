package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"single-sentence/pkg/apiX"
	e "single-sentence/pkg/errorCode"
	"single-sentence/serializer"
	"single-sentence/service"
)

func ShowUserInfoHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	ID := apiX.AssertId(c, code)
	if ID == 0 {
		return
	}
	user, err := service.UserInfoShow(c, ID)
	if err != nil {
		log.Println(err)
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "未成功找到用户",
			"error":  err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"data":   user,
	})
	return
}
func UpdateUserInfoHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	ID := apiX.AssertId(c, code)
	if ID == 0 {
		return
	}
	var req serializer.UserUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		code = e.InvalidParams
		c.JSON(http.StatusNotFound, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "绑定JSON数据失败",
			"error":  err.Error(),
		})
		return
	}
	err = service.UserInfoUpdate(c, &req, ID)
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "更新用户信息失败",
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
func DeleteUserHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	ID := apiX.AssertId(c, code)
	if ID == 0 {
		return
	}
	err := service.UserDelete(c, ID)
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
func RegisterUserHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	var req serializer.UserRegister
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		code = e.InvalidParams
		c.JSON(http.StatusNotFound, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "绑定JSON数据失败",
			"error":  err.Error(),
		})
		return
	}
	//fmt.Println(req)
	err = service.UserRegister(c, &req)
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "用户注册失败",
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
func LoginUserHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	var req serializer.UserLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		code = e.InvalidParams
		c.JSON(http.StatusNotFound, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "绑定JSON数据失败",
			"error":  err.Error(),
		})
		return
	}
	token, err := service.UserLogin(c, &req)
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "用户登录错误",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"token":  token,
	})
	return

}
