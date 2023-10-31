package serializer

type UserRegister struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}
type UserUpdate struct {
	Password string `json:"password" `
	NickName string `json:"nickName"`
}
type UserLogin struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
type SentencePost struct {
	Xyq777 string `json:"xyq777" binding:"required"`
	Type   string `json:"type" binding:"required,correctType"`
	From   string `json:"from" binding:"required"`
}
type SentenceUpdate struct {
	Xyq777 string `json:"xyq777" `
	Type   string `json:"type" binding:"correctType"`
	From   string `json:"from" `
}
