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


type RegistryDetail struct {

    /* 实例的主键id (Optional) */
    Id int `json:"id"`

    /* 区域 (Optional) */
    Region string `json:"region"`

    /* 京舰请求Id（创建实例的请求id） (Optional) */
    RequestId string `json:"requestId"`

    /* 集群网络（实例所在的用户vpcId） (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网（实例所在的用户vpc内的子网id） (Optional) */
    SubNetId string `json:"subNetId"`

    /* 实例ID（注册中心实例id） (Optional) */
    RegistryId string `json:"registryId"`

    /* 集群名称（注册中心名称） (Optional) */
    RegistryName string `json:"registryName"`

    /* 集群类型（注册中心类型） (Optional) */
    RegistryType string `json:"registryType"`

    /* 集群状态（注册中心状态） (Optional) */
    RegistryStatus string `json:"registryStatus"`

    /* 创建账号的主账号信息 (Optional) */
    MasterUserPin string `json:"masterUserPin"`

    /* 创建账号的主账号id (Optional) */
    MasterUserId string `json:"masterUserId"`

    /* 创建实例的用户的用户pin（如果为子账号创建则为子账号pin 主账号创建为主账号pin） (Optional) */
    SubUserPin string `json:"subUserPin"`

    /* 创建实例的账号id  （如果为子账号创建则为子账号id 主账号创建为主账号id） (Optional) */
    SubUserId string `json:"subUserId"`

    /* 创建日期 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新日期 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 过期时间 (Optional) */
    ExpireTime string `json:"expireTime"`

    /* 订单Id (Optional) */
    OrderNum string `json:"orderNum"`

    /* 订单Source Id (Optional) */
    OrderSourceId string `json:"orderSourceId"`

    /* 计费类型 (Optional) */
    ChargeType string `json:"chargeType"`

    /* 注册中心版本 (Optional) */
    Version string `json:"version"`

    /* 注册中心服务实例数 (Optional) */
    InstanceNum string `json:"instanceNum"`

    /* 可用区（可用区A） (Optional) */
    RZ string `json:"rZ"`

    /* 可用区类型（如-系统分配） (Optional) */
    RzType string `json:"rzType"`

    /* 集群配置参数 (Optional) */
    ShowArguments string `json:"showArguments"`

    /* 集群扩容参数 (Optional) */
    ShowResizeArguments string `json:"showResizeArguments"`

    /* 集群描述 (Optional) */
    RegistryDesc string `json:"registryDesc"`
}