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


type IpResourceInfo struct {

    /* 公网 IP 地址 (Optional) */
    Ip string `json:"ip"`

    /* 安全状态, 0: 安全, 1: 清洗, 2: 黑洞 (Optional) */
    SafeStatus int `json:"safeStatus"`

    /* 所属地域 Id, cn-north-1: 华北-北京, cn-east-1: 华东-宿迁, cn-east-2: 华东-上海, cn-south-1: 华南-广州 (Optional) */
    Region string `json:"region"`

    /* 黑洞阈值，单位 bps (Optional) */
    BlackHoleThreshold int64 `json:"blackHoleThreshold"`

    /* 触发清洗的流量速率，单位 bps (Optional) */
    CleanThresholdBps int64 `json:"cleanThresholdBps"`

    /* 触发清洗的包速率，单位 pps (Optional) */
    CleanThresholdPps int64 `json:"cleanThresholdPps"`
}
