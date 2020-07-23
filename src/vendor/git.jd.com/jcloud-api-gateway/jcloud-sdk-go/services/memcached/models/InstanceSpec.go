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

type InstanceSpec struct {

    /* 实例规格代码：
MC-S-1C1G	- 单机版，1核CPU，内存1G
MC-S-1C2G	- 单机版，1核CPU，内存2G
MC-S-1C4G	- 单机版，1核CPU，内存4G
MC-S-1C8G	- 单机版，1核CPU，内存8G
MC-S-1C16G - 单机版，1核CPU，内存16G
MC-S-1C32G - 单机版，1核CPU，内存32G
  */
    InstanceClass string `json:"instanceClass"`

    /* 实例类型（single：单机版，master-slave：主从版，cluster：集群版）
主从版、集群版暂不支持
  */
    InstanceType string `json:"instanceType"`

    /* 可用区  */
    AzId string `json:"azId"`

    /* 所属VPC的ID  */
    VpcId string `json:"vpcId"`

    /* 所属子网的ID  */
    SubnetId string `json:"subnetId"`

    /* 实例名称  */
    InstanceName string `json:"instanceName"`

    /* 实例描述 (Optional) */
    InstanceDescription *string `json:"instanceDescription"`

    /* Memcached的版本号  */
    McVersion string `json:"mcVersion"`

    /* 认证方式（true：需要认证，false：免密）  */
    McAuth bool `json:"mcAuth"`

    /* 密码 (Optional) */
    McPswd *string `json:"mcPswd"`

    /* 计费信息 (Optional) */
    Charge *charge.ChargeSpec `json:"charge"`

    /* IP协议版本，值为"v4&v6"表示支持IPv4和IPv6，值为"v4"或为空表示只支持IPv4 (Optional) */
    IpVersion *string `json:"ipVersion"`
}
