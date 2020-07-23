package deploy

// Legacy 是根据生产环境 es 中查询出来的部署事件，模拟的部署事件结构，根据实际情况
// 替换成相应的部署事件 struct
type Legacy struct {
	Id     string `json:"id"`
	Tenant string `json:"tenant"`
	Corp   string `json:"corp"`
	Dept   string `json:"dept"`
	Pdl    string `json:"pdl"`
	Sys    string `json:"sys"`
	App    string `json:"app"`

	Operation  string `json:"operation"`
	Status     string `json:"status"`
	Operator   string `json:"operator"`
	ObjectType string `json:"objectType"`

	Group      []string `json:"group"`
	Deployment string   `json:"deployment"`
	Env        string   `json:"env"`
	Time       int64    `json:"start_time"`
	Version    string   `json:"version"`
}
