package dao

import (
	"log"
	"single-sentence/repository/model"
)

func GetUserById(ID uint) (user *model.User, err error) {
	if err = MysqlDao.Where("ID = ?", ID).Find(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
func GetUserByUsername(username string) (user *model.User, err error) {
	if err = MysqlDao.Where("User_name = ?", username).Find(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
func UpdateUser(user *model.User) (err error) {
	if err = MysqlDao.Save(user).Error; err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
func DeleteUserById(id uint) (err error) {
	if err = MysqlDao.Where("ID = ?", id).Delete(&model.User{}).Error; err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
func CreateUser(user *model.User) (err error) {
	if err = MysqlDao.Create(user).Error; err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}
