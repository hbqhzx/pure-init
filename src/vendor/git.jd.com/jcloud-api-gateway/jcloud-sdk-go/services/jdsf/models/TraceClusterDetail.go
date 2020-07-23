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


type TraceClusterDetail struct {

    /* 链路跟踪集群 id (Optional) */
    InstanceId string `json:"instanceId"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 集群状态 (Optional) */
    TraceClusterState string `json:"traceClusterState"`

    /* 集群实例配置名称 (Optional) */
    TraceConfigName string `json:"traceConfigName"`

    /* 集群名称 (Optional) */
    TraceClusterName string `json:"traceClusterName"`

    /* 描述信息 (Optional) */
    TraceClusterDesc string `json:"traceClusterDesc"`

    /* 用户虚拟网络 id (Optional) */
    VpcId string `json:"vpcId"`

    /* 用户虚拟网络名称 (Optional) */
    VpcName string `json:"vpcName"`

    /* 用户虚拟子网 id (Optional) */
    SubnetId string `json:"subnetId"`

    /* 用户虚拟子网名称 (Optional) */
    SubnetName string `json:"subnetName"`

    /* 用户连接 调用链服务收集器 地址 (Optional) */
    HttpEndPoint string `json:"httpEndPoint"`

    /* 用户连接服务收集器 thrift 协议连接地址 (Optional) */
    TchannelEndPoint string `json:"tchannelEndPoint"`

    /* 用户数据收集服务 zipkin 兼容协议连接地址 (Optional) */
    ZipkinEndPoint string `json:"zipkinEndPoint"`

    /* 日志上报TPS计费项配置 (Optional) */
    TpsSpec string `json:"tpsSpec"`
}