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


type TaskInfo struct {

    /* 探测异常数  ，null代表缺值。只统计探测失败，超时的个数。 (Optional) */
    AbnormalCount int64 `json:"abnormalCount"`

    /* task的探测地址 (Optional) */
    Address string `json:"address"`

    /* 任务状态[false：己禁用，true：己启用] (Optional) */
    Enabled bool `json:"enabled"`

    /* task名称 (Optional) */
    Name string `json:"name"`

    /* task的可用率 (Optional) */
    ProbeAvailability float64 `json:"probeAvailability"`

    /* 该task的探测源个数 (Optional) */
    ProbeCount int64 `json:"probeCount"`

    /* task的探测类型，1：http，2：telnet (Optional) */
    ProbeType int64 `json:"probeType"`

    /* task的探测平均响应时间 (Optional) */
    ResponseTime float64 `json:"responseTime"`

    /* task的id (Optional) */
    TaskId string `json:"taskId"`
}
