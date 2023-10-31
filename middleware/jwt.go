package middleware

import (
	"github.com/gin-gonic/gin"
	e "single-sentence/pkg/errorCode"
	"single-sentence/pkg/jwtX"
	"single-sentence/repository/dao"
)

func AuthMiddleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		Token := c.GetHeader("Authorization")
		if Token == "" {
			code = e.InvalidParams
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token不能为空",
			})
			c.Abort()
			return
		}
		claims, err := jwtX.ParseToken(Token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   err.Error(),
			})
			c.Abort()
			return
		}
		user, err := dao.GetUserById(claims.ID)
		if err != nil || user.UserName != claims.Username {
			code = e.ErrorAuthCheckTokenFail
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   err.Error(),
			})
			c.Abort()
			return
		}
		if user.Type == "admin" {
			c.Set("isAdmin", true)
			c.Next()
			return
		}
		c.Set("username", claims.Username)
		c.Set("ID", claims.ID)
		c.Next()

	}
}
