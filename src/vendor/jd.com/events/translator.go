package events

// 帮助相对容易地转换云翼事件

type Translator struct {
	EventContextDict map[string]string
	ArkContextDict   map[string]string
	DataWhiteList    map[string]struct{}
}
