package events

type Resp struct {
	OK   []string `json:"ok"`
	Fail []string `json:"fail"`
}
