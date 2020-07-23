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

type Instance struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称，具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Cloud-Database-and-Cache/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例规格 (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* CPU核数 (Optional) */
    InstanceCPU int `json:"instanceCPU"`

    /* 内存，单位MB (Optional) */
    InstanceMemoryMB int `json:"instanceMemoryMB"`

    /* 实例类型 (Optional) */
    Engine string `json:"engine"`

    /* 实例版本 (Optional) */
    EngineVersion string `json:"engineVersion"`

    /* 实例域名 (Optional) */
    InstanceDomain string `json:"instanceDomain"`

    /* 实例端口 (Optional) */
    InstancePort string `json:"instancePort"`

    /* 可用区ID，第一个为主实例在的可用区，参见[地域及可用区对照表](../Enum-Definitions/Regions-AZ.md) (Optional) */
    AzId []string `json:"azId"`

    /* VPC ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 实例状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 实例更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 是否只读 (Optional) */
    Readonly string `json:"readonly"`

    /* Tag信息 (Optional) */
    Tags []Tag `json:"tags"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`
}
