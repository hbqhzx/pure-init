package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"jd.com/events"
	"jd.com/events/convert/deploy"
)

var byt = []byte(`{
	"app": "wizard-tp-used",
	"cfg_version": [
		""
	],
	"corp": "product",
	"deployment": "deploy-server",
	"end_time": 1556278040,
	"env": "STAGING",
	"group": [
		"online"
	],
	"host_ip": [
		"10.160.13.243"
	],
	"id": "deploy-server_227341",
	"idc": [],
	"ins": [
		"0.wizard-tp-used"
	],
	"operation": "DeployTypeUpdate",
	"operator": "zhoushuke",
	"pdl": "ark",
	"start_time": 1556278039,
	"status": "DEPLOY_STATUS_FINISH",
	"sys": "tpmonitor-ark",
	"system": "deploy-server",
	"tenant": "",
	"update_time": 1556278040
}`)

func main() {
	var sample deploy.Legacy
	json.Unmarshal(byt, &sample)
	e, err := ConvertToEvent(sample)
	if err != nil {
		fmt.Println(err)
		return
	}
	var uri = "http://eventdb-saver-api-hdpre.hawkeye.jcloud.com/v1/cloudevent/_push"
	//uri = "http://10.226.145.170:8080/v1/cloudevent/_push"

	client := events.NewHttpClient(uri)
	err = client.Push(e)
	fmt.Println(err)

	//fmt.Printf("%+v", sample)
}

// 部署事件  Data 中只存储如下字段即可
var DataWhiteList = []string{"pdl", "deployment", "version", "env", "group"}

// operation/status 中英文映射，方便拼写 desc
var dict = map[string]map[string]string{
	"operation": map[string]string{
		"DeployTypeUpdate":   "上线",
		"DeployTypeStart":    "启动",
		"DeployTypeStop":     "停止",
		"DeployTypeRestart":  "重启",
		"DeployTypeConf":     "配置更新",
		"DeployTypeRollback": "回滚",
		"TurnOn":             "流量开启",
		"TurnOff":            "流量关闭",
	},
	"status": map[string]string{
		"DEPLOY_STATUS_DEPLOYING":        "正在部署",
		"DEPLOY_STATUS_FINISH":           "部署成功",
		"DEPLOY_STATUS_ERROR":            "部署失败",
		"DEPLOY_STATUS_CANCELED":         "部署取消",
		"DEPLOY_STATUS_ROLLBACKING":      "回滚中",
		"DEPLOY_STATUS_ROLLBACK_SUCCESS": "回滚成功",
		"DEPLOY_STATUS_ROLLBACK_FAIL":    "回滚失败",
	},
}

func ConvertToEvent(any interface{}) (*events.Event, error) {
	var anyMap map[string]interface{}
	byts, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewBuffer(byts))
	decoder.UseNumber()
	decoder.Decode(&anyMap)
	//json.Unmarshal(byts, &anyMap)
	//fmt.Println(string(byts))

	var ts int64
	ts, err = anyMap["start_time"].(json.Number).Int64()
	if err != nil {
		ts = time.Now().UnixNano() / 1e6
	} else {
		ts *= 1e3
	}
	// 设置 ec
	ec := events.NewEventContext("deploy"). //, "APP", []string{anyMap["app"].(string)}).
						SetTime(ts).
						SetId(anyMap["id"].(string))

	// 设置 ac
	ac := events.NewArkContext().
		SetTenant(anyMap["tenant"].(string)).
		SetProduct(anyMap["pdl"].(string)).
		SetRelatedProduct(anyMap["pdl"].(string)).
		SetCorp(anyMap["corp"].(string)).
		SetDept(anyMap["dept"].(string)).
		SetSys(anyMap["sys"].(string)).
		SetApp(anyMap["app"].(string)).
		SetOperator(anyMap["operator"].(string)).
		SetObjectType("APP").
		SetObjectName([]string{anyMap["app"].(string)}...).
		SetOperation(anyMap["operation"].(string))

	//TODO: 拼装部署事件desc
	descFields := []string{"operation", "status"}
	var descValues []string
	for _, field := range descFields {
		descValues = append(descValues, dict[field][anyMap[field].(string)])
	}
	desc := strings.Join(descValues, ", ")
	ac.SetDesc(desc)

	data := make(map[string]interface{})
	for _, key := range DataWhiteList {
		data[key] = anyMap[key]
	}

	event := events.NewEvent(*ec, *ac, data)
	event.Validate()

	byts, err = json.Marshal(event)
	fmt.Println(string(byts))
	return event, nil

}
