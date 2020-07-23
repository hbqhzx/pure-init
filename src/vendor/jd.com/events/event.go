package events

import (
	"fmt"
	"time"
)

const (
	DefaultVersion = "v0.2"
	DefaultTenant  = "jcloud"
	GlobalSeen     = "GLOBAL"
)

// Event represents the canonical representation of a CloudEvent.
type Event struct {
	EventContext
	ArkContext
	Data Data `json:"data"`
}

func NewEvent(ec EventContext, ac ArkContext, data Data) *Event {
	return &Event{
		EventContext: ec,
		ArkContext:   ac,
		Data:         data,
	}
}

func (e *Event) Validate() error {
	err := e.EventContext.Validate()
	if err != nil {
		return err
	}
	err = e.ArkContext.Validate()
	if err != nil {
		return err
	}

	// 保证source == objectType, subject == objectName
	// TODO: 如何把 operation/objectType/objectName 跟
	// cloudevent 的 type/source/subject
	// 对应起来
	e.ensureObject()
	return nil
}

func (e *Event) ensureObject() {
	if e.EventContext.Source == "" {
		e.EventContext.Source = e.ArkContext.ObjectType
	}

	if len(e.EventContext.Subject) == 0 {
		e.EventContext.Subject = e.ArkContext.ObjectName
	}
}

type EventContext struct {
	// 可选,表示 event spec 版本。目前我们遵循 v0.2
	SpecVersion string `json:"specversion"`
	// 事件类型
	// 部署: deploy
	// 服务树: lord
	// 门神: doorgod
	// 其他类型待定
	Type string `json:"type"`
	// 根据 cloud event spec , 为事件产生者，倒排 uri 形式
	// 在 ark  环境中，暂时设置为等于 objectType
	Source string `json:"source"`
	// 根据 cloud event spec, 为具体事件主体，
	// 在 ark 环境中，暂时设置为等于 objectName
	Subject []string `json:"subject"`
	// 事件 id, 可选。
	Id string `json:"id"`
	// 事件时间，毫秒时间戳
	Time int64 `json:"time,omitempty"`
	// schema 仓库地址, 可选
	SchemaURL *string `json:"schemaurl,omitempty"`
	// ContentType， 可选
	ContentType *string `json:"contenttype,omitempty"`
}

func NewEventContext(typ string) *EventContext {
	return &EventContext{
		Type: typ,
	}
}

func (ec *EventContext) SetSource(source string) *EventContext     { ec.Source = source; return ec }
func (ec *EventContext) SetSubject(subject []string) *EventContext { ec.Subject = subject; return ec }

func (ec *EventContext) SetTime(ts int64) *EventContext {
	ec.Time = ts
	return ec
}

func (ec *EventContext) SetId(id string) *EventContext {
	ec.Id = id
	return ec
}

func (ec *EventContext) Validate() error {
	if ec.Type == "" {
		return fmt.Errorf("type field must be filled")
	}
	if ec.SpecVersion == "" {
		ec.SpecVersion = DefaultVersion
	}
	// earlier than 2018/1/1 00:00:00
	if ec.Time <= 1514736000000 {
		ec.Time = time.Now().UnixNano() / 1e6
	}
	return nil
}

type ArkContext struct {
	// 租户
	Tenant string `json:"tenant"`
	// 体系
	Corp string `json:"corp"`
	// 部门
	Dept string `json:"dept"`
	// 产品线
	Product string `json:"product"`
	// 此事件/操作影响到的产品线，即希望哪些产品线的人看到该时间
	// 例子:
	// 1. 操作没有产生跨产品线影响，则 RelatedProduct: []string{} ,留空列表
	// 2. 操作只影响本产品线自身，留空
	// 3. 操作影响别的产品线，譬如 产品线 prod0 的操作影响到了产品线 p1, p2，则
	//				Product: "prod0",
	//              RelatedProduct: ["p1", "p2"]
	// 4. 操作影响范围很大，为全局影响，则写入
	//              RelatedProduct: ["GLOBAL"]
	//    则在任何一个产品线节点都可以看到该事件
	RelatedProduct []string `json:"relatedProduct"`
	// 系统
	Sys string `json:"sys"`
	// 应用
	App string `json:"app"`
	// 操作人
	Operator string `json:"operator"`

	// 操作对象类型
	ObjectType string `json:"objectType"`
	// 操作对象名字,列表
	ObjectName []string `json:"objectName"`
	Operation  string   `json:"operation"`

	// 操作内容
	Desc string `json:"desc"`
}

func (ac *ArkContext) Validate() error {
	if ac.Tenant == "" {
		ac.SetTenant(DefaultTenant)
	}
	return nil
}

func NewArkContext() *ArkContext {
	return &ArkContext{}
}

func (ac *ArkContext) SetTenant(tenant string) *ArkContext {
	ac.Tenant = tenant
	return ac
}
func (ac *ArkContext) SetProduct(product string) *ArkContext {
	ac.Product = product
	return ac
}
func (ac *ArkContext) SetRelatedProduct(products ...string) *ArkContext {
	ac.RelatedProduct = products
	return ac
}

func (ac *ArkContext) GlobalSeen() *ArkContext { ac.SetRelatedProduct(GlobalSeen); return ac }

func (ac *ArkContext) SetCorp(corp string) *ArkContext         { ac.Corp = corp; return ac }
func (ac *ArkContext) SetDept(dept string) *ArkContext         { ac.Dept = dept; return ac }
func (ac *ArkContext) SetSys(sys string) *ArkContext           { ac.Sys = sys; return ac }
func (ac *ArkContext) SetApp(app string) *ArkContext           { ac.App = app; return ac }
func (ac *ArkContext) SetOperator(operator string) *ArkContext { ac.Operator = operator; return ac }

func (ac *ArkContext) SetDesc(desc string) *ArkContext { ac.Desc = desc; return ac }

func (ac *ArkContext) SetObjectType(ot string) *ArkContext       { ac.ObjectType = ot; return ac }
func (ac *ArkContext) SetObjectName(ob ...string) *ArkContext    { ac.ObjectName = ob; return ac }
func (ac *ArkContext) SetOperation(operation string) *ArkContext { ac.Operation = operation; return ac }

type Data map[string]interface{}
