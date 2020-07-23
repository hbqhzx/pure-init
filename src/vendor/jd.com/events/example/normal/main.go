package main

import (
	"fmt"

	"jd.com/events"
)

var (
	// 测试环境地址
	uri = "http://10.226.145.170:8080/v1/cloudevent/_push"
)

func main() {

	// 设置 event context
	// 必须提供事件类型
	ec := events.NewEventContext("something").
		SetTime(11111111111)

	/*
		示例:
		假设事件发生在
		租户 jcloud
		产品体系 product
		部门 工具部 tools
		产品线 ark
		系统 hawkeye
		应用 tsdb-saver
		操作人 somebody
		操作对象 tsdb-saver
		操作类型 应用 APP

	*/
	// 设置 ark context
	ac := events.NewArkContext().
		// jcloud 租户
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

	// 事件的其他附加数据:
	//  key: value
	data := make(map[string]interface{})
	data["key"] = "value"
	event := events.NewEvent(*ec, *ac, data)

	client := events.NewHttpClient(uri)

	err := client.Push(event)
	if err != nil {
		fmt.Println(err)
	}

}
