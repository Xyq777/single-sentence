package routes

import (
	"github.com/gin-gonic/gin"
	"single-sentence/api"
	"single-sentence/middleware"
)

func NewRouter() *gin.Engine {

	r := gin.Default()
	user := r.Group("/user")

	{
		user.Use(middleware.AuthMiddleware("common"))

		user.GET("/", api.ShowUserInfoHandler)
		user.PUT("/", api.UpdateUserInfoHandler)
		user.DELETE("/", api.DeleteUserHandler)
		user.GET("/sentences", api.GetUserSentencesHandler)
		user.POST("/sentences", api.AddUserSentenceHandler)
		user.PUT("/sentences/:sentenceId", api.UpdateUserSentenceHandler)
		user.DELETE("/sentences/:sentenceId", api.DeleteUserSentenceHandler)
		//user.GET("/logout", api.logout)
	}
	admin := r.Group("/admin")
	{
		admin.Use(middleware.AuthMiddleware("admin"))

		admin.DELETE("/:userId", api.DeleteUser)
		admin.DELETE("/sentences/:sentenceId", api.DeleteSentence)
		//admin.GET("/logout", api.logout)
	}
	r.GET("/", api.GetSentenceHandler)
	r.POST("/register", api.RegisterUserHandler)
	r.POST("/login", api.LoginUserHandler)

	return r
}
