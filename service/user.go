package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"single-sentence/pkg/jwtX"
	"single-sentence/repository/dao"
	"single-sentence/repository/model"
	"single-sentence/serializer"
)

func UserInfoShow(c *gin.Context, id uint) (user *model.User, err error) {
	user, err = dao.GetUserById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil

}
func UserInfoUpdate(c *gin.Context, req *serializer.UserUpdate, id uint) (err error) {
	var user = &model.User{}
	err = user.NewUser(id, "", req.Password, req.NickName)
	if err != nil {
		log.Panicln(err)
		return err
	}
	err = dao.UpdateUser(user)
	if err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
func UserDelete(c *gin.Context, id uint) (err error) {
	err = dao.DeleteUserById(id)
	if err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
func UserRegister(c *gin.Context, req *serializer.UserRegister) (err error) {
	var user = &model.User{}
	err = user.NewUser(0, req.UserName, req.Password, req.NickName)
	if err != nil {
		log.Println(err)
		return err
	}
	err = dao.CreateUser(user)
	if err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
func UserLogin(c *gin.Context, req *serializer.UserLogin) (token string, err error) {
	var user = &model.User{}
	user, err = dao.GetUserByUsername(req.UserName)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if !user.CheckPassword(req.Password) {
		return "", errors.New("密码错误")
	}
	token, err = jwtX.GenerateToken(user.ID, user.UserName)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}
