package dao

type Group struct {
	Id            int    `json:"id" gorm:"primary_key" comment:"id"`
	Tenant        string `json:"tenant" len:"255" comment:"`
	AppName       string `json:"app_name" len:"255" comment:"app名" unique:"app_name"`
	GroupName     string `json:"group_name" len:"255" comment:"Group名"`
	Env           string `json:"env" len:"4095" comment:"环境变量"`
	TokenExpireAt int    `json:"token_expire_at" comment:"token过期时间"` // 内网安全
	TrafficJson   string `json:"traffic_json" len:"2047" comment:"分组流量配置"`
	ModelBase
}
