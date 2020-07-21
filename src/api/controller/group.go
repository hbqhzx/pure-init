/* controller.go -   */

/*
modification history
--------------------
2017/6/9, by lovejoy, create
2017/6/9 by lovejoy, modify,
*/
/*
DESCRIPTION
controller:
Usage:
    import ""

*/
package controller

import (
	"manager"
	. "manager/dao"
)

func ListGroup(ctx *Context) {
	query := ctx.query(R("app_name"), []O{"tenant", "category", "type"})
	appName := ctx.paramString("app_name")
	tenant := ctx.paramString("tenant")

	ctx.jsonSuccess(manager.ListGroup(query, appName, tenant))
}

//单个app的详情
func ListGroupDetail(ctx *Context) {
	query := ctx.query([]R{"app_name", "group_name"}, []O{"tenant"})
	ctx.jsonSuccess(manager.ListGroupDetail(query))
}

//更新group
func UpdateGroup(ctx *Context) {
	group := &Group{}
	ctx.param(group, []R{"app_name", "group_name"}, []O{"tenant", "env", "traffic_json", "token_expire_at"})
	category := ctx.paramString("category")
	ctx.jsonSuccess(manager.UpdateGroup(group, category))
}

// 更新分组 lb 配置
func UpdateGroupLB(ctx *Context) {
	group := &Group{}
	ctx.param(group, []R{"app_name", "group_name"}, []O{"tenant"})
	lb_type := ctx.paramString("lb_type")
	lb_cluster := ctx.paramString("lb_cluster")
	if lb_type != "jcloud" && lb_type != "jd" {
		ctx.jsonFailure("lb_type 只能是 jd 或 jcloud")
	}

	manager.UpdateGroupLB(group, lb_type, lb_cluster)
	ctx.jsonSuccess()
}
