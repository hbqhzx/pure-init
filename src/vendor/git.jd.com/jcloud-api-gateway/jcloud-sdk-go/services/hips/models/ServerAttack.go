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


type ServerAttack struct {

    /* 时间 (Optional) */
    Time string `json:"time"`

    /* ip地址 (Optional) */
    CreateTime string `json:"createTime"`

    /* 区域 (Optional) */
    RegionId string `json:"regionId"`

    /* 攻击类型 (Optional) */
    CrackType string `json:"crackType"`

    /* IP (Optional) */
    Ip string `json:"ip"`

    /* 端口 (Optional) */
    Port int `json:"port"`

    /* 公网IP (Optional) */
    FloatingIp string `json:"floatingIp"`

    /* 固定IP (Optional) */
    FixIp string `json:"fixIp"`

    /* 详细信息 (Optional) */
    DetailInfo string `json:"detailInfo"`

    /* 事件类型 (Optional) */
    Type string `json:"type"`

    /* 主机名 (Optional) */
    ServerName string `json:"serverName"`
}
