package dao

type Token struct {
	Id       int    `json:"id" gorm:"primary_key" comment:"id"`
	User     string `json:"user" len:"31" comment:"用户名"`
	Token    string `json:"token" gorm:"unique" len:"63" comment:"密码"`
	UserId   int    `json:"user_id" comment:"erp_id"`
	Email    string `json:"email" len:"31" comment:"邮箱" `
	Mobile   string `json:"mobile" len:"31" comment:"电话"`
	Platform string `json:"platform" gorm:"index:platform" len:"63" comment:"token来源"`
	ModelBase
}
