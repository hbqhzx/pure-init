# Event golang sdk


## api 地址

### 测试
env | uri
---- | ---
测试 | 10.226.145.170:8080/v1/cloudevent/_push
华东预发布二期 | eventdb-saver-api-hdpre.hawkeye.jcloud.com/v1/cloudevent/_push
预发布 | 10.180.152.219:8080/v1/cloudevent/_push

### 生产
env | uri
---- | ---
私有云 | eventdb-saver-api.hawkeye.jd.com/v1/cloudevent/_push
公有云 | eventdb-saver-api-hdpre.hawkeye.jcloud.com/v1/cloudevent/_push
DevOps | eventdb-saver-api-devops.hawkeye.jcloud.com/v1/cloudevent/_push

## 示例
详细请参考`example/normal/main.go`

```
var (
	uri = "http://10.226.145.170:8080/v1/cloudevent/_push"
)

func main() {

	// 设置 event context
	ec := events.NewEventContext("something"). //, "APP", []string{anyMap["app"].(string)}).
							SetTime(11111111111)

	// 设置 ark context
	ac := events.NewArkContext().
		SetTenant("jcloud").
		SetProduct("ark").
		SetRelatedProduct("ark").
		SetCorp("product").
		SetDept("tools").
		SetSys("hawkeye").
		SetApp("tsdb-saver").
		SetOperator("somebody").
		SetObjectType("APP").
		SetObjectName("tsdb-saver").
		SetOperation("do something").
		SetDesc("do something. success")

	data := make(map[string]interface{})
	data["key"] = "value"
	event := events.NewEvent(*ec, *ac, data)

	client := events.NewHttpClient(uri)
	err := client.Push(event)
	if err != nil {
		fmt.Println(err)
	}

}
```