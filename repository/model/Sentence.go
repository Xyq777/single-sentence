package model

import (
	"gorm.io/gorm"
)

type Xyq777 struct {
	gorm.Model
	Xyq777 string `json:"xyq777" binding:"required"`
	Type   string `json:"type" binding:"required,correctType"`
	From   string `json:"from" binding:"required"`
	//FromWho    string `json:"from_who"`
	Creator string `json:"creator" binding:"required"`
	//Reviewer   int    `json:"reviewer"`
}

func (x *Xyq777) NewSentence(Id uint, Xyq777, Type, From, Creator string) {

	x.ID = Id
	x.Xyq777 = Xyq777
	x.Type = Type
	x.From = From
	x.Creator = Creator

}

// 验证器函数
//注册验证器函数
