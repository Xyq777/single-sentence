package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"single-sentence/pkg/apiX"
	e "single-sentence/pkg/errorCode"
	"single-sentence/pkg/randX"
	"single-sentence/repository/dao"
	"single-sentence/serializer"
	"single-sentence/service"
	"strconv"
)

func GetSentenceHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	queries := c.QueryArray("c")
	param := randX.GetRandParam(queries)
	if param == "" {
		code = e.InvalidParams
		c.JSON(http.StatusNotFound, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "无效的查询参数",
		})
		return
	}
	sentence, err := service.SentenceGet(param)
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "句子获取失败",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"data":   sentence,
	})
}
func GetUserSentencesHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	//从redis再拿的话，得再建一个数据结构，唉我也不知道要不要，先用mysql吧
	username := apiX.AssertUsername(c, code)
	//	fmt.Println(username)
	if username == "" {
		return
	}
	sentences, err := dao.GetAllSentences(username)
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "句子获取失败",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"data":   sentences,
	})
}
func AddUserSentenceHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	var req serializer.SentencePost
	err := c.ShouldBindJSON(&req)
	if err != nil {
		code = e.InvalidParams
		c.JSON(http.StatusNotFound, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "绑定JSON失败",
		})
		return
	}
	username := apiX.AssertUsername(c, code)
	if username == "" {
		return
	}
	sentence, err := service.SentencePost(&req, username)
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "句子返回失败",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"data":   sentence,
	})
}

func UpdateUserSentenceHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	var req serializer.SentenceUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		code = e.InvalidParams
		c.JSON(http.StatusNotFound, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "绑定JSON失败",
		})
		return
	}
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
	sentence, err := service.SentenceUpdate(&req, uint(sentenceId))
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "句子更新失败",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"data":   sentence,
	})

}
func DeleteUserSentenceHandler(c *gin.Context) {
	var code int
	code = e.SUCCESS
	username := apiX.AssertUsername(c, code)
	if username == "" {
		return
	}
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
	sentence, err := dao.GetSentenceById(uint(sentenceId))
	if err != nil {
		code = e.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "未找到要删除的句子",
			"error":  err.Error(),
		})
		return
	}
	if sentence.Creator != username {
		code = e.InvalidParams
		c.JSON(http.StatusForbidden, gin.H{
			"status": code,
			"msg":    e.GetMsg(code),
			"data":   "无法删除他人句子",
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
