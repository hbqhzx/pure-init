package manager

import (
	"encoding/json"
	"lib/db"
	. "lib/exception"
	. "manager/dao"
	"regexp"
)

type GroupDetail struct {
	Group       *Group
	ConfVersion int    `json:"conf_vesion"`
	ConfUrl     string `json:"conf_url"`
	Category    string `json:"category"`
}
type GroupMeta struct {
	Group         *GroupResponse
	InstanceCount int    `json:"instance_count"`
	ConfVersion   int    `json:"conf_version"`
	ConfUrl       string `json:"conf_url"`
	Category      string `json:"category"`
}

type GroupResponse struct {
	Tenant      string `json:"tenant" len:"255" comment:"`
	AppName     string `json:"app_name" len:"255" comment:"app名"`
	GroupName   string `json:"group_name" len:"255" comment:"Group名"`
	GroupNameCn string `json:"group_name_cn" len:"255" comment:"Group中文名"`
	Env         string `json:"env" len:"4095" comment:"环境变量"`
	TrafficJson string `json:"traffic_json" len:"2047" comment:"分组流量配置"`
}

func ListGroup(query *db.Query, appName string, tenant string) []GroupMeta {

	//通过group和app获取所有instance.如果仅获取配置，不需要统计这个数据
	groupMetas := []GroupMeta{}

	var groups []Group
	TryCatch(func() {
		db.FindAll(&groups, map[string]interface{}{
			"tenant":     tenant,
			"app_name":   appName,
		})
	}, func(e interface{}) {
		if e == db.ErrRecordNotFound {
		} else {
			Throw(e)
		}
	})

	groupMap := make(map[string]Group, len(groups))
	if len(groups) > 0 {
		for _, group := range groups {
			groupMap[group.GroupName] = group
		}
	}

	return groupMetas
}

func ListGroupDetail(query *db.Query) *GroupDetail {
	group := &Group{}
	ret := GroupDetail{group, 0, "", ""}
	return &ret
}

type Env struct {
	Key  string `json:"key"`
	Name string `json:"value"`
}

func UpdateGroup(group *Group, category string) *GroupDetail {
	existsGroup := Group{}
	TryCatch(func() {
		db.Find(&existsGroup, map[string]interface{}{
			"tenant":     group.Tenant,
			"app_name":   group.AppName,
			"group_name": group.GroupName,
		})
	}, func(e interface{}) {
	})

	if group.TrafficJson == "" {
		if existsGroup.TrafficJson != "" {
			group.TrafficJson = existsGroup.TrafficJson
		} else {
			group.TrafficJson = "{}"
		}
	} else {
		trfc := map[string]string{}
		if err := json.Unmarshal([]byte(group.TrafficJson), &trfc); err != nil {
			Throw("流量配置不是合法的 json")
		}
	}

	if group.Env == "" {
		if existsGroup.Env != "" {
			group.Env = existsGroup.Env
		} else {
			group.Env = "[]"
		}
	} else {
		var envs []Env
		if err := json.Unmarshal([]byte(group.Env), &envs); err != nil {
			Throw("环境变量参数格式有误")
		}
		for _, v := range envs {
			if matched, err := regexp.MatchString("^[A-Za-z_][A-Za-z0-9_]*$", v.Key); !matched || err != nil {
				Throw(v.Key + ": 环境变量需要以字母开头仅包含字母和数字^[A-Za-z_][A-Za-z0-9_]*")
			}
		}
	}

	if existsGroup.Id > 0 {
		existsGroup.Env = group.Env
		existsGroup.TrafficJson = group.TrafficJson
		existsGroup.TokenExpireAt = group.TokenExpireAt
		db.Save(&existsGroup)
	} else {
		db.Create(group)
	}

	return &GroupDetail{group, 0, "", category}
}

func UpdateGroupLB(group *Group, lb_type, lb_cluster string) {
	traffic_json, err := json.Marshal(map[string]string{
		"lb_type":    lb_type,
		"lb_cluster": lb_cluster,
	})
	if err != nil {
		Throw("参数有误: " + err.Error())
	}

	db.UpdateOne(&Group{}, map[string]interface{}{
		"traffic_json": string(traffic_json),
	}, map[string]interface{}{
		"tenant":     group.Tenant,
		"app_name":   group.AppName,
		"group_name": group.GroupName,
	})
}
