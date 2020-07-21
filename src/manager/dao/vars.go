/* dao.go -   */

/*
modification history
--------------------
2017/6/10, by lovejoy, create
2017/6/10 by lovejoy, modify,
*/
/*
DESCRIPTION
dao: 这个里面放所有的变量定义
Usage:
    import ""

*/
package dao

//app 表tower_type字段的定义
const (
	TOWER_TYPE_NORMAL = iota //包部署
	TOWER_TYPE_IAAS          //IAAS
)

//deploy_list中的type字段的定义
const (
	DeployTypeUpdate   = iota + 1 // 升级
	DeployTypeConf                // 配置升级
	DeployTypeRestart             // 重启
	DeployTypeRollback            //回滚 这个不会创建上线单 并且可能不会出现在这里
	DeployTypeStop                //停止
	DeployTypeStart               //启动
)

var IaasDeployTypeMap = map[int]string{
	DeployTypeUpdate:   "DEPLOY",
	DeployTypeRestart:  "RESTART",
	DeployTypeRollback: "ROLLBACK",
}

//上线单的状态
const (
	DEPLOY_STATUS_INIT        = iota // 创建中 前端理论不会get到
	DEPLOY_STATUS_DEPLOYING          // 正在部署
	DEPLOY_STATUS_FINISH             // 部署成功
	DEPLOY_STATUS_ERROR              // 部署失败
	DEPLOY_STATUS_ROLLBACKING        //回滚中
	DEPLOY_STATUS_ROLLBACK_SUCCESS
	DEPLOY_STATUS_ROLLBACK_FAIL
	DEPLOY_STATUS_CANCELED
)
const (
	INSTANCE_STATUS_INIT      = iota + 10 //未开始
	INSTANCE_STATUS_DEPLOYING             //部署中
	INSTANCE_STATUS_FINISH                //部署成功
	INSTANCE_STATUS_ERROR                 //部署失败

	INSTANCE_STAUTS_ROLLBACK_INIT //回滚未开始 14
	INSTANCE_STATUS_ROLLBACKING
	INSTANCE_STATUS_ROLLBACK_SUCCESS
	INSTANCE_STATUS_ROLLBACK_FAIL
	INSTANCE_STATUS_CANCELED
)

var DeployType map[int]string = map[int]string{
	1: "上线",
	2: "配置升级",
	3: "重启",
	4: "回滚",
	5: "停止",
	6: "启动",
}

const (
	GroupEnvPre     = "pre"
	GroupEnvProduct = "product"
	COMPILE_MODULE  = "COMPILE_MODULE"
	SHARE_MODULE    = "SHARE_MODULE"
)

const (
	ARKJCS = "ARKJCS"
)

const (
	MAX_DEPLOYLIST_COUNT = 10000
)
