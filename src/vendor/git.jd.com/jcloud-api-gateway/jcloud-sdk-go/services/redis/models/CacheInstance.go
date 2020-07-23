// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models

import charge "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/charge/models"

type CacheInstance struct {

    /* 实例ID (Optional) */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 实例名称 (Optional) */
    CacheInstanceName string `json:"cacheInstanceName"`

    /* 规格代码，参考 https://docs.jdcloud.com/cn/jcs-for-redis/specifications (Optional) */
    CacheInstanceClass string `json:"cacheInstanceClass"`

    /* 实例的总内存（MB） (Optional) */
    CacheInstanceMemoryMB int `json:"cacheInstanceMemoryMB"`

    /* 实例状态：creating表示创建中，running表示运行中，error表示错误，changing表示变更规格中，deleting表示删除中，configuring表示修改参数中，restoring表示备份恢复中 (Optional) */
    CacheInstanceStatus string `json:"cacheInstanceStatus"`

    /* 实例描述 (Optional) */
    CacheInstanceDescription string `json:"cacheInstanceDescription"`

    /* 创建时间（ISO 8601标准的UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ） (Optional) */
    CreateTime string `json:"createTime"`

    /* az信息 (Optional) */
    AzId AzId `json:"azId"`

    /* 所属VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 所属子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 访问域名 (Optional) */
    ConnectionDomain string `json:"connectionDomain"`

    /* 端口 (Optional) */
    Port int `json:"port"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 实例的详细版本号，形如x.x-x.x (Optional) */
    InstanceVersion string `json:"instanceVersion"`

    /* 连接redis实例时，是否需要密码认证，false表示无密码 (Optional) */
    Auth bool `json:"auth"`

    /* 创建实例时选择的redis引擎版本：目前支持2.8和4.0 (Optional) */
    RedisVersion string `json:"redisVersion"`

    /* 实例类型：master-slave表示主从版，cluster表示集群版 (Optional) */
    CacheInstanceType string `json:"cacheInstanceType"`

    /* 是否支持IPv6，0表示不支持（只能用IPv4），1表示支持 (Optional) */
    Ipv6On int `json:"ipv6On"`

    /* 标签信息 (Optional) */
    Tags []Tag `json:"tags"`
}
