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


type Instance struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例版本 (Optional) */
    InstanceVersion string `json:"instanceVersion"`

    /* 实例状态，running：运行，error：错误，creating：创建中，changing：变配中 (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例硬件配置规格 例如:2C4G20G (Optional) */
    InstanceFlavor string `json:"instanceFlavor"`

    /* az信息 (Optional) */
    AzId string `json:"azId"`

    /* 所属VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 所属子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* exposed domain (Optional) */
    ExposedDomain string `json:"exposedDomain"`

    /* 实例服务配置 (Optional) */
    ServiceConfig string `json:"serviceConfig"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`
}