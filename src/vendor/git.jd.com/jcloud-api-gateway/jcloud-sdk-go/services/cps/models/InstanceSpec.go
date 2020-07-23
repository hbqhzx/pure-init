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

    /* 可用区, 如 cn-east-1  */
    Az string `json:"az"`

    /* 实例类型, 如 cps.c.normal  */
    DeviceType string `json:"deviceType"`

    /* 主机名 (Optional) */
    Hostname *string `json:"hostname"`

    /* 镜像类型, 取值范围：standard  */
    ImageType string `json:"imageType"`

    /* 操作系统类型ID  */
    OsTypeId string `json:"osTypeId"`

    /* 系统盘RAID类型ID  */
    SysRaidTypeId string `json:"sysRaidTypeId"`

    /* 数据盘RAID类型ID  */
    DataRaidTypeId string `json:"dataRaidTypeId"`

    /* 子网编号 (Optional) */
    SubnetId *string `json:"subnetId"`

    /* 是否启用外网，取值范围：yes、no (Optional) */
    EnableInternet *string `json:"enableInternet"`

    /* 是否启用IPv6，取值范围：yes、no (Optional) */
    EnableIpv6 *string `json:"enableIpv6"`

    /* 网络类型，取值范围：basic、vpc  */
    NetworkType string `json:"networkType"`

    /* 网络CIDR (Optional) */
    Cidr *string `json:"cidr"`

    /* 内网IP (Optional) */
    PrivateIp *string `json:"privateIp"`

    /* 外网链路类型, 目前只支持bgp (Optional) */
    LineType *string `json:"lineType"`

    /* 外网带宽, 范围[1,200] 单位Mbps (Optional) */
    Bandwidth *int `json:"bandwidth"`

    /* 云物理服务器名称  */
    Name string `json:"name"`

    /* 云物理服务器描述 (Optional) */
    Description *string `json:"description"`

    /* 密码  */
    Password string `json:"password"`

    /* 购买数量  */
    Count int `json:"count"`

    /* 计费配置  */
    Charge *charge.ChargeSpec `json:"charge"`
}
