package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const PassWordCost = 10 //密码加密难度
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	NickName       string
	Type           string `gorm:"default:'common'"`
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}
func (u *User) NewUser(Id uint, username, password, nickname string) error {
	//fmt.Println(password)
	err := u.SetPassword(password)
	if err != nil {
		return err
	}
	u.ID = Id
	u.UserName = username
	u.NickName = nickname

	return nil

}
